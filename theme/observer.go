package theme

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Observer receives notifications about theme system events.
type Observer interface {
	OnThemeRegistered(ctx context.Context, event *ThemeEvent)
	OnThemeUpdated(ctx context.Context, event *ThemeEvent)
	OnThemeDeleted(ctx context.Context, event *ThemeEvent)
	OnThemeCompiled(ctx context.Context, event *CompilationEvent)
	OnValidationFailed(ctx context.Context, event *ValidationEvent)
	OnCacheHit(ctx context.Context, event *CacheEvent)
	OnCacheMiss(ctx context.Context, event *CacheEvent)
	OnTenantConfigured(ctx context.Context, event *TenantEvent)
	OnShutdown(ctx context.Context) error
}

// ObservableManager wraps a Manager with observer support.
type ObservableManager struct {
	*Manager
	observers []Observer
	mu        sync.RWMutex
}

// NewObservableManager creates a manager with observer support.
func NewObservableManager(config *ManagerConfig, storage Storage, evaluator ConditionEvaluator) (*ObservableManager, error) {
	manager, err := NewManager(config, storage, evaluator)
	if err != nil {
		return nil, err
	}

	return &ObservableManager{
		Manager:   manager,
		observers: make([]Observer, 0),
	}, nil
}

// AddObserver adds an observer to receive events.
func (om *ObservableManager) AddObserver(observer Observer) {
	om.mu.Lock()
	defer om.mu.Unlock()
	om.observers = append(om.observers, observer)
}

// RemoveObserver removes an observer.
func (om *ObservableManager) RemoveObserver(observer Observer) {
	om.mu.Lock()
	defer om.mu.Unlock()

	for i, obs := range om.observers {
		if obs == observer {
			om.observers = append(om.observers[:i], om.observers[i+1:]...)
			break
		}
	}
}

// RegisterTheme registers a theme and notifies observers.
func (om *ObservableManager) RegisterTheme(ctx context.Context, theme *Theme) error {
	err := om.Manager.RegisterTheme(ctx, theme)
	if err != nil {
		return err
	}

	event := &ThemeEvent{
		ThemeID:   theme.ID,
		ThemeName: theme.Name,
		Action:    "registered",
		Timestamp: time.Now(),
	}

	om.notifyThemeRegistered(ctx, event)
	return nil
}

// UpdateTheme updates a theme and notifies observers.
func (om *ObservableManager) UpdateTheme(ctx context.Context, theme *Theme) error {
	err := om.Manager.UpdateTheme(ctx, theme)
	if err != nil {
		return err
	}

	event := &ThemeEvent{
		ThemeID:   theme.ID,
		ThemeName: theme.Name,
		Action:    "updated",
		Timestamp: time.Now(),
	}

	om.notifyThemeUpdated(ctx, event)
	return nil
}

// DeleteTheme deletes a theme and notifies observers.
func (om *ObservableManager) DeleteTheme(ctx context.Context, themeID string) error {
	err := om.Manager.DeleteTheme(ctx, themeID)
	if err != nil {
		return err
	}

	event := &ThemeEvent{
		ThemeID:   themeID,
		Action:    "deleted",
		Timestamp: time.Now(),
	}

	om.notifyThemeDeleted(ctx, event)
	return nil
}

// GetTheme retrieves and compiles a theme, notifying observers of compilation.
func (om *ObservableManager) GetTheme(ctx context.Context, themeID string, evalData map[string]any) (*CompiledTheme, error) {
	start := time.Now()

	compiled, err := om.Manager.GetTheme(ctx, themeID, evalData)
	if err != nil {
		return nil, err
	}

	event := &CompilationEvent{
		ThemeID:   themeID,
		TenantID:  getTenantID(ctx),
		Duration:  time.Since(start),
		CSSSize:   len(compiled.CSS),
		Timestamp: time.Now(),
	}

	om.notifyThemeCompiled(ctx, event)
	return compiled, nil
}

// notifyThemeRegistered notifies all observers of theme registration.
func (om *ObservableManager) notifyThemeRegistered(ctx context.Context, event *ThemeEvent) {
	om.mu.RLock()
	observers := make([]Observer, len(om.observers))
	copy(observers, om.observers)
	om.mu.RUnlock()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for _, observer := range observers {
		select {
		case <-ctx.Done():
			// TODO: Log timeout/cancellation
			return
		default:
			wg.Add(1)
			go func(obs Observer) {
				defer wg.Done()
				obs.OnThemeRegistered(ctx, event)
			}(observer)
		}
	}

	// NOTE: Optional: wait for all started goroutines to complete
	wg.Wait()
}

// notifyThemeUpdated notifies all observers of theme updates.
func (om *ObservableManager) notifyThemeUpdated(ctx context.Context, event *ThemeEvent) {
	om.mu.RLock()
	observers := make([]Observer, len(om.observers))
	copy(observers, om.observers)
	om.mu.RUnlock()

	for _, observer := range observers {
		go observer.OnThemeUpdated(ctx, event)
	}
}

// notifyThemeDeleted notifies all observers of theme deletion.
func (om *ObservableManager) notifyThemeDeleted(ctx context.Context, event *ThemeEvent) {
	om.mu.RLock()
	observers := make([]Observer, len(om.observers))
	copy(observers, om.observers)
	om.mu.RUnlock()

	for _, observer := range observers {
		go observer.OnThemeDeleted(ctx, event)
	}
}

// notifyThemeCompiled notifies all observers of theme compilation.
func (om *ObservableManager) notifyThemeCompiled(ctx context.Context, event *CompilationEvent) {
	om.mu.RLock()
	observers := make([]Observer, len(om.observers))
	copy(observers, om.observers)
	om.mu.RUnlock()

	for _, observer := range observers {
		go observer.OnThemeCompiled(ctx, event)
	}
}

// notifyShutdown notifies all observers of system shutdown with graceful termination.
func (om *ObservableManager) notifyShutdown(ctx context.Context) error {
	om.mu.RLock()
	observers := make([]Observer, len(om.observers))
	copy(observers, om.observers)
	om.mu.RUnlock()

	if len(observers) == 0 {
		return nil
	}

	// Set a reasonable timeout for shutdown notifications
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		errors []error
	)

	// Notify all observers concurrently
	for _, observer := range observers {
		wg.Add(1)
		go func(obs Observer) {
			defer wg.Done()

			if err := obs.OnShutdown(ctx); err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("observer %T: %w", obs, err))
				mu.Unlock()
			}
		}(observer)
	}

	// Wait for all notifications to complete or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait for completion or context cancellation
	select {
	case <-done:
		// All observers notified
	case <-ctx.Done():
		// Timeout or cancellation occurred
		mu.Lock()
		errors = append(errors, fmt.Errorf("shutdown notification timeout: %w", ctx.Err()))
		mu.Unlock()
	}

	// Return combined errors if any occurred
	if len(errors) > 0 {
		return fmt.Errorf("shutdown notifications completed with errors: %v", errors)
	}

	return nil
}

// Shutdown gracefully shuts down the observable manager and notifies all observers.
func (om *ObservableManager) Shutdown(ctx context.Context) error {
	// First notify all observers of shutdown
	if err := om.notifyShutdown(ctx); err != nil {
		// TODO:Log the error but continue with manager shutdownc
		// You might want to inject a logger here
		fmt.Printf("Warning during observer shutdown: %v\n", err)
	}

	// Then perform any manager-specific cleanup
	// This would call the underlying Manager's shutdown if it exists
	return nil
}

// ThemeEvent represents a theme lifecycle event.
type ThemeEvent struct {
	ThemeID   string
	ThemeName string
	Action    string
	Timestamp time.Time
	Metadata  map[string]any
}

// CompilationEvent represents a theme compilation event.
type CompilationEvent struct {
	ThemeID   string
	TenantID  string
	Duration  time.Duration
	CSSSize   int
	FromCache bool
	Timestamp time.Time
	Metadata  map[string]any
}

// ValidationEvent represents a validation event.
type ValidationEvent struct {
	ThemeID    string
	Valid      bool
	IssueCount int
	ErrorCount int
	Timestamp  time.Time
	Issues     []ValidationIssue
	Metadata   map[string]any
}

// CacheEvent represents a cache operation event.
type CacheEvent struct {
	Key       string
	Operation string
	Hit       bool
	Duration  time.Duration
	Timestamp time.Time
}

// TenantEvent represents a tenant configuration event.
type TenantEvent struct {
	TenantID  string
	Action    string
	ThemeID   string
	Timestamp time.Time
	Metadata  map[string]any
}

// LoggingObserver is an observer that logs events using structured logging.
type LoggingObserver struct {
	logger Logger
}

// Logger interface for structured logging.
type Logger interface {
	Info(msg string, keysAndValues ...any)
	Error(msg string, err error, keysAndValues ...any)
	Debug(msg string, keysAndValues ...any)
}

// NewLoggingObserver creates a new logging observer.
func NewLoggingObserver(logger Logger) *LoggingObserver {
	return &LoggingObserver{logger: logger}
}

// OnThemeRegistered logs theme registration.
func (lo *LoggingObserver) OnThemeRegistered(ctx context.Context, event *ThemeEvent) {
	lo.logger.Info("theme registered",
		"themeID", event.ThemeID,
		"themeName", event.ThemeName,
		"timestamp", event.Timestamp,
	)
}

// OnThemeUpdated logs theme updates.
func (lo *LoggingObserver) OnThemeUpdated(ctx context.Context, event *ThemeEvent) {
	lo.logger.Info("theme updated",
		"themeID", event.ThemeID,
		"themeName", event.ThemeName,
		"timestamp", event.Timestamp,
	)
}

// OnThemeDeleted logs theme deletion.
func (lo *LoggingObserver) OnThemeDeleted(ctx context.Context, event *ThemeEvent) {
	lo.logger.Info("theme deleted",
		"themeID", event.ThemeID,
		"timestamp", event.Timestamp,
	)
}

// OnThemeCompiled logs theme compilation.
func (lo *LoggingObserver) OnThemeCompiled(ctx context.Context, event *CompilationEvent) {
	lo.logger.Debug("theme compiled",
		"themeID", event.ThemeID,
		"tenantID", event.TenantID,
		"duration", event.Duration,
		"cssSize", event.CSSSize,
		"fromCache", event.FromCache,
	)
}

// OnValidationFailed logs validation failures.
func (lo *LoggingObserver) OnValidationFailed(ctx context.Context, event *ValidationEvent) {
	lo.logger.Error("theme validation failed", nil,
		"themeID", event.ThemeID,
		"issueCount", event.IssueCount,
		"errorCount", event.ErrorCount,
	)
}

// OnCacheHit logs cache hits.
func (lo *LoggingObserver) OnCacheHit(ctx context.Context, event *CacheEvent) {
	lo.logger.Debug("cache hit",
		"key", event.Key,
		"duration", event.Duration,
	)
}

// OnCacheMiss logs cache misses.
func (lo *LoggingObserver) OnCacheMiss(ctx context.Context, event *CacheEvent) {
	lo.logger.Debug("cache miss",
		"key", event.Key,
	)
}

// OnTenantConfigured logs tenant configuration changes.
func (lo *LoggingObserver) OnTenantConfigured(ctx context.Context, event *TenantEvent) {
	lo.logger.Info("tenant configured",
		"tenantID", event.TenantID,
		"action", event.Action,
		"themeID", event.ThemeID,
	)
}

// OnShutdown handles logging observer shutdown.
func (lo *LoggingObserver) OnShutdown(ctx context.Context) error {
	lo.logger.Info("logging observer shutting down")
	// TODO: Perform any cleanup needed for the logging observer
	// For example, flush any buffered logs
	return nil
}

// MetricsObserver is an observer that collects metrics.
type MetricsObserver struct {
	mu sync.RWMutex

	themeRegistrations uint64
	themeUpdates       uint64
	themeDeletions     uint64
	themeCompilations  uint64
	validationFailures uint64
	cacheHits          uint64
	cacheMisses        uint64

	totalCompilationTime time.Duration
	avgCSSSize           float64
	cssCount             uint64
}

// NewMetricsObserver creates a new metrics observer.
func NewMetricsObserver() *MetricsObserver {
	return &MetricsObserver{}
}

// OnThemeRegistered records theme registration metrics.
func (mo *MetricsObserver) OnThemeRegistered(ctx context.Context, event *ThemeEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	mo.themeRegistrations++
}

// OnThemeUpdated records theme update metrics.
func (mo *MetricsObserver) OnThemeUpdated(ctx context.Context, event *ThemeEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	mo.themeUpdates++
}

// OnThemeDeleted records theme deletion metrics.
func (mo *MetricsObserver) OnThemeDeleted(ctx context.Context, event *ThemeEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	mo.themeDeletions++
}

// OnThemeCompiled records compilation metrics.
func (mo *MetricsObserver) OnThemeCompiled(ctx context.Context, event *CompilationEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()

	mo.themeCompilations++
	mo.totalCompilationTime += event.Duration

	mo.cssCount++
	mo.avgCSSSize = (mo.avgCSSSize*float64(mo.cssCount-1) + float64(event.CSSSize)) / float64(mo.cssCount)
}

// OnValidationFailed records validation failure metrics.
func (mo *MetricsObserver) OnValidationFailed(ctx context.Context, event *ValidationEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	mo.validationFailures++
}

// OnCacheHit records cache hit metrics.
func (mo *MetricsObserver) OnCacheHit(ctx context.Context, event *CacheEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	mo.cacheHits++
}

// OnCacheMiss records cache miss metrics.
func (mo *MetricsObserver) OnCacheMiss(ctx context.Context, event *CacheEvent) {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	mo.cacheMisses++
}

// OnTenantConfigured records tenant configuration metrics.
func (mo *MetricsObserver) OnTenantConfigured(ctx context.Context, event *TenantEvent) {
	// TODO: POptional: track tenant configuration changes
}

// OnShutdown handles metrics observer shutdown.
func (mo *MetricsObserver) OnShutdown(ctx context.Context) error {
	// TODO: Perform any cleanup needed for the metrics observer
	// For example, export final metrics or close connections
	return nil
}

// GetMetrics returns current metrics.
func (mo *MetricsObserver) GetMetrics() *Metrics {
	mo.mu.RLock()
	defer mo.mu.RUnlock()

	var avgCompilationTime time.Duration
	if mo.themeCompilations > 0 {
		avgCompilationTime = mo.totalCompilationTime / time.Duration(mo.themeCompilations)
	}

	cacheHitRate := 0.0
	totalCacheOps := mo.cacheHits + mo.cacheMisses
	if totalCacheOps > 0 {
		cacheHitRate = float64(mo.cacheHits) / float64(totalCacheOps)
	}

	return &Metrics{
		ThemeRegistrations: mo.themeRegistrations,
		ThemeUpdates:       mo.themeUpdates,
		ThemeDeletions:     mo.themeDeletions,
		ThemeCompilations:  mo.themeCompilations,
		ValidationFailures: mo.validationFailures,
		CacheHits:          mo.cacheHits,
		CacheMisses:        mo.cacheMisses,
		CacheHitRate:       cacheHitRate,
		AvgCompilationTime: avgCompilationTime,
		AvgCSSSize:         int(mo.avgCSSSize),
	}
}

// Metrics contains system metrics.
type Metrics struct {
	ThemeRegistrations uint64
	ThemeUpdates       uint64
	ThemeDeletions     uint64
	ThemeCompilations  uint64
	ValidationFailures uint64
	CacheHits          uint64
	CacheMisses        uint64
	CacheHitRate       float64
	AvgCompilationTime time.Duration
	AvgCSSSize         int
}

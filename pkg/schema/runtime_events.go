// pkg/schema/runtime/events.go
package schema
import (
	"context"
	"fmt"
	"sync"
	"time"
)
// EventHandler manages runtime events with dependency injection
// No direct dependency on Runtime to avoid circular references
// Uses types from schema package to avoid import cycles
type EventHandler struct {
	// Event callbacks by type
	handlers map[EventType][]EventCallback
	// Configuration
	validationTiming ValidationTiming
	enabled          bool
	// Event tracking (optional)
	tracker *EventTracker
	// Concurrency control
	mu sync.RWMutex
}
// NewEventHandler creates a new event handler
func NewEventHandler() *EventHandler {
	return NewEventHandlerWithConfig(DefaultEventHandlerConfig())
}
// NewEventHandlerWithConfig creates a new event handler with custom config
func NewEventHandlerWithConfig(config *EventHandlerConfig) *EventHandler {
	handler := &EventHandler{
		handlers:         make(map[EventType][]EventCallback),
		validationTiming: config.ValidationTiming,
		enabled:          true,
	}
	if config.EnableTracking {
		handler.tracker = NewEventTracker()
	}
	return handler
}
// ═══════════════════════════════════════════════════════════════════════════
// Event Handling Methods
// ═══════════════════════════════════════════════════════════════════════════
// OnChange handles field value changes
func (h *EventHandler) OnChange(ctx context.Context, event *Event) error {
	if !h.enabled {
		return nil
	}
	if event.FieldName == "" {
		return fmt.Errorf("field name is required for change event")
	}
	// Track event if tracker is enabled
	if h.tracker != nil {
		h.tracker.TrackEvent(event)
	}
	// Trigger registered callbacks
	return h.triggerCallbacks(ctx, EventChange, event)
}
// OnBlur handles field blur (focus lost)
func (h *EventHandler) OnBlur(ctx context.Context, event *Event) error {
	if !h.enabled {
		return nil
	}
	if event.FieldName == "" {
		return fmt.Errorf("field name is required for blur event")
	}
	// Track event if tracker is enabled
	if h.tracker != nil {
		h.tracker.TrackEvent(event)
	}
	// Trigger registered callbacks
	return h.triggerCallbacks(ctx, EventBlur, event)
}
// OnFocus handles field focus (focus gained)
func (h *EventHandler) OnFocus(ctx context.Context, event *Event) error {
	if !h.enabled {
		return nil
	}
	if event.FieldName == "" {
		return fmt.Errorf("field name is required for focus event")
	}
	// Track event if tracker is enabled
	if h.tracker != nil {
		h.tracker.TrackEvent(event)
	}
	// Trigger registered callbacks
	return h.triggerCallbacks(ctx, EventFocus, event)
}
// OnSubmit handles form submission
func (h *EventHandler) OnSubmit(ctx context.Context) error {
	if !h.enabled {
		return nil
	}
	event := &Event{
		Type:      EventSubmit,
		Timestamp: time.Now(),
	}
	// Track event if tracker is enabled
	if h.tracker != nil {
		h.tracker.TrackEvent(event)
	}
	// Trigger registered callbacks
	return h.triggerCallbacks(ctx, EventSubmit, event)
}
// OnReset handles form reset
func (h *EventHandler) OnReset(ctx context.Context) error {
	if !h.enabled {
		return nil
	}
	event := &Event{
		Type:      EventReset,
		Timestamp: time.Now(),
	}
	// Track event if tracker is enabled
	if h.tracker != nil {
		h.tracker.TrackEvent(event)
	}
	// Trigger registered callbacks
	return h.triggerCallbacks(ctx, EventReset, event)
}
// OnInit handles form initialization
func (h *EventHandler) OnInit(ctx context.Context) error {
	if !h.enabled {
		return nil
	}
	event := &Event{
		Type:      EventInit,
		Timestamp: time.Now(),
	}
	// Track event if tracker is enabled
	if h.tracker != nil {
		h.tracker.TrackEvent(event)
	}
	// Trigger registered callbacks
	return h.triggerCallbacks(ctx, EventInit, event)
}
// ═══════════════════════════════════════════════════════════════════════════
// Event Registration and Configuration
// ═══════════════════════════════════════════════════════════════════════════
// Register adds a custom event handler
func (h *EventHandler) Register(eventType EventType, callback EventCallback) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.handlers[eventType] = append(h.handlers[eventType], callback)
}
// RegisterMultiple registers a callback for multiple event types
func (h *EventHandler) RegisterMultiple(eventTypes []EventType, callback EventCallback) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, eventType := range eventTypes {
		h.handlers[eventType] = append(h.handlers[eventType], callback)
	}
}
// Unregister removes all event handlers for a specific event type
func (h *EventHandler) Unregister(eventType EventType) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.handlers, eventType)
}
// UnregisterAll removes all event handlers
func (h *EventHandler) UnregisterAll() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.handlers = make(map[EventType][]EventCallback)
}
// SetValidationTiming configures when validation occurs
func (h *EventHandler) SetValidationTiming(timing ValidationTiming) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.validationTiming = timing
}
// GetValidationTiming returns current validation timing strategy
func (h *EventHandler) GetValidationTiming() ValidationTiming {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.validationTiming
}
// ShouldValidateOn checks if validation should occur for given event type
func (h *EventHandler) ShouldValidateOn(eventType EventType) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch h.validationTiming {
	case ValidateOnChange:
		return eventType == EventChange
	case ValidateOnBlur:
		return eventType == EventBlur
	case ValidateOnSubmit:
		return eventType == EventSubmit
	case ValidateNever:
		return false
	default:
		return false
	}
}
// Enable enables event handling
func (h *EventHandler) Enable() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.enabled = true
}
// Disable disables event handling
func (h *EventHandler) Disable() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.enabled = false
}
// IsEnabled checks if event handling is enabled
func (h *EventHandler) IsEnabled() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.enabled
}
// GetTracker returns the event tracker (if enabled)
func (h *EventHandler) GetTracker() *EventTracker {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.tracker
}
// EnableTracking enables event tracking
func (h *EventHandler) EnableTracking() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.tracker == nil {
		h.tracker = NewEventTracker()
	}
}
// DisableTracking disables event tracking
func (h *EventHandler) DisableTracking() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.tracker = nil
}
// HasHandlers checks if there are any handlers registered for an event type
func (h *EventHandler) HasHandlers(eventType EventType) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.handlers[eventType]) > 0
}
// GetHandlerCount returns the number of handlers for an event type
func (h *EventHandler) GetHandlerCount(eventType EventType) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.handlers[eventType])
}
// ═══════════════════════════════════════════════════════════════════════════
// Private Helper Methods
// ═══════════════════════════════════════════════════════════════════════════
// triggerCallbacks executes all registered callbacks for an event type
func (h *EventHandler) triggerCallbacks(ctx context.Context, eventType EventType, event *Event) error {
	h.mu.RLock()
	callbacks := make([]EventCallback, len(h.handlers[eventType]))
	copy(callbacks, h.handlers[eventType])
	h.mu.RUnlock()
	// Execute callbacks outside of lock to prevent deadlocks
	for i, callback := range callbacks {
		if err := callback(ctx, event); err != nil {
			return fmt.Errorf("event callback %d failed: %w", i, err)
		}
	}
	return nil
}
// ═══════════════════════════════════════════════════════════════════════════
// Event Tracker
// ═══════════════════════════════════════════════════════════════════════════
// EventTracker tracks event statistics
type EventTracker struct {
	stats *EventStats
	mu    sync.RWMutex
}
// NewEventTracker creates a new event tracker
func NewEventTracker() *EventTracker {
	return &EventTracker{
		stats: &EventStats{
			EventsByType:  make(map[string]int),
			EventsByField: make(map[string]int),
		},
	}
}
// TrackEvent records an event for statistics
func (t *EventTracker) TrackEvent(event *Event) {
	t.mu.Lock()
	defer t.mu.Unlock()
	// First event
	// Track first event time using LastEventTime when TotalEvents is 0
	if t.stats.TotalEvents == 0 {
		t.stats.LastEventTime = event.Timestamp
	}
	t.stats.TotalEvents++
	t.stats.EventsByType[string(event.Type)]++
	if event.FieldName != "" {
		t.stats.EventsByField[event.FieldName]++
	}
	t.stats.LastEventTime = event.Timestamp
	t.stats.LastEventType = string(event.Type)
	t.stats.LastEventField = event.FieldName
}
// GetStats returns current event statistics
func (t *EventTracker) GetStats() *EventStats {
	t.mu.RLock()
	defer t.mu.RUnlock()
	// Return a deep copy
	statsCopy := &EventStats{
		TotalEvents:    t.stats.TotalEvents,
		EventsByType:   make(map[string]int),
		EventsByField:  make(map[string]int),
		LastEventTime:  t.stats.LastEventTime,
		LastEventType:  t.stats.LastEventType,
		LastEventField: t.stats.LastEventField,
	}
	for eventType, count := range t.stats.EventsByType {
		statsCopy.EventsByType[eventType] = count
	}
	for field, count := range t.stats.EventsByField {
		statsCopy.EventsByField[field] = count
	}
	return statsCopy
}
// Reset clears all statistics
func (t *EventTracker) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.stats = &EventStats{
		EventsByType:  make(map[string]int),
		EventsByField: make(map[string]int),
	}
}
// ═══════════════════════════════════════════════════════════════════════════
// Debounced Event Handler
// ═══════════════════════════════════════════════════════════════════════════
// DebouncedEventHandler wraps events with debouncing for better UX
type DebouncedEventHandler struct {
	handler *EventHandler
	timers  map[string]*time.Timer
	delays  map[EventType]time.Duration
	mu      sync.Mutex
}
// NewDebouncedEventHandler creates a debounced event handler
func NewDebouncedEventHandler(handler *EventHandler, config *DebouncedConfig) *DebouncedEventHandler {
	if config == nil {
		config = DefaultDebouncedConfig()
	}
	return &DebouncedEventHandler{
		handler: handler,
		timers:  make(map[string]*time.Timer),
		delays: map[EventType]time.Duration{
			EventChange: config.ChangeDelay,
			EventBlur:   config.BlurDelay,
			EventFocus:  config.FocusDelay,
		},
	}
}
// OnChange handles change events with debouncing
func (d *DebouncedEventHandler) OnChange(ctx context.Context, event *Event) error {
	return d.debounce(ctx, EventChange, event)
}
// OnBlur handles blur events with debouncing
func (d *DebouncedEventHandler) OnBlur(ctx context.Context, event *Event) error {
	return d.debounce(ctx, EventBlur, event)
}
// OnFocus handles focus events (may not be debounced)
func (d *DebouncedEventHandler) OnFocus(ctx context.Context, event *Event) error {
	delay := d.delays[EventFocus]
	if delay == 0 {
		// No debouncing for focus
		return d.handler.OnFocus(ctx, event)
	}
	return d.debounce(ctx, EventFocus, event)
}
// OnSubmit handles submit events (not debounced)
func (d *DebouncedEventHandler) OnSubmit(ctx context.Context) error {
	return d.handler.OnSubmit(ctx)
}
// debounce applies debouncing to an event
func (d *DebouncedEventHandler) debounce(ctx context.Context, eventType EventType, event *Event) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Create unique key for this field + event type
	key := fmt.Sprintf("%s:%s", eventType, event.FieldName)
	// Cancel existing timer for this field
	if timer, exists := d.timers[key]; exists {
		timer.Stop()
	}
	// Get delay for this event type
	delay := d.delays[eventType]
	if delay == 0 {
		// No debouncing - execute immediately
		delete(d.timers, key)
		return d.triggerEvent(ctx, eventType, event)
	}
	// Create new timer
	d.timers[key] = time.AfterFunc(delay, func() {
		// Execute the actual handler after delay
		if err := d.triggerEvent(ctx, eventType, event); err != nil {
			// In production, you might want to send this to an error handler
		}
		// Clean up timer
		d.mu.Lock()
		delete(d.timers, key)
		d.mu.Unlock()
	})
	return nil
}
// triggerEvent triggers the appropriate event handler
func (d *DebouncedEventHandler) triggerEvent(ctx context.Context, eventType EventType, event *Event) error {
	switch eventType {
	case EventChange:
		return d.handler.OnChange(ctx, event)
	case EventBlur:
		return d.handler.OnBlur(ctx, event)
	case EventFocus:
		return d.handler.OnFocus(ctx, event)
	default:
		return fmt.Errorf("unsupported event type: %s", eventType)
	}
}
// CancelAll cancels all pending debounced events
func (d *DebouncedEventHandler) CancelAll() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, timer := range d.timers {
		timer.Stop()
	}
	d.timers = make(map[string]*time.Timer)
}
// CancelField cancels all pending debounced events for a specific field
func (d *DebouncedEventHandler) CancelField(fieldName string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for key, timer := range d.timers {
		if len(key) > len(fieldName) && key[len(key)-len(fieldName):] == fieldName {
			timer.Stop()
			delete(d.timers, key)
		}
	}
}
// GetPendingCount returns the number of pending debounced events
func (d *DebouncedEventHandler) GetPendingCount() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.timers)
}
// ═══════════════════════════════════════════════════════════════════════════
// Batch Event Handler
// ═══════════════════════════════════════════════════════════════════════════
// BatchEventHandler handles multiple events efficiently
type BatchEventHandler struct {
	handler      *EventHandler
	batchSize    int
	batchTimeout time.Duration
	queue        []*Event
	mu           sync.Mutex
	timer        *time.Timer
	onFlush      func(context.Context, []*Event) error
}
// NewBatchEventHandler creates a batch event handler
func NewBatchEventHandler(handler *EventHandler, batchSize int, batchTimeout time.Duration) *BatchEventHandler {
	return &BatchEventHandler{
		handler:      handler,
		batchSize:    batchSize,
		batchTimeout: batchTimeout,
		queue:        make([]*Event, 0, batchSize),
	}
}
// Add adds an event to the batch queue
func (b *BatchEventHandler) Add(event *Event) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.queue = append(b.queue, event)
	// Reset timeout timer
	if b.timer != nil {
		b.timer.Stop()
	}
	b.timer = time.AfterFunc(b.batchTimeout, func() {
		b.Flush(context.Background())
	})
	// Flush if batch is full
	if len(b.queue) >= b.batchSize {
		b.flushUnsafe(context.Background())
	}
}
// Flush processes all queued events
func (b *BatchEventHandler) Flush(ctx context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.flushUnsafe(ctx)
}
// flushUnsafe flushes without acquiring lock (must be called with lock held)
func (b *BatchEventHandler) flushUnsafe(ctx context.Context) error {
	if len(b.queue) == 0 {
		return nil
	}
	// Stop timer
	if b.timer != nil {
		b.timer.Stop()
		b.timer = nil
	}
	// Process events
	events := make([]*Event, len(b.queue))
	copy(events, b.queue)
	b.queue = b.queue[:0] // Clear queue
	// Call flush callback if set
	if b.onFlush != nil {
		return b.onFlush(ctx, events)
	}
	// Default: process each event individually
	for _, event := range events {
		var err error
		switch event.Type {
		case EventChange:
			err = b.handler.OnChange(ctx, event)
		case EventBlur:
			err = b.handler.OnBlur(ctx, event)
		case EventFocus:
			err = b.handler.OnFocus(ctx, event)
		case EventSubmit:
			err = b.handler.OnSubmit(ctx)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
// SetFlushCallback sets a custom callback for batch processing
func (b *BatchEventHandler) SetFlushCallback(callback func(context.Context, []*Event) error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.onFlush = callback
}
// GetQueueSize returns the current queue size
func (b *BatchEventHandler) GetQueueSize() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.queue)
}

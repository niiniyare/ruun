package validation

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// PerformanceValidator validates performance characteristics
type PerformanceValidator struct {
	config     PerformanceConfig
	metrics    *MetricsCollector
	thresholds PerformanceThresholds
}

// PerformanceConfig configures performance validation
type PerformanceConfig struct {
	EnableBundleSize     bool                  `json:"enableBundleSize"`
	EnableRenderTime     bool                  `json:"enableRenderTime"`
	EnableMemoryUsage    bool                  `json:"enableMemoryUsage"`
	EnableComplexity     bool                  `json:"enableComplexity"`
	EnableNetworkMetrics bool                  `json:"enableNetworkMetrics"`
	SamplingInterval     time.Duration         `json:"samplingInterval"`
	MaxSamples           int                   `json:"maxSamples"`
	Thresholds           PerformanceThresholds `json:"thresholds"`
}

// PerformanceThresholds defines performance thresholds
type PerformanceThresholds struct {
	MaxBundleSize      int64         `json:"maxBundleSize"`  // bytes
	MaxRenderTime      time.Duration `json:"maxRenderTime"`  // milliseconds
	MaxMemoryUsage     int64         `json:"maxMemoryUsage"` // bytes
	MaxComplexity      int           `json:"maxComplexity"`  // complexity score
	MaxNetworkRequests int           `json:"maxNetworkRequests"`
	MaxPayloadSize     int64         `json:"maxPayloadSize"`  // bytes
	MaxResponseTime    time.Duration `json:"maxResponseTime"` // milliseconds
	MinFrameRate       float64       `json:"minFrameRate"`    // FPS
	MaxCPUUsage        float64       `json:"maxCPUUsage"`     // percentage
}

// PerformanceMetrics contains performance measurements
type PerformanceMetrics struct {
	BundleSize      *BundleSizeMetrics `json:"bundleSize,omitempty"`
	RenderTime      *RenderTimeMetrics `json:"renderTime,omitempty"`
	MemoryUsage     *MemoryMetrics     `json:"memoryUsage,omitempty"`
	Complexity      *ComplexityMetrics `json:"complexity,omitempty"`
	Network         *NetworkMetrics    `json:"network,omitempty"`
	Runtime         *RuntimeMetrics    `json:"runtime,omitempty"`
	MeetsThresholds bool               `json:"meetsThresholds"`
	Score           float64            `json:"score"` // 0-100
	Grade           PerformanceGrade   `json:"grade"`
	Timestamp       time.Time          `json:"timestamp"`
}

// BundleSizeMetrics contains bundle size measurements
type BundleSizeMetrics struct {
	TotalSize       int64            `json:"totalSize"`
	GzippedSize     int64            `json:"gzippedSize"`
	ComponentSizes  map[string]int64 `json:"componentSizes"`
	DependencySizes map[string]int64 `json:"dependencySizes"`
	ThresholdsMet   bool             `json:"thresholdsMet"`
	Recommendations []string         `json:"recommendations,omitempty"`
}

// RenderTimeMetrics contains render performance measurements
type RenderTimeMetrics struct {
	InitialRender  time.Duration            `json:"initialRender"`
	ReRenders      []time.Duration          `json:"reRenders,omitempty"`
	AverageRender  time.Duration            `json:"averageRender"`
	P95Render      time.Duration            `json:"p95Render"`
	P99Render      time.Duration            `json:"p99Render"`
	FrameRate      float64                  `json:"frameRate"`
	ComponentTimes map[string]time.Duration `json:"componentTimes"`
	ThresholdsMet  bool                     `json:"thresholdsMet"`
	Bottlenecks    []string                 `json:"bottlenecks,omitempty"`
}

// MemoryMetrics contains memory usage measurements
type MemoryMetrics struct {
	HeapUsage       int64            `json:"heapUsage"`
	HeapSize        int64            `json:"heapSize"`
	ObjectCount     int64            `json:"objectCount"`
	GCCount         int64            `json:"gcCount"`
	ComponentMemory map[string]int64 `json:"componentMemory"`
	LeakDetected    bool             `json:"leakDetected"`
	ThresholdsMet   bool             `json:"thresholdsMet"`
	Warnings        []string         `json:"warnings,omitempty"`
}

// ComplexityMetrics contains complexity measurements
type ComplexityMetrics struct {
	CyclomaticComplexity int                    `json:"cyclomaticComplexity"`
	CognitiveComplexity  int                    `json:"cognitiveComplexity"`
	ComponentComplexity  map[string]int         `json:"componentComplexity"`
	DependencyDepth      int                    `json:"dependencyDepth"`
	PropComplexity       int                    `json:"propComplexity"`
	StateComplexity      int                    `json:"stateComplexity"`
	ThresholdsMet        bool                   `json:"thresholdsMet"`
	Suggestions          []ComplexitySuggestion `json:"suggestions,omitempty"`
}

// NetworkMetrics contains network performance measurements
type NetworkMetrics struct {
	RequestCount   int                       `json:"requestCount"`
	TotalPayload   int64                     `json:"totalPayload"`
	ResponseTimes  []time.Duration           `json:"responseTimes"`
	FailedRequests int                       `json:"failedRequests"`
	CacheHitRate   float64                   `json:"cacheHitRate"`
	RequestDetails map[string]RequestMetrics `json:"requestDetails"`
	ThresholdsMet  bool                      `json:"thresholdsMet"`
	Optimizations  []string                  `json:"optimizations,omitempty"`
}

// RuntimeMetrics contains runtime performance measurements
type RuntimeMetrics struct {
	CPUUsage       float64                  `json:"cpuUsage"`
	MemoryUsage    int64                    `json:"memoryUsage"`
	GoroutineCount int                      `json:"goroutineCount"`
	GCPauses       []time.Duration          `json:"gcPauses"`
	ResponseTimes  map[string]time.Duration `json:"responseTimes"`
	ThroughputRPS  float64                  `json:"throughputRPS"`
	ThresholdsMet  bool                     `json:"thresholdsMet"`
	Alerts         []string                 `json:"alerts,omitempty"`
}

// RequestMetrics contains individual request metrics
type RequestMetrics struct {
	URL          string        `json:"url"`
	Method       string        `json:"method"`
	ResponseTime time.Duration `json:"responseTime"`
	PayloadSize  int64         `json:"payloadSize"`
	StatusCode   int           `json:"statusCode"`
	Cached       bool          `json:"cached"`
}

// ComplexitySuggestion contains suggestions for reducing complexity
type ComplexitySuggestion struct {
	Component  string `json:"component"`
	Issue      string `json:"issue"`
	Suggestion string `json:"suggestion"`
	Impact     string `json:"impact"` // high, medium, low
	AutoFix    bool   `json:"autoFix"`
}

// PerformanceGrade represents performance grade
type PerformanceGrade string

const (
	PerformanceGradeA PerformanceGrade = "A"
	PerformanceGradeB PerformanceGrade = "B"
	PerformanceGradeC PerformanceGrade = "C"
	PerformanceGradeD PerformanceGrade = "D"
	PerformanceGradeF PerformanceGrade = "F"
)

// MetricsCollector collects and analyzes performance metrics
type MetricsCollector struct {
	samples         []PerformanceMetrics
	config          PerformanceConfig
	startTime       time.Time
	renderTimes     []time.Duration
	memorySnapshots []runtime.MemStats
}

// NewPerformanceValidator creates a new performance validator
func NewPerformanceValidator() *PerformanceValidator {
	config := PerformanceConfig{
		EnableBundleSize:     true,
		EnableRenderTime:     true,
		EnableMemoryUsage:    true,
		EnableComplexity:     true,
		EnableNetworkMetrics: true,
		SamplingInterval:     time.Millisecond * 100,
		MaxSamples:           1000,
		Thresholds: PerformanceThresholds{
			MaxBundleSize:      5 * 1024 * 1024, // 5MB
			MaxRenderTime:      time.Millisecond * 100,
			MaxMemoryUsage:     100 * 1024 * 1024, // 100MB
			MaxComplexity:      50,
			MaxNetworkRequests: 20,
			MaxPayloadSize:     1024 * 1024, // 1MB
			MaxResponseTime:    time.Millisecond * 500,
			MinFrameRate:       60.0,
			MaxCPUUsage:        80.0,
		},
	}

	return &PerformanceValidator{
		config:     config,
		metrics:    NewMetricsCollector(config),
		thresholds: config.Thresholds,
	}
}

// NewMetricsCollector creates a new metrics collector
func NewMetricsCollector(config PerformanceConfig) *MetricsCollector {
	return &MetricsCollector{
		config:          config,
		startTime:       time.Now(),
		samples:         make([]PerformanceMetrics, 0),
		renderTimes:     make([]time.Duration, 0),
		memorySnapshots: make([]runtime.MemStats, 0),
	}
}

// ValidatePerformance validates performance characteristics
func (pv *PerformanceValidator) ValidatePerformance(ctx *ValidationContext, value any) *PerformanceMetrics {
	start := time.Now()

	metrics := &PerformanceMetrics{
		MeetsThresholds: true,
		Timestamp:       start,
	}

	// Collect bundle size metrics
	if pv.config.EnableBundleSize {
		metrics.BundleSize = pv.collectBundleSizeMetrics(value)
		if !metrics.BundleSize.ThresholdsMet {
			metrics.MeetsThresholds = false
		}
	}

	// Collect render time metrics
	if pv.config.EnableRenderTime {
		metrics.RenderTime = pv.collectRenderTimeMetrics(value)
		if !metrics.RenderTime.ThresholdsMet {
			metrics.MeetsThresholds = false
		}
	}

	// Collect memory usage metrics
	if pv.config.EnableMemoryUsage {
		metrics.MemoryUsage = pv.collectMemoryMetrics(value)
		if !metrics.MemoryUsage.ThresholdsMet {
			metrics.MeetsThresholds = false
		}
	}

	// Collect complexity metrics
	if pv.config.EnableComplexity {
		metrics.Complexity = pv.collectComplexityMetrics(value)
		if !metrics.Complexity.ThresholdsMet {
			metrics.MeetsThresholds = false
		}
	}

	// Collect network metrics
	if pv.config.EnableNetworkMetrics {
		metrics.Network = pv.collectNetworkMetrics(value)
		if !metrics.Network.ThresholdsMet {
			metrics.MeetsThresholds = false
		}
	}

	// Collect runtime metrics
	metrics.Runtime = pv.collectRuntimeMetrics(value)
	if !metrics.Runtime.ThresholdsMet {
		metrics.MeetsThresholds = false
	}

	// Calculate overall score and grade
	metrics.Score = pv.calculatePerformanceScore(metrics)
	metrics.Grade = pv.calculatePerformanceGrade(metrics.Score)

	// Store metrics
	pv.metrics.AddSample(*metrics)

	return metrics
}

// collectBundleSizeMetrics collects bundle size metrics
func (pv *PerformanceValidator) collectBundleSizeMetrics(value any) *BundleSizeMetrics {
	metrics := &BundleSizeMetrics{
		ComponentSizes:  make(map[string]int64),
		DependencySizes: make(map[string]int64),
		Recommendations: make([]string, 0),
	}

	// Calculate estimated component size
	size := pv.estimateComponentSize(value)
	metrics.TotalSize = size
	metrics.GzippedSize = int64(float64(size) * 0.7) // Estimate 70% compression

	// Check thresholds
	if metrics.TotalSize <= pv.thresholds.MaxBundleSize {
		metrics.ThresholdsMet = true
	} else {
		metrics.ThresholdsMet = false
		metrics.Recommendations = append(metrics.Recommendations,
			fmt.Sprintf("Bundle size %d bytes exceeds threshold %d bytes",
				metrics.TotalSize, pv.thresholds.MaxBundleSize))
	}

	// Analyze component sizes
	if componentMap, ok := value.(map[string]any); ok {
		if components, exists := componentMap["components"]; exists {
			if compArray, ok := components.([]any); ok {
				for i, comp := range compArray {
					compSize := pv.estimateComponentSize(comp)
					compName := fmt.Sprintf("component_%d", i)
					if compMap, ok := comp.(map[string]any); ok {
						if name, exists := compMap["type"]; exists {
							compName = fmt.Sprintf("%v", name)
						}
					}
					metrics.ComponentSizes[compName] = compSize
				}
			}
		}
	}

	return metrics
}

// collectRenderTimeMetrics collects render performance metrics
func (pv *PerformanceValidator) collectRenderTimeMetrics(value any) *RenderTimeMetrics {
	metrics := &RenderTimeMetrics{
		ComponentTimes: make(map[string]time.Duration),
		Bottlenecks:    make([]string, 0),
	}

	// Simulate render timing
	renderStart := time.Now()
	pv.simulateRender(value)
	renderTime := time.Since(renderStart)

	metrics.InitialRender = renderTime
	metrics.AverageRender = renderTime
	metrics.FrameRate = 1000.0 / float64(renderTime.Milliseconds())

	// Add to render times for statistics
	pv.metrics.renderTimes = append(pv.metrics.renderTimes, renderTime)

	// Calculate percentiles if we have enough samples
	if len(pv.metrics.renderTimes) > 5 {
		metrics.P95Render = pv.calculatePercentile(pv.metrics.renderTimes, 95)
		metrics.P99Render = pv.calculatePercentile(pv.metrics.renderTimes, 99)
	}

	// Check thresholds
	if renderTime <= pv.thresholds.MaxRenderTime {
		metrics.ThresholdsMet = true
	} else {
		metrics.ThresholdsMet = false
		metrics.Bottlenecks = append(metrics.Bottlenecks,
			fmt.Sprintf("Render time %v exceeds threshold %v",
				renderTime, pv.thresholds.MaxRenderTime))
	}

	if metrics.FrameRate < pv.thresholds.MinFrameRate {
		metrics.ThresholdsMet = false
		metrics.Bottlenecks = append(metrics.Bottlenecks,
			fmt.Sprintf("Frame rate %.1f FPS below threshold %.1f FPS",
				metrics.FrameRate, pv.thresholds.MinFrameRate))
	}

	return metrics
}

// collectMemoryMetrics collects memory usage metrics
func (pv *PerformanceValidator) collectMemoryMetrics(value any) *MemoryMetrics {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	metrics := &MemoryMetrics{
		HeapUsage:       int64(m.HeapInuse),
		HeapSize:        int64(m.HeapSys),
		ObjectCount:     int64(m.Mallocs - m.Frees),
		GCCount:         int64(m.NumGC),
		ComponentMemory: make(map[string]int64),
		Warnings:        make([]string, 0),
	}

	// Store snapshot
	pv.metrics.memorySnapshots = append(pv.metrics.memorySnapshots, m)

	// Detect memory leaks (simplified)
	if len(pv.metrics.memorySnapshots) > 10 {
		oldSnapshot := pv.metrics.memorySnapshots[len(pv.metrics.memorySnapshots)-10]
		if m.HeapInuse > oldSnapshot.HeapInuse*2 {
			metrics.LeakDetected = true
			metrics.Warnings = append(metrics.Warnings, "Potential memory leak detected")
		}
	}

	// Check thresholds
	if metrics.HeapUsage <= pv.thresholds.MaxMemoryUsage {
		metrics.ThresholdsMet = true
	} else {
		metrics.ThresholdsMet = false
		metrics.Warnings = append(metrics.Warnings,
			fmt.Sprintf("Memory usage %d bytes exceeds threshold %d bytes",
				metrics.HeapUsage, pv.thresholds.MaxMemoryUsage))
	}

	return metrics
}

// collectComplexityMetrics collects complexity metrics
func (pv *PerformanceValidator) collectComplexityMetrics(value any) *ComplexityMetrics {
	metrics := &ComplexityMetrics{
		ComponentComplexity: make(map[string]int),
		Suggestions:         make([]ComplexitySuggestion, 0),
	}

	// Calculate complexity scores
	complexity := pv.calculateComplexity(value)
	metrics.CyclomaticComplexity = complexity.cyclomatic
	metrics.CognitiveComplexity = complexity.cognitive
	metrics.PropComplexity = complexity.props
	metrics.StateComplexity = complexity.state
	metrics.DependencyDepth = complexity.depth

	totalComplexity := complexity.cyclomatic + complexity.cognitive + complexity.props + complexity.state

	// Check thresholds
	if totalComplexity <= pv.thresholds.MaxComplexity {
		metrics.ThresholdsMet = true
	} else {
		metrics.ThresholdsMet = false

		if complexity.cyclomatic > 10 {
			metrics.Suggestions = append(metrics.Suggestions, ComplexitySuggestion{
				Component:  "main",
				Issue:      "High cyclomatic complexity",
				Suggestion: "Break down complex functions into smaller ones",
				Impact:     "high",
				AutoFix:    false,
			})
		}

		if complexity.props > 10 {
			metrics.Suggestions = append(metrics.Suggestions, ComplexitySuggestion{
				Component:  "main",
				Issue:      "Too many props",
				Suggestion: "Group related props into objects or use composition",
				Impact:     "medium",
				AutoFix:    false,
			})
		}

		if complexity.depth > 5 {
			metrics.Suggestions = append(metrics.Suggestions, ComplexitySuggestion{
				Component:  "main",
				Issue:      "Deep dependency tree",
				Suggestion: "Flatten component hierarchy",
				Impact:     "medium",
				AutoFix:    false,
			})
		}
	}

	return metrics
}

// collectNetworkMetrics collects network performance metrics
func (pv *PerformanceValidator) collectNetworkMetrics(value any) *NetworkMetrics {
	metrics := &NetworkMetrics{
		RequestDetails: make(map[string]RequestMetrics),
		Optimizations:  make([]string, 0),
	}

	// Simulate network metrics based on component structure
	requestCount := pv.estimateNetworkRequests(value)
	metrics.RequestCount = requestCount
	metrics.TotalPayload = int64(requestCount) * 1024 // Estimate 1KB per request
	metrics.CacheHitRate = 0.8                        // Assume 80% cache hit rate

	// Check thresholds
	meetsThresholds := true

	if metrics.RequestCount > pv.thresholds.MaxNetworkRequests {
		meetsThresholds = false
		metrics.Optimizations = append(metrics.Optimizations,
			fmt.Sprintf("Reduce network requests from %d to %d or fewer",
				metrics.RequestCount, pv.thresholds.MaxNetworkRequests))
	}

	if metrics.TotalPayload > pv.thresholds.MaxPayloadSize {
		meetsThresholds = false
		metrics.Optimizations = append(metrics.Optimizations,
			fmt.Sprintf("Reduce total payload from %d to %d bytes or fewer",
				metrics.TotalPayload, pv.thresholds.MaxPayloadSize))
	}

	metrics.ThresholdsMet = meetsThresholds

	return metrics
}

// collectRuntimeMetrics collects runtime performance metrics
func (pv *PerformanceValidator) collectRuntimeMetrics(value any) *RuntimeMetrics {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	metrics := &RuntimeMetrics{
		MemoryUsage:    int64(m.Sys),
		GoroutineCount: runtime.NumGoroutine(),
		ResponseTimes:  make(map[string]time.Duration),
		Alerts:         make([]string, 0),
	}

	// Simulate CPU usage (in real implementation, this would use actual monitoring)
	metrics.CPUUsage = 45.0 + float64(time.Now().UnixNano()%20) // Simulate 45-65% usage

	// Simulate throughput
	metrics.ThroughputRPS = 100.0 + float64(time.Now().UnixNano()%50) // Simulate 100-150 RPS

	// Check thresholds
	meetsThresholds := true

	if metrics.CPUUsage > pv.thresholds.MaxCPUUsage {
		meetsThresholds = false
		metrics.Alerts = append(metrics.Alerts,
			fmt.Sprintf("CPU usage %.1f%% exceeds threshold %.1f%%",
				metrics.CPUUsage, pv.thresholds.MaxCPUUsage))
	}

	if metrics.MemoryUsage > pv.thresholds.MaxMemoryUsage {
		meetsThresholds = false
		metrics.Alerts = append(metrics.Alerts,
			fmt.Sprintf("Memory usage %d bytes exceeds threshold %d bytes",
				metrics.MemoryUsage, pv.thresholds.MaxMemoryUsage))
	}

	metrics.ThresholdsMet = meetsThresholds

	return metrics
}

// Helper methods

type complexityScores struct {
	cyclomatic int
	cognitive  int
	props      int
	state      int
	depth      int
}

func (pv *PerformanceValidator) estimateComponentSize(value any) int64 {
	// Simplified size estimation based on structure
	if value == nil {
		return 0
	}

	size := int64(100) // Base size

	if valueMap, ok := value.(map[string]any); ok {
		for key, val := range valueMap {
			size += int64(len(key)) * 2 // Key size
			size += pv.estimateValueSize(val)
		}
	} else if valueSlice, ok := value.([]any); ok {
		for _, val := range valueSlice {
			size += pv.estimateValueSize(val)
		}
	} else {
		size += int64(len(fmt.Sprintf("%v", value)))
	}

	return size
}

func (pv *PerformanceValidator) estimateValueSize(value any) int64 {
	if value == nil {
		return 4
	}

	switch v := value.(type) {
	case string:
		return int64(len(v) * 2) // Unicode overhead
	case map[string]any:
		return pv.estimateComponentSize(v)
	case []any:
		size := int64(0)
		for _, item := range v {
			size += pv.estimateValueSize(item)
		}
		return size
	default:
		return int64(len(fmt.Sprintf("%v", v)))
	}
}

func (pv *PerformanceValidator) simulateRender(value any) {
	// Simulate render work based on component complexity
	complexity := pv.calculateComplexity(value)
	workDuration := time.Microsecond * time.Duration(complexity.cyclomatic*10)
	time.Sleep(workDuration)
}

func (pv *PerformanceValidator) calculateComplexity(value any) complexityScores {
	scores := complexityScores{}

	if valueMap, ok := value.(map[string]any); ok {
		// Count props
		if props, exists := valueMap["props"]; exists {
			if propsMap, ok := props.(map[string]any); ok {
				scores.props = len(propsMap)
			}
		}

		// Count state complexity
		if state, exists := valueMap["state"]; exists {
			if stateMap, ok := state.(map[string]any); ok {
				scores.state = len(stateMap)
			}
		}

		// Calculate depth
		scores.depth = pv.calculateDepth(valueMap, 0)

		// Estimate cyclomatic complexity
		scores.cyclomatic = scores.props + scores.state + scores.depth

		// Estimate cognitive complexity
		scores.cognitive = scores.cyclomatic + (scores.depth * 2)
	}

	return scores
}

func (pv *PerformanceValidator) calculateDepth(value map[string]any, currentDepth int) int {
	maxDepth := currentDepth

	for _, val := range value {
		if childMap, ok := val.(map[string]any); ok {
			depth := pv.calculateDepth(childMap, currentDepth+1)
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if childSlice, ok := val.([]any); ok {
			for _, child := range childSlice {
				if childMap, ok := child.(map[string]any); ok {
					depth := pv.calculateDepth(childMap, currentDepth+1)
					if depth > maxDepth {
						maxDepth = depth
					}
				}
			}
		}
	}

	return maxDepth
}

func (pv *PerformanceValidator) estimateNetworkRequests(value any) int {
	requests := 0

	if valueMap, ok := value.(map[string]any); ok {
		// Check for API endpoints
		for key, val := range valueMap {
			if strings.Contains(strings.ToLower(key), "api") ||
				strings.Contains(strings.ToLower(key), "url") ||
				strings.Contains(strings.ToLower(key), "endpoint") {
				requests++
			}

			if valStr, ok := val.(string); ok {
				if strings.HasPrefix(valStr, "http") {
					requests++
				}
			}
		}

		// Recursively check nested structures
		if components, exists := valueMap["components"]; exists {
			if compSlice, ok := components.([]any); ok {
				for _, comp := range compSlice {
					requests += pv.estimateNetworkRequests(comp)
				}
			}
		}
	}

	return requests
}

func (pv *PerformanceValidator) calculatePercentile(values []time.Duration, percentile float64) time.Duration {
	if len(values) == 0 {
		return 0
	}

	// Sort values (simplified - in production use proper sorting)
	sorted := make([]time.Duration, len(values))
	copy(sorted, values)

	// Simple bubble sort for demonstration
	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted)-1-i; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	index := int(float64(len(sorted)) * percentile / 100.0)
	if index >= len(sorted) {
		index = len(sorted) - 1
	}

	return sorted[index]
}

func (pv *PerformanceValidator) calculatePerformanceScore(metrics *PerformanceMetrics) float64 {
	score := 100.0

	// Penalize threshold violations
	if metrics.BundleSize != nil && !metrics.BundleSize.ThresholdsMet {
		score -= 15.0
	}
	if metrics.RenderTime != nil && !metrics.RenderTime.ThresholdsMet {
		score -= 20.0
	}
	if metrics.MemoryUsage != nil && !metrics.MemoryUsage.ThresholdsMet {
		score -= 15.0
	}
	if metrics.Complexity != nil && !metrics.Complexity.ThresholdsMet {
		score -= 10.0
	}
	if metrics.Network != nil && !metrics.Network.ThresholdsMet {
		score -= 10.0
	}
	if metrics.Runtime != nil && !metrics.Runtime.ThresholdsMet {
		score -= 20.0
	}

	// Additional penalties for specific metrics
	if metrics.RenderTime != nil && metrics.RenderTime.FrameRate < 30.0 {
		score -= 10.0 // Severe frame rate penalty
	}
	if metrics.MemoryUsage != nil && metrics.MemoryUsage.LeakDetected {
		score -= 25.0 // Memory leak penalty
	}

	// Ensure score is between 0 and 100
	if score < 0 {
		score = 0
	}

	return score
}

func (pv *PerformanceValidator) calculatePerformanceGrade(score float64) PerformanceGrade {
	switch {
	case score >= 90:
		return PerformanceGradeA
	case score >= 80:
		return PerformanceGradeB
	case score >= 70:
		return PerformanceGradeC
	case score >= 60:
		return PerformanceGradeD
	default:
		return PerformanceGradeF
	}
}

// AddSample adds a performance sample to the collector
func (mc *MetricsCollector) AddSample(sample PerformanceMetrics) {
	mc.samples = append(mc.samples, sample)

	// Keep only the most recent samples
	if len(mc.samples) > mc.config.MaxSamples {
		mc.samples = mc.samples[1:]
	}
}

// GetAverageMetrics returns average metrics across all samples
func (mc *MetricsCollector) GetAverageMetrics() PerformanceMetrics {
	if len(mc.samples) == 0 {
		return PerformanceMetrics{}
	}

	// Calculate averages (simplified implementation)
	var totalScore float64
	var totalRenderTime time.Duration

	for _, sample := range mc.samples {
		totalScore += sample.Score
		if sample.RenderTime != nil {
			totalRenderTime += sample.RenderTime.AverageRender
		}
	}

	avgScore := totalScore / float64(len(mc.samples))
	avgRenderTime := time.Duration(int64(totalRenderTime) / int64(len(mc.samples)))

	return PerformanceMetrics{
		Score:     avgScore,
		Grade:     mc.calculateGrade(avgScore),
		Timestamp: time.Now(),
		RenderTime: &RenderTimeMetrics{
			AverageRender: avgRenderTime,
		},
	}
}

func (mc *MetricsCollector) calculateGrade(score float64) PerformanceGrade {
	switch {
	case score >= 90:
		return PerformanceGradeA
	case score >= 80:
		return PerformanceGradeB
	case score >= 70:
		return PerformanceGradeC
	case score >= 60:
		return PerformanceGradeD
	default:
		return PerformanceGradeF
	}
}

// GetTrends returns performance trends over time
func (mc *MetricsCollector) GetTrends() map[string][]float64 {
	trends := map[string][]float64{
		"score":       make([]float64, 0),
		"renderTime":  make([]float64, 0),
		"memoryUsage": make([]float64, 0),
	}

	for _, sample := range mc.samples {
		trends["score"] = append(trends["score"], sample.Score)

		if sample.RenderTime != nil {
			trends["renderTime"] = append(trends["renderTime"], float64(sample.RenderTime.AverageRender.Milliseconds()))
		}

		if sample.MemoryUsage != nil {
			trends["memoryUsage"] = append(trends["memoryUsage"], float64(sample.MemoryUsage.HeapUsage))
		}
	}

	return trends
}

// Reset clears all collected metrics
func (mc *MetricsCollector) Reset() {
	mc.samples = make([]PerformanceMetrics, 0)
	mc.renderTimes = make([]time.Duration, 0)
	mc.memorySnapshots = make([]runtime.MemStats, 0)
	mc.startTime = time.Now()
}

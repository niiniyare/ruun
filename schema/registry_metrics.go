// pkg/schema/registry_metrics.go
package schema

import (
	"sync/atomic"
	"time"
)



// NewRegistryMetrics creates metrics
func NewRegistryMetrics() *RegistryMetrics {
	return &RegistryMetrics{
		operations: make(map[string]*operationMetrics),
	}
}

func (m *RegistryMetrics) IncrementSchemaCount() {
	atomic.AddInt64(&m.schemaCount, 1)
}

func (m *RegistryMetrics) DecrementSchemaCount() {
	atomic.AddInt64(&m.schemaCount, -1)
}

func (m *RegistryMetrics) IncrementCacheHit() {
	atomic.AddInt64(&m.cacheHits, 1)
}

func (m *RegistryMetrics) IncrementCacheMiss() {
	atomic.AddInt64(&m.cacheMisses, 1)
}

func (m *RegistryMetrics) RecordOperation(op string, duration time.Duration) {
	if _, ok := m.operations[op]; !ok {
		m.operations[op] = &operationMetrics{}
	}

	metrics := m.operations[op]
	atomic.AddInt64(&metrics.count, 1)
	atomic.AddInt64(&metrics.duration, int64(duration))
}

func (m *RegistryMetrics) GetStats() map[string]any {
	return map[string]any{
		"schema_count": atomic.LoadInt64(&m.schemaCount),
		"cache_hits":   atomic.LoadInt64(&m.cacheHits),
		"cache_misses": atomic.LoadInt64(&m.cacheMisses),
		"cache_ratio":  m.getCacheHitRatio(),
	}
}

func (m *RegistryMetrics) getCacheHitRatio() float64 {
	hits := atomic.LoadInt64(&m.cacheHits)
	misses := atomic.LoadInt64(&m.cacheMisses)
	total := hits + misses

	if total == 0 {
		return 0
	}

	return float64(hits) / float64(total)
}

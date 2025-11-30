// pkg/schema/cache_memory.go
package schema

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MemoryCache provides in-memory caching
type MemoryCache struct {
	items    map[string]*cacheItem
	maxSize  int
	mu       sync.RWMutex
	stopChan chan struct{}
}

type cacheItem struct {
	data      []byte
	expiresAt time.Time
}

// NewMemoryCache creates memory cache
func NewMemoryCache(maxSize int) *MemoryCache {
	cache := &MemoryCache{
		items:    make(map[string]*cacheItem),
		maxSize:  maxSize,
		stopChan: make(chan struct{}),
	}

	// Start cleanup goroutine
	go cache.cleanup()

	return cache
}

func (c *MemoryCache) Get(ctx context.Context, key string) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	if time.Now().After(item.expiresAt) {
		return nil, fmt.Errorf("key expired")
	}

	return item.data, nil
}

func (c *MemoryCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check size limit
	if len(c.items) >= c.maxSize {
		c.evictOldest()
	}

	c.items[key] = &cacheItem{
		data:      value,
		expiresAt: time.Now().Add(ttl),
	}

	return nil
}

func (c *MemoryCache) Delete(ctx context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	return nil
}

func (c *MemoryCache) Clear(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*cacheItem)
	return nil
}

func (c *MemoryCache) Health(ctx context.Context) error {
	return nil
}

func (c *MemoryCache) Stop() {
	close(c.stopChan)
}

func (c *MemoryCache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.removeExpired()
		case <-c.stopChan:
			return
		}
	}
}

func (c *MemoryCache) removeExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, item := range c.items {
		if now.After(item.expiresAt) {
			delete(c.items, key)
		}
	}
}

func (c *MemoryCache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, item := range c.items {
		if oldestKey == "" || item.expiresAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = item.expiresAt
		}
	}

	if oldestKey != "" {
		delete(c.items, oldestKey)
	}
}

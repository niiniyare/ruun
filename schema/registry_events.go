package schema

import (
	"sync"
	"time"
)

// RegistryEventType defines event types
type RegistryEventType string

const (
	EventSchemaRegistered RegistryEventType = "schema.registered"
	EventSchemaUpdated    RegistryEventType = "schema.updated"
	EventSchemaDeleted    RegistryEventType = "schema.deleted"
	EventSchemaAccessed   RegistryEventType = "schema.accessed"
)

// RegistryEvent represents a registry event
type RegistryEvent struct {
	Type      RegistryEventType
	SchemaID  string
	Timestamp time.Time
	Data      map[string]any
}

// RegistryEventHandler handles registry events
type RegistryEventHandler func(*RegistryEvent)

// RegistryEventBus manages event subscriptions
type RegistryEventBus struct {
	handlers []RegistryEventHandler
	mu       sync.RWMutex
}

// NewRegistryEventBus creates event bus
func NewRegistryEventBus() *RegistryEventBus {
	return &RegistryEventBus{
		handlers: []RegistryEventHandler{},
	}
}

// Subscribe adds event handler
func (b *RegistryEventBus) Subscribe(handler RegistryEventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers = append(b.handlers, handler)
}

// Emit publishes event
func (b *RegistryEventBus) Emit(event *RegistryEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, handler := range b.handlers {
		go handler(event)
	}
}

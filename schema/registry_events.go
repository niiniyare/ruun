package schema

import (
	"sync"
)




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

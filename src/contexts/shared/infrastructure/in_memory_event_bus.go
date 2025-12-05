package infrastructure

import (
	"context"
	"sync"

	"github.com/ezflow-me/identity-management-service/src/contexts/shared/domain"
)

// InMemoryEventBus is an in-memory implementation of EventBus for testing
type InMemoryEventBus struct {
	mu     sync.Mutex
	events []domain.DomainEvent
}

// NewInMemoryEventBus creates a new InMemoryEventBus
func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		events: make([]domain.DomainEvent, 0),
	}
}

// Publish stores events in memory
func (bus *InMemoryEventBus) Publish(ctx context.Context, events ...domain.DomainEvent) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.events = append(bus.events, events...)
	return nil
}

// PublishedEvents returns all published events
func (bus *InMemoryEventBus) PublishedEvents() []domain.DomainEvent {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	return bus.events
}

// Clear clears all published events
func (bus *InMemoryEventBus) Clear() {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.events = make([]domain.DomainEvent, 0)
}

// LastEvent returns the last published event
func (bus *InMemoryEventBus) LastEvent() domain.DomainEvent {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	if len(bus.events) == 0 {
		return nil
	}
	return bus.events[len(bus.events)-1]
}

// EventCount returns the number of published events
func (bus *InMemoryEventBus) EventCount() int {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	return len(bus.events)
}

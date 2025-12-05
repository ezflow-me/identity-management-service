package domain

import "context"

// EventBus defines the interface for publishing domain events
type EventBus interface {
	Publish(ctx context.Context, events ...DomainEvent) error
}

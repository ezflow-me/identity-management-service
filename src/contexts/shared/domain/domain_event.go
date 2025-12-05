package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// DomainEvent represents a domain event that occurred in the system
type DomainEvent interface {
	EventID() string
	AggregateID() string
	EventName() string
	OccurredOn() time.Time
	Topic() string
	ToPrimitives() map[string]interface{}
	ToJSON() ([]byte, error)
}

// BaseDomainEvent provides common functionality for all domain events
type BaseDomainEvent struct {
	eventID     string
	aggregateID string
	eventName   string
	occurredOn  time.Time
	topic       string
}

// NewBaseDomainEvent creates a new BaseDomainEvent
func NewBaseDomainEvent(aggregateID, eventName, topic string) BaseDomainEvent {
	return BaseDomainEvent{
		eventID:     uuid.New().String(),
		aggregateID: aggregateID,
		eventName:   eventName,
		occurredOn:  time.Now().UTC(),
		topic:       topic,
	}
}

// EventID returns the unique identifier of the event
func (e BaseDomainEvent) EventID() string {
	return e.eventID
}

// AggregateID returns the aggregate identifier
func (e BaseDomainEvent) AggregateID() string {
	return e.aggregateID
}

// EventName returns the name of the event
func (e BaseDomainEvent) EventName() string {
	return e.eventName
}

// OccurredOn returns when the event occurred
func (e BaseDomainEvent) OccurredOn() time.Time {
	return e.occurredOn
}

// Topic returns the Kafka topic for this event
func (e BaseDomainEvent) Topic() string {
	return e.topic
}

// ToJSON serializes the event metadata to JSON for Kafka
func (e BaseDomainEvent) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"event_id":     e.eventID,
		"aggregate_id": e.aggregateID,
		"event_name":   e.eventName,
		"occurred_on":  e.occurredOn.Format(time.RFC3339Nano),
	})
}

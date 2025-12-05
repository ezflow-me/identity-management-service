package domain

import (
	"encoding/json"
	"maps"
	"time"

	shared "github.com/ezflow-me/identity-management-service/src/contexts/shared/domain"
)

const (
	UserRegisteredEventName  = "user.registered"
	UserRegisteredEventTopic = "identity.users.registered"
)

type UserRegisteredEvent struct {
	shared.BaseDomainEvent
	User
}

func NewUserRegisteredEvent(user *User) UserRegisteredEvent {
	return UserRegisteredEvent{
		BaseDomainEvent: shared.NewBaseDomainEvent(user.ID(), UserRegisteredEventName, UserRegisteredEventTopic),
		User:            *user,
	}
}

func (e UserRegisteredEvent) ToPrimitives() map[string]interface{} {
	event := map[string]interface{}{
		"event_id":     e.EventID(),
		"aggregate_id": e.AggregateID(),
		"event_name":   e.EventName(),
		"occurred_on":  e.OccurredOn().Format(time.RFC3339Nano),
	}

	maps.Copy(event, e.User.ToPrimitives())

	return event
}

func (e UserRegisteredEvent) ToJSON() ([]byte, error) {
	event := e.ToPrimitives()

	json, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return json, nil
}

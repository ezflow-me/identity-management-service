package domain

// AggregateRoot is the base for all aggregates
type AggregateRoot struct {
	events []DomainEvent
}

// PullDomainEvents returns and clears all recorded domain events
func (a *AggregateRoot) PullDomainEvents() []DomainEvent {
	events := a.events
	a.events = nil
	return events
}

// Record adds a domain event to the aggregate
func (a *AggregateRoot) Record(event DomainEvent) {
	a.events = append(a.events, event)
}

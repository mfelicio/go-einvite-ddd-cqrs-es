package cqrs

type EventBus interface {
	Publish(event DomainEvent)
	Subscribe(eventName string, handler EventHandler)
}

type EventHandler func(event DomainEvent)

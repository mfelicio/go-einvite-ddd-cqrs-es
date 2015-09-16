package memory

import (
	. "github.com/mfelicio/einvite/backend/cqrs"
	"log"
)

type eventBus struct {
	subscriptions map[string][]EventHandler
	invokeHandler func(handler EventHandler, event DomainEvent)
}

// ctor
func NewEventBus(async bool) *eventBus {

	var invokeHandler func(handler EventHandler, event DomainEvent)
	if async {
		invokeHandler = func(handler EventHandler, event DomainEvent) { go handler(event) }
	} else {
		invokeHandler = func(handler EventHandler, event DomainEvent) { handler(event) }
	}

	return &eventBus{
		subscriptions: make(map[string][]EventHandler),
		invokeHandler: invokeHandler,
	}
}

func (this *eventBus) Publish(event DomainEvent) {

	if handlers, ok := this.subscriptions[event.GetName()]; ok {

		for _, handler := range handlers {

			this.invokeHandler(handler, event)
		}
	} else {

		log.Println("No EventHandlers found for event %s", event.GetName())
	}
}

func (this *eventBus) Subscribe(eventName string, handler EventHandler) {

	if handlers, ok := this.subscriptions[eventName]; !ok {
		this.subscriptions[eventName] = []EventHandler{handler}
	} else {
		this.subscriptions[eventName] = append(handlers, handler)
	}

}

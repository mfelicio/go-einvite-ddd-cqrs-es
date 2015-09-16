package memory

import (
	"github.com/mfelicio/einvite/backend/cqrs"
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
	check "gopkg.in/check.v1"
	"time"
)

type meb struct{}

var _ = check.Suite(&meb{}) //SetupSuite, SetupTest, TearDownTest and TearDownSuite not needed

func (t *meb) TestEventBusPubSynchronous(c *check.C) {

	bus := NewEventBus(false)

	invokes := make(map[string]int)
	invokes["ev1"] = 0
	invokes["ev2"] = 0
	invokes["ev3"] = 0

	handler := func(ev cqrs.DomainEvent) {
		invokes[ev.GetName()] = invokes[ev.GetName()] + 1
	}

	bus.Subscribe("ev1", handler)
	bus.Subscribe("ev1", handler)
	bus.Subscribe("ev2", handler)
	bus.Subscribe("ev3", handler)

	bus.Publish(&sourcing.Event{Name: "ev1"})
	bus.Publish(&sourcing.Event{Name: "ev2"})

	c.Check(invokes["ev1"], check.Equals, 2)
	c.Check(invokes["ev2"], check.Equals, 1)
}

func (t *meb) TestEventBusPubAsynchronous(c *check.C) {

	bus := NewEventBus(true)

	invokes := make(map[string]int)
	invokes["ev1"] = 0
	invokes["ev2"] = 0
	invokes["ev3"] = 0

	handler := func(ev cqrs.DomainEvent) {
		invokes[ev.GetName()] = invokes[ev.GetName()] + 1
	}

	bus.Subscribe("ev1", handler)
	bus.Subscribe("ev1", handler)
	bus.Subscribe("ev2", handler)
	bus.Subscribe("ev3", handler)

	bus.Publish(&sourcing.Event{Name: "ev1"})
	bus.Publish(&sourcing.Event{Name: "ev2"})

	c.Check(invokes["ev1"], check.Equals, 0)
	c.Check(invokes["ev2"], check.Equals, 0)

	time.Sleep(10 * time.Millisecond)

	c.Check(invokes["ev1"], check.Equals, 2)
	c.Check(invokes["ev2"], check.Equals, 1)
}

package sourcing

import (
	"time"
)

type Event struct {
	AggregateId string
	SequenceId  int
	Time        time.Time
	Name        string
	Payload     interface{}
}

func (this *Event) GetAggregateId() string {
	return this.AggregateId
}

func (this *Event) GetSequenceId() int {
	return this.SequenceId
}

func (this *Event) GetTime() time.Time {
	return this.Time
}

func (this *Event) GetName() string {
	return this.Name
}

func (this *Event) GetPayload() interface{} {
	return this.Payload
}

package cqrs

import (
	"time"
)

type DomainEvent interface {
	GetAggregateId() string
	GetSequenceId() int
	GetTime() time.Time
	GetName() string
	GetPayload() interface{}
}

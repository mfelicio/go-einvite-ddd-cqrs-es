package queries

import (
	"time"
)

type AggregateQuery struct {
	Id string
	Source AggregateQuerySource
}

type AggregateQuerySource struct {
	Version int
	Date    time.Time
}
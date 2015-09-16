package sourcing

type EventStore interface {
	Load(aggregateId string) ([]*Event, error)
	LoadSince(aggregateId string, start int) ([]*Event, error)
	Append(aggregateId string, expectedVersion int, events []*Event) error
}

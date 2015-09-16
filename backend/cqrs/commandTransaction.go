package cqrs

type CommandId string

type CommandTransaction interface {
	Command() interface{}
	Name() string
	Id() CommandId

	Commit(data interface{})
	Failed(err error)
}

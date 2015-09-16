package cqrs

type CommandStore interface {
	Push(id CommandId, name string, command interface{}) error
	Pop() (CommandTransaction, error)

	Check(id CommandId) (CommandResult, error)
}

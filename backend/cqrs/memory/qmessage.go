package memory

import (
	. "github.com/mfelicio/einvite/backend/cqrs"
)

const (
	qMessage_push    = 0
	qMessage_pop     = 1
	qMessage_check   = 2
	qMessage_commit  = 3
	qMessage_failed  = 4
	qMessage_command = 5
)

type qMessage struct {
	msgType int
	msg     interface{}
}

type pushMessage struct {
	id      CommandId
	name    string
	command interface{}
	out     chan (error)
}

type popMessage struct {
	out chan (CommandTransaction)
}

type commandMessage struct {
	id  CommandId
	out chan (CommandTransaction)
}

type checkMessage struct {
	id  CommandId
	out chan (*CommandResult)
}

type failedMessage struct {
	id  CommandId
	err error
}

type commitMessage struct {
	id   CommandId
	data interface{}
}

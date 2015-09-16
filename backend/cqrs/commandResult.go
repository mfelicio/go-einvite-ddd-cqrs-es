package cqrs

import (
	"time"
)

type CommandResult struct {
	Status  CommandStatus
	Updated time.Time
	Data    interface{}
}

type CommandStatus int

const (
	CommandStatus_Queued     CommandStatus = 1
	CommandStatus_Processing CommandStatus = 2
	CommandStatus_Ok         CommandStatus = 3
	CommandStatus_Error      CommandStatus = 4
)

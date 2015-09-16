package commands

import (
	"github.com/mfelicio/einvite/backend/cqrs"
	"github.com/mfelicio/einvite/backend/cqrs/services"
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
)

func BindHandlers(framework sourcing.Framework, cmdService *services.CommandService, eventBus cqrs.EventBus) {

	meetingCmdHandler := NewMeetingCommandHandler(framework, eventBus)

	cmdService.Bind(
		Command_CreateMeeting,
		func(cmd interface{}) (interface{}, error) {

			return meetingCmdHandler.HandleCreateMeeting(cmd.(*CreateMeeting))
		})

	cmdService.Bind(
		Command_ChangeTitle,
		func(cmd interface{}) (interface{}, error) {

			return meetingCmdHandler.HandleChangeTitle(cmd.(*ChangeTitle))
		})
}

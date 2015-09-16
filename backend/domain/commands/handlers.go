package commands

import (
	"github.com/mfelicio/einvite/backend/cqrs"
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
	. "github.com/mfelicio/einvite/backend/domain"
	. "github.com/mfelicio/einvite/backend/domain/model"
)

//handler for CreateMeetingCommand should create MeetingAggregate and then associate meeting to UserAggregate
//after storing MeetingAggregate events. The processing of the command only concludes when both aggregates are created, not necessarily in the same transaction
//however, the CreateMeetingCommand must be idempotent. If associate to user fails, the reprocessing of the command will have no side effects on MeetingAggregate
//and will only do the unfinished part, associate meeting to user

type MeetingCommandHandler struct {
	framework sourcing.Framework
	eventBus  cqrs.EventBus
}

func NewMeetingCommandHandler(framework sourcing.Framework, eventBus cqrs.EventBus) *MeetingCommandHandler {

	return &MeetingCommandHandler{
		framework: framework,
		eventBus:  eventBus,
	}
}

func (this *MeetingCommandHandler) HandleCreateMeeting(cmd *CreateMeeting) (interface{}, error) {

	var ref *MeetingRef

	// creates a new meeting aggregate
	events, err := this.framework.Update(AggregateType_Meeting, cmd.MeetingId,

		func(e interface{}) {

			ref = e.(*Meeting).Create(cmd.User, cmd.What, cmd.Title, cmd.Suggestions, cmd.Invitations)
		})

	if err != nil {
		return nil, err
	}

	// publishes the events in a finally block, regardless of the user association outcome
	defer this.publish(events)

	if cmd.User.Type != UserType_Anonymous {
		//associates to user aggregate
		//publishes resulting events
	}

	return ref, err
}

func (this *MeetingCommandHandler) HandleChangeTitle(cmd *ChangeTitle) (interface{}, error) {

	events, err := this.framework.Update(AggregateType_Meeting, cmd.MeetingId,

		func(e interface{}) {

			e.(*Meeting).ChangeTitle(cmd.RequesterId, cmd.Title)
		})

	defer this.publish(events)

	//no data to return, just whatever possible errors
	return nil, err
}

func (this *MeetingCommandHandler) publish(events []*sourcing.Event) {

	if events == nil {
		return
	}

	// publish events
	for _, ev := range events {

		this.eventBus.Publish(ev)
	}
}

package model

import (
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
	. "github.com/mfelicio/einvite/backend/domain/events"
	. "github.com/mfelicio/einvite/backend/domain/services"
	"reflect"
)

/**

  Integration with go-eventsourcing

  This file contains bindings for the following aggregates
	- Meeting
	- User

**/

//Meeting Aggregate bindings
var (
	AggregateType_Meeting = reflect.TypeOf(&Meeting{})
	AggregateName_Meeting = "meeting"

	createMeeting = func(refSvc MeetingRefService) func() interface{} {
		//closure to capture refSvc from global Bind function
		return func() interface{} { return NewMeeting(refSvc) }
	}
	initMeeting = func(m interface{}, a sourcing.Aggregate) {
		m.(*Meeting).Aggregate = a
	}
)

//Meeting Event bindings
var (
	invitationAcceptedEvent  = reflect.TypeOf(&InvitationAccepted{})
	invitationDeclinedEvent  = reflect.TypeOf(&InvitationDeclined{})
	invitationCreatedEvent   = reflect.TypeOf(&InvitationCreated{})
	invitationDeletedEvent   = reflect.TypeOf(&InvitationDeleted{})
	meetingCreatedEvent      = reflect.TypeOf(&MeetingCreated{})
	meetingTitleChangedEvent = reflect.TypeOf(&MeetingTitleChanged{})
	participantAddedEvent    = reflect.TypeOf(&ParticipantAdded{})
	participantRemovedEvent  = reflect.TypeOf(&ParticipantRemoved{})
	suggestionAddedEvent     = reflect.TypeOf(&SuggestionAdded{})
	suggestionRemovedEvent   = reflect.TypeOf(&SuggestionRemoved{})
	voteAddedEvent           = reflect.TypeOf(&VoteAdded{})
	voteChangedEvent         = reflect.TypeOf(&VoteChanged{})
	voteRemovedEvent         = reflect.TypeOf(&VoteRemoved{})

	onInvitationAcceptedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onInvitationAccepted(e.(*InvitationAccepted))
	}

	onInvitationDeclinedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onInvitationDeclined(e.(*InvitationDeclined))
	}

	onInvitationCreatedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onInvitationCreated(e.(*InvitationCreated))
	}

	onInvitationDeletedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onInvitationDeleted(e.(*InvitationDeleted))
	}

	onMeetingCreatedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onCreated(e.(*MeetingCreated))
	}

	onMeetingTitleChangedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onTitleChanged(e.(*MeetingTitleChanged))
	}

	onParticipantAddedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onParticipantAdded(e.(*ParticipantAdded))
	}

	onParticipantRemovedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onParticipantRemoved(e.(*ParticipantRemoved))
	}

	onSuggestionAddedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onSuggestionAdded(e.(*SuggestionAdded))
	}

	onSuggestionRemovedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onSuggestionRemoved(e.(*SuggestionRemoved))
	}

	onVoteAddedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onVoteAdded(e.(*VoteAdded))
	}

	onVoteChangedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onVoteChanged(e.(*VoteChanged))
	}

	onVoteRemovedEvent = func(s interface{}, e interface{}) {
		s.(*Meeting).state.onVoteRemoved(e.(*VoteRemoved))
	}
)

//SystemUser Aggregate bindings
var (
	systemUserAggregate     = reflect.TypeOf(&SystemUser{})
	systemUserAggregateName = "systemUser"
)

func BindEvents(f sourcing.Framework, refService MeetingRefService) {

	//Bind MeetingAggregate
	b := f.BindNamed(AggregateType_Meeting, AggregateName_Meeting, createMeeting(refService), initMeeting)

	b.On(invitationAcceptedEvent, onInvitationAcceptedEvent)
	b.On(invitationDeclinedEvent, onInvitationDeclinedEvent)
	b.On(invitationCreatedEvent, onInvitationCreatedEvent)
	b.On(invitationDeletedEvent, onInvitationDeletedEvent)
	b.On(meetingCreatedEvent, onMeetingCreatedEvent)
	b.On(meetingTitleChangedEvent, onMeetingTitleChangedEvent)
	b.On(participantAddedEvent, onParticipantAddedEvent)
	b.On(participantRemovedEvent, onParticipantRemovedEvent)
	b.On(suggestionAddedEvent, onSuggestionAddedEvent)
	b.On(suggestionRemovedEvent, onSuggestionRemovedEvent)
	b.On(voteAddedEvent, onVoteAddedEvent)
	b.On(voteChangedEvent, onVoteChangedEvent)
	b.On(voteRemovedEvent, onVoteRemovedEvent)

}

package events

import (
	. "github.com/mfelicio/einvite/backend/domain"
)

type MeetingCreated struct {
	What What
	Ref  *MeetingRef
	User *MeetingUser
}

type MeetingTitleChanged struct {
	Value string
}

type InvitationCreated struct {
	Id         ParticipantId
	Invitation *Invitation
}

type InvitationDeleted struct {
	Id ParticipantId
}

type InvitationAccepted struct {
	Id ParticipantId
}

type InvitationDeclined struct {
	Id ParticipantId
}

type ParticipantAdded struct {
	Id   ParticipantId
	Type ParticipantType

	User        *MeetingUser
	CompanionOf *MeetingUser
}

type ParticipantRemoved struct {
	Id ParticipantId
}

type SuggestionAdded struct {
	Id            SuggestionId
	ParticipantId ParticipantId
	Data          *SuggestionValue
}

type SuggestionRemoved struct {
	Id         SuggestionId
	WhoRemoved ParticipantId
}

type VoteAdded struct {
	ParticipantId ParticipantId
	SuggestionId  SuggestionId
	Value         VoteType
}

type VoteChanged struct {
	ParticipantId ParticipantId
	SuggestionId  SuggestionId
	Old           VoteType
	New           VoteType
}

type VoteRemoved struct {
	ParticipantId ParticipantId
	SuggestionId  SuggestionId
	Value         VoteType
}

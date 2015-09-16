package commands

import (
	. "github.com/mfelicio/einvite/backend/domain"
)

type CreateMeeting struct {
	MeetingId string

	User *MeetingUser

	Title       string
	What        What
	Suggestions []*SuggestionValue
	Invitations []*Invitation
}

type ChangeTitle struct {
	MeetingId string

	RequesterId ParticipantId
	Title       string
}

const (
	Command_CreateMeeting = "CreateMeeting"
	Command_ChangeTitle   = "ChangeTitle"
)

package queries

import (
	. "github.com/mfelicio/einvite/backend/domain"
)

type MeetingDetails struct {
	AggregateQuery

	Title string
	What  What
	Where SuggestionId
	When  SuggestionId
	Which SuggestionId

	Invitations  []*MeetingInvitation
	Participants []*MeetingParticipant
	Suggestions  []*MeetingSuggestion
}

type MeetingParticipant struct {
	Id         ParticipantId
	Type       ParticipantType
	Commitment ParticipantCommitment
	User       *MeetingUser
}

type MeetingInvitation struct {
	Invitation

	Id ParticipantId //invitation share same id as participant
}

type MeetingSuggestion struct {
	SuggestionValue

	Id    SuggestionId
	Votes map[VoteType][]ParticipantId
}

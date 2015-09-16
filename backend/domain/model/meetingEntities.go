package model

import (
	. "github.com/mfelicio/einvite/backend/domain"
)

// Participant

type Participant struct {
	Id          ParticipantId
	Type        ParticipantType
	Commitment  ParticipantCommitment
	User        *MeetingUser
	CompanionOf *MeetingUser
}

func NewParticipant(id ParticipantId, pType ParticipantType, user *MeetingUser, companionOf *MeetingUser) *Participant {

	return &Participant{
		Id:          id,
		Type:        pType,
		Commitment:  ParticipantCommitment_Undefined,
		User:        user,
		CompanionOf: companionOf,
	}
}

// Suggestion

type Suggestion struct {
	Id           SuggestionId
	Data         *SuggestionValue
	WhoSuggested ParticipantId
	Voters       map[ParticipantId]VoteType
}

func NewSuggestion(id SuggestionId, data *SuggestionValue, participantId ParticipantId) *Suggestion {
	return &Suggestion{
		Id:           id,
		Data:         data,
		WhoSuggested: participantId,
		Voters:       make(map[ParticipantId]VoteType),
	}
}

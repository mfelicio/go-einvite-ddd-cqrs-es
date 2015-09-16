package contracts

import (
	"github.com/mfelicio/einvite/backend/domain"
)

type CreateMeetingRequest struct {
	ApiCommand

	Title       string
	What        domain.What
	Suggestions []*domain.SuggestionValue
	Invitations []*domain.Invitation
}

type ChangeTitleRequest struct {
	ApiCommand

	Title string
}

type AddSuggestionRequest struct {
	ApiCommand

	domain.SuggestionValue
}

type VoteRequest struct {
	ApiCommand

	Value domain.VoteType
}

type AddInvitationRequest struct {
	ApiCommand

	domain.Invitation
}

type AutoEnrollRequest struct {
	ApiCommand

	domain.Invitation
}

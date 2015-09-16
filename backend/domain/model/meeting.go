package model

import (
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
	. "github.com/mfelicio/einvite/backend/domain"
	. "github.com/mfelicio/einvite/backend/domain/events"
	. "github.com/mfelicio/einvite/backend/domain/services"
)

//must have a participantList per User, usually 1 user = 1 participant but the user can bring more people with him (wife/husband)

type Meeting struct {
	sourcing.Aggregate

	refService MeetingRefService

	state *meetingState
}

//Ctor
func NewMeeting(refService MeetingRefService) *Meeting {
	return &Meeting{
		refService: refService,
		state:      newState(),
	}
}

//Commands
func (this *Meeting) Create(creator *MeetingUser, what What, title string, suggestions []*SuggestionValue, invitations []*Invitation) *MeetingRef {

	if this.state.Ref != nil {

		//already created
		return this.state.Ref
	}

	ref := this.refService.NewRef(this.Id())

	this.On(&MeetingCreated{
		User: creator,
		What: what,
		Ref:  ref,
	})

	if title != this.state.Title {
		this.On(&MeetingTitleChanged{
			Value: title,
		})
	}

	//TODO: set settings

	for _, suggestion := range suggestions {
		//TODO: creator participant Id should not be hardcoded
		this.CreateSuggestion(1, suggestion)
	}

	for _, i := range invitations {
		this.On(&InvitationCreated{
			Id:         this.state.NextParticipantId,
			Invitation: i,
		})
	}

	return ref
}

func (this *Meeting) ChangeTitle(requesterId ParticipantId, title string) {

	if p, ok := this.state.Participants[requesterId]; ok {

		if p.Type == ParticipantType_Admin {

			if title != this.state.Title {

				this.On(&MeetingTitleChanged{
					Value: title,
				})
			}
		}
	}
}

func (this *Meeting) AddInvitation(requesterId ParticipantId, invitation *Invitation) {

	//TODO: validate that the user isn't in the meeting already, use lookup service??

	if p, ok := this.state.Participants[requesterId]; ok {

		if p.Type == ParticipantType_Admin {

			this.On(&InvitationCreated{
				Id:         this.state.NextParticipantId,
				Invitation: invitation,
			})
		}
	}

}

func (this *Meeting) RemoveInvitation(requesterId ParticipantId, participantId ParticipantId) {

	if _, ok := this.state.Invitations[participantId]; ok {

		if p, pOk := this.state.Participants[requesterId]; pOk {

			if p.Type == ParticipantType_Admin {

				this.On(&InvitationDeleted{
					Id: participantId,
				})
			}
		}
	}
}

func (this *Meeting) AddCompanion(requesterId ParticipantId, user *MeetingUser) {

	if p, ok := this.state.Participants[requesterId]; ok {
		//TODO: validate that the user isn't in the meeting already, use lookup service??

		this.On(&ParticipantAdded{
			Id:          this.state.NextParticipantId,
			Type:        ParticipantType_UserCompanion,
			User:        user,
			CompanionOf: p.User,
		})
	}
}

func (this *Meeting) RemoveParticipant(requesterId ParticipantId, participantId ParticipantId) {

	if p, ok := this.state.Participants[participantId]; ok {

		if p.Type == ParticipantType_Admin || requesterId == participantId {

			//remove participant
			this.On(&ParticipantRemoved{
				Id: participantId,
			})
		}
	}
}

func (this *Meeting) AutoEnrollParticipant(participantType ParticipantType, user *MeetingUser) {

	//TODO: validate settings
	//TODO: validate that the user isn't in the meeting already, use lookup service??

	this.On(&ParticipantAdded{
		Id:   this.state.NextParticipantId,
		Type: participantType,
		User: user,
	})
}

func (this *Meeting) AcceptInvitation(requesterId ParticipantId) {
	// Validate that participant exists and status is pending
	if _, ok := this.state.Invitations[requesterId]; ok {

		this.On(&InvitationAccepted{requesterId})
	}
}

func (this *Meeting) DeclineInvitation(requesterId ParticipantId) {

	// Validate that participant exists and status is pending
	if _, ok := this.state.Invitations[requesterId]; ok {

		this.On(&InvitationDeclined{requesterId})
	}
}

func (this *Meeting) CreateSuggestion(requesterId ParticipantId, value *SuggestionValue) {

	//TODO: validate settings for suggestion type (can suggestions be made?)

	// Validate that participant exists and accepted the event
	if _, ok := this.state.Participants[requesterId]; ok {

		this.On(&SuggestionAdded{
			Id:            this.state.NextSuggestionId,
			ParticipantId: requesterId,
			Data:          value,
		})
	}
}

func (this *Meeting) RemoveSuggestion(requesterId ParticipantId, suggestionId SuggestionId) {

	if p, ok := this.state.Participants[requesterId]; ok {

		if s, sOk := this.state.Suggestions[suggestionId]; sOk {

			if p.Type == ParticipantType_Admin || requesterId == s.WhoSuggested {

				this.On(&SuggestionRemoved{
					Id:         suggestionId,
					WhoRemoved: requesterId,
				})
			}
		}
	}
}

func (this *Meeting) SetVote(requesterId ParticipantId, suggestionId SuggestionId, value VoteType) {

	if _, ok := this.state.Participants[requesterId]; ok {

		if s, sOk := this.state.Suggestions[suggestionId]; sOk {

			//TODO: validate settings

			//validade that user hasnt voted yet or his vote is different
			oldValue, exists := s.Voters[requesterId]

			if !exists {

				this.On(&VoteAdded{
					ParticipantId: requesterId,
					SuggestionId:  suggestionId,
					Value:         value,
				})
			} else if oldValue != value {

				this.On(&VoteChanged{
					ParticipantId: requesterId,
					SuggestionId:  suggestionId,
					Old:           oldValue,
					New:           value,
				})
			}

		}

	}
}

func (this *Meeting) RemoveVote(requesterId ParticipantId, suggestionId SuggestionId) {

	if _, ok := this.state.Participants[requesterId]; ok {

		// Validate that suggestion exists and can be voted
		if s, sOk := this.state.Suggestions[suggestionId]; sOk {

			//validade that user has voted
			if v, vOk := s.Voters[requesterId]; vOk {

				this.On(&VoteRemoved{
					ParticipantId: requesterId,
					SuggestionId:  suggestionId,
					Value:         v,
				})
			}
		}

	}
}

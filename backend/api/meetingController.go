package api

import (
	"github.com/emicklei/go-restful"
	"github.com/mfelicio/einvite/backend/api/contracts"
	"github.com/mfelicio/einvite/backend/cqrs"
	"github.com/mfelicio/einvite/backend/cqrs/services"
	"github.com/mfelicio/einvite/backend/domain"
	"github.com/mfelicio/einvite/backend/domain/queries"
	"net/http"
	"time"
)

type MeetingController struct {
	commandStore cqrs.CommandStore
	querySvc     services.QueryService
}

func NewMeetingController(commandStore cqrs.CommandStore) *MeetingController {
	return &MeetingController{
		commandStore: commandStore,
	}
}

//GET /meeting/{id}
func (c *MeetingController) GetDetails(request *restful.Request, response *restful.Response) {
	sampleQueryResponse(request, response)
}

//POST /meeting
func (c *MeetingController) Create(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//POST /meeting/{id}
func (c *MeetingController) ChangeTitle(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//POST /meeting/{id}/invitations
func (c *MeetingController) AddInvitation(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//DELETE /meeting/{id}/invitation/{invitationId}
func (c *MeetingController) RemoveInvitation(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//PUT /meeting/{id}/invitation/{invitationId}/accept
func (c *MeetingController) AcceptInvitation(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//PUT /meeting/{id}/invitation/{invitationId}/decline
func (c *MeetingController) DeclineInvitation(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//POST /meeting/{id}/participants
func (c *MeetingController) AutoEnrollParticipant(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//DELETE /meeting/{id}/participant/{participantId}
func (c *MeetingController) RemoveParticipant(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//PUT /meeting/{id}/participant/{participantId}/companion
func (c *MeetingController) AddParticipantCompanion(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//PUT /meeting/{id}/suggestions
func (c *MeetingController) AddSuggestion(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//DELETE /meeting/{id}/suggestion/{suggestionId}
func (c *MeetingController) RemoveSuggestion(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//PUT /meeting/{id}/suggestion/{suggestionId}/votes/{participantId}
func (c *MeetingController) Vote(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

//DELETE /meeting/{id}/suggestion/{suggestionId}/votes/{participantId}
func (c *MeetingController) RemoveVote(request *restful.Request, response *restful.Response) {
	sampleCommandResponse(request, response)
}

func sampleQueryResponse(request *restful.Request, response *restful.Response) {

	details := &contracts.MeetingDetailsResponse{
		MeetingDetails: queries.MeetingDetails{
			AggregateQuery: queries.AggregateQuery{
				Id: "the meeting Id",
				Source: queries.AggregateQuerySource{
					Version: 22,
					Date:    time.Now(),
				},
			},

			Title: "the meeting title",
			What:  domain.What_Restaurant,
			Where: 1,
			When:  1,
			Which: -1,

			Invitations: []*queries.MeetingInvitation{
				&queries.MeetingInvitation{
					Id: 3,
					Invitation: domain.Invitation{
						Type: domain.ParticipantType_User,
						User: &domain.MeetingUser{
							Type:       domain.UserType_Google,
							ExternalId: 9182746124367,
						},
					},
				},
				&queries.MeetingInvitation{
					Id: 5,
					Invitation: domain.Invitation{
						Type: domain.ParticipantType_User,
						User: &domain.MeetingUser{
							Type:       domain.UserType_Email,
							ExternalId: "foo@bar.com",
						},
					},
				},
			},

			Participants: []*queries.MeetingParticipant{
				&queries.MeetingParticipant{
					Id:         1,
					Type:       domain.ParticipantType_Admin,
					Commitment: domain.ParticipantCommitment_Going,
					User: &domain.MeetingUser{
						Type:       domain.UserType_Google,
						ExternalId: 123192371273,
					},
				},
				&queries.MeetingParticipant{
					Id:         2,
					Type:       domain.ParticipantType_User,
					Commitment: domain.ParticipantCommitment_Undefined,
					User: &domain.MeetingUser{
						Type:       domain.UserType_Anonymous,
						ExternalId: "John doe, an anonymous user",
					},
				},
				&queries.MeetingParticipant{
					Id:         4,
					Type:       domain.ParticipantType_User,
					Commitment: domain.ParticipantCommitment_Maybe,
					User: &domain.MeetingUser{
						Type:       domain.UserType_Google,
						ExternalId: 432423412312,
					},
				},
			},

			Suggestions: []*queries.MeetingSuggestion{
				&queries.MeetingSuggestion{
					Id: 1,
					SuggestionValue: domain.SuggestionValue{
						Type:      domain.SuggestionType_When,
						ValueType: domain.SuggestionValueType_Text,
						Value:     "sexta-feira 13",
					},
					Votes: make(map[domain.VoteType][]domain.ParticipantId),
				},
				&queries.MeetingSuggestion{
					Id: 2,
					SuggestionValue: domain.SuggestionValue{
						Type:      domain.SuggestionType_Where,
						ValueType: domain.SuggestionValueType_Text,
						Value:     "no prestige",
					},
					Votes: make(map[domain.VoteType][]domain.ParticipantId),
				},
			},
		},
	}

	//details
	details.Suggestions[0].Votes[domain.VoteType_Yes] = []domain.ParticipantId{1, 2}
	details.Suggestions[0].Votes[domain.VoteType_Maybe] = []domain.ParticipantId{}
	details.Suggestions[0].Votes[domain.VoteType_No] = []domain.ParticipantId{4}

	details.Suggestions[1].Votes[domain.VoteType_Yes] = []domain.ParticipantId{1}
	details.Suggestions[1].Votes[domain.VoteType_Maybe] = []domain.ParticipantId{4}
	details.Suggestions[1].Votes[domain.VoteType_No] = []domain.ParticipantId{2}

	response.WriteEntity(details)
}

func sampleCommandResponse(request *restful.Request, response *restful.Response) {
	response.WriteHeader(http.StatusAccepted)
	response.AddHeader("Location", "/command/here-goes-the-commandId-received-in-the-request/status")

	// due to a bug on go-restful the line below is necessary or else the Accepted header won't be sent
	response.WriteEntity("")
}

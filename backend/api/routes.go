package api

import (
	"github.com/emicklei/go-restful"
	"github.com/mfelicio/einvite/backend/api/contracts"
	"github.com/mfelicio/einvite/backend/cqrs"
	"log"
)

func Register(cmdStore cqrs.CommandStore) {

	log.Println("Registering webapi")

	registerMeetingRoutes(NewMeetingController(cmdStore))
	registerUserRoutes(NewUserController(cmdStore))
	registerCommandRoutes(NewCommandController(cmdStore))

	registerLoginRoutes()

}

func registerLoginRoutes() {

}

func registerCommandRoutes(controller *CommandController) {

	ws := new(restful.WebService)

	idParam := ws.PathParameter("id", "command id").DataType("string").Required(true)
	tokenParam := ws.QueryParameter("token", "access token").DataType("string").Required(true)

	ws.Path("/command").Doc("Command operations").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/{id}/status").To(controller.GetStatus).
		// docs
		Doc("Get status").
		Operation("getStatus").
		Param(tokenParam).
		Param(idParam).
		Writes(contracts.ApiCommandStatus{})) // on the response

	ws.Route(ws.DELETE("/{id}").To(controller.Cancel).
		// docs
		Doc("Cancel command").
		Operation("cancelCommand").
		Param(tokenParam).
		Param(idParam).
		Writes(contracts.ApiCommandResponse{})) // on the response

	restful.Add(ws)
}

func registerMeetingRoutes(controller *MeetingController) {

	ws := new(restful.WebService)

	idParam := ws.PathParameter("id", "meeting id").DataType("string").Required(true)
	suggestionIdParam := ws.PathParameter("suggestionId", "suggestion id").DataType("int").Required(true)
	participantIdParam := ws.PathParameter("participantId", "participant id").DataType("int").Required(true)
	invitationIdParam := ws.PathParameter("invitationId", "invitation id").DataType("int").Required(true)

	tokenParam := ws.QueryParameter("token", "access token").DataType("string").Required(true)

	ws.Path("/meeting").Doc("Meeting operations").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.POST("").To(controller.Create).
		// docs
		Doc("Create meeting").
		Operation("createMeeting").
		Param(tokenParam).
		Reads(contracts.CreateMeetingRequest{}). // from the request
		Writes(contracts.ApiCommandResponse{}))  // on the response

	ws.Route(ws.GET("/{id}").To(controller.GetDetails).
		// docs
		Doc("Get details").
		Operation("getDetails").
		Param(tokenParam).
		Param(idParam).
		Writes(contracts.MeetingDetailsResponse{})) // on the response

	ws.Route(ws.PUT("/{id}/title").To(controller.ChangeTitle).
		// docs
		Doc("Change title").
		Operation("changeTitle").
		Param(tokenParam).
		Param(idParam).
		Reads(contracts.ChangeTitleRequest{}).
		Writes(contracts.ApiCommandResponse{})) // on the response

	ws.Route(ws.POST("/{id}/invitations").To(controller.AddInvitation).
		// docs
		Doc("Add invitation").
		Operation("addInvitation").
		Param(tokenParam).
		Param(idParam).
		Reads(contracts.AddInvitationRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.DELETE("/{id}/invitation/{invitationId}").To(controller.RemoveInvitation).
		// docs
		Doc("Remove invitation").
		Operation("removeInvitation").
		Param(tokenParam).
		Param(idParam).Param(invitationIdParam).
		Reads(contracts.ApiCommand{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.PUT("/{id}/invitation/{invitationId}/accept").To(controller.AcceptInvitation).
		// docs
		Doc("Accept invitation").
		Operation("acceptInvitation").
		Param(tokenParam).
		Param(idParam).Param(invitationIdParam).
		Reads(contracts.ApiCommand{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.PUT("/{id}/invitation/{invitationId}/decline").To(controller.DeclineInvitation).
		// docs
		Doc("Decline invitation").
		Operation("declineInvitation").
		Param(tokenParam).
		Param(idParam).Param(invitationIdParam).
		Reads(contracts.ApiCommand{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.POST("/{id}/participants").To(controller.AutoEnrollParticipant).
		// docs
		Doc("Auto enroll participant").
		Operation("autoEnrollParticipant").
		Param(tokenParam).
		Param(idParam).
		Reads(contracts.AutoEnrollRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.DELETE("/{id}/participant/{participantId}").To(controller.RemoveParticipant).
		// docs
		Doc("Remove participant").
		Operation("removeParticipant").
		Param(tokenParam).
		Param(idParam).Param(participantIdParam).
		Reads(contracts.ApiCommand{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.PUT("/{id}/participant/{participantId}/companion").To(controller.AddParticipantCompanion).
		// docs
		Doc("Add participant companion").
		Operation("addParticipantCompanion").
		Param(tokenParam).
		Param(idParam).Param(participantIdParam).
		Reads(contracts.AutoEnrollRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.POST("/{id}/suggestions").To(controller.AddSuggestion).
		// docs
		Doc("Add suggestion").
		Operation("addSuggestion").
		Param(tokenParam).
		Param(idParam).
		Reads(contracts.AddSuggestionRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.DELETE("/{id}/suggestion/{suggestionId}").To(controller.RemoveSuggestion).
		// docs
		Doc("Remove suggestion").
		Operation("removeSuggestion").
		Param(tokenParam).
		Param(idParam).Param(suggestionIdParam).
		Reads(contracts.AutoEnrollRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.PUT("/{id}/suggestion/{suggestionId}/votes/{participantId}").To(controller.Vote).
		// docs
		Doc("Vote").
		Operation("vote").
		Param(tokenParam).
		Param(idParam).Param(suggestionIdParam).Param(participantIdParam).
		Reads(contracts.VoteRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	ws.Route(ws.DELETE("/{id}/suggestion/{suggestionId}/votes/{participantId}").To(controller.RemoveVote).
		// docs
		Doc("Remove vote").
		Operation("removeVote").
		Param(tokenParam).
		Param(idParam).Param(suggestionIdParam).Param(participantIdParam).
		Reads(contracts.AutoEnrollRequest{}).
		Writes(contracts.ApiCommandResponse{}))

	restful.Add(ws)
}

func registerUserRoutes(controller *UserController) {

	ws := new(restful.WebService)

	idParam := ws.PathParameter("id", "user id").DataType("string").Required(true)

	tokenParam := ws.QueryParameter("token", "access token").DataType("string").Required(true)

	ws.Path("/user").Doc("User operations").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/{id}/meetings").To(controller.List).
		// docs
		Doc("List meetings").
		Operation("listMeetings").
		Param(tokenParam).
		Param(idParam).
		Writes(contracts.UserMeetingsResponse{})) // on the response

	restful.Add(ws)
}

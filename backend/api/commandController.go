package api

import (
	"github.com/emicklei/go-restful"
	//"github.com/mfelicio/einvite/backend/api/contracts"
	"github.com/mfelicio/einvite/backend/cqrs"
)

type CommandController struct {
	commandStore cqrs.CommandStore
}

func NewCommandController(commandStore cqrs.CommandStore) *CommandController {
	return &CommandController{
		commandStore: commandStore,
	}
}

//GET /command/{id}/status
func (c *CommandController) GetStatus(request *restful.Request, response *restful.Response) {
	//sample..
	obj := make(map[string]interface{})

	obj["MeetingId"] = "the resulting meeting Id, assuming the command was a CreateMeetingCommand"

	response.WriteEntity(obj)
}

//DELETE /command/{id}
func (c *CommandController) Cancel(request *restful.Request, response *restful.Response) {

}

package api

import (
	"github.com/emicklei/go-restful"
	"github.com/mfelicio/einvite/backend/cqrs"
)

type UserController struct {
	commandStore cqrs.CommandStore
}

func NewUserController(commandStore cqrs.CommandStore) *UserController {
	return &UserController{
		commandStore: commandStore,
	}
}

//GET /user/{id}/meetings
func (c *UserController) List(request *restful.Request, response *restful.Response) {

}

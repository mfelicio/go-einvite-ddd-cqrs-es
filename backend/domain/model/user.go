package model

import (
	//. "github.com/mfelicio/einvite/backend/domain"
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
)

//activity streams?
//social stuff here?
//multiple credential management? same system user with google id, facebook id
type SystemUser struct {
	sourcing.Aggregate
}

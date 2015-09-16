package contracts

import (
	"github.com/mfelicio/einvite/backend/domain/queries"
)

type MeetingDetailsResponse struct {
	ApiQueryResponse

	queries.MeetingDetails
}

type UserMeetingsResponse struct {
	ApiQueryResponse
}

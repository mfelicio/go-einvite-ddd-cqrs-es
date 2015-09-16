package services

import (
	"fmt"
	. "github.com/mfelicio/einvite/backend/domain"
)

//must be properly implemented somewhere.. use some cryptography
type MeetingRefService interface {
	NewRef(meetingId string) *MeetingRef
	GetMeetingId(ref *MeetingRef) string
}

type dummyMeetingRefService struct{}

func NewDummyMeetingService() MeetingRefService {
	return &dummyMeetingRefService{}
}

func (this *dummyMeetingRefService) NewRef(meetingId string) *MeetingRef {

	return &MeetingRef{
		ParticipantRef: fmt.Sprintf("ParticipantRef:%s", meetingId),
		AdminRef:       fmt.Sprintf("AdminRef:%s", meetingId),
	}
}

func (this *dummyMeetingRefService) GetMeetingId(ref *MeetingRef) string {

	return ref.AdminRef[len("AdminRef:"):]
}

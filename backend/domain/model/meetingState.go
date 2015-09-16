package model

import (
	. "github.com/mfelicio/einvite/backend/domain"
	. "github.com/mfelicio/einvite/backend/domain/events"
)

type meetingState struct {
	Ref     *MeetingRef
	Creator *MeetingUser

	Title string
	What  What

	NextParticipantId ParticipantId
	NextSuggestionId  SuggestionId

	Invitations map[ParticipantId]*Invitation

	Participants map[ParticipantId]*Participant
	Suggestions  map[SuggestionId]*Suggestion

	SelectedWhere SuggestionId
	SelectedWhen  SuggestionId
	SelectedWhich SuggestionId
}

//Ctor
func newState() *meetingState {
	return &meetingState{
		NextParticipantId: 1,
		NextSuggestionId:  1,
		Invitations:       make(map[ParticipantId]*Invitation),
		Participants:      make(map[ParticipantId]*Participant),
		Suggestions:       make(map[SuggestionId]*Suggestion),
		SelectedWhere:     -1, //use proper const
		SelectedWhen:      -1, //use proper const
		SelectedWhich:     -1, //use proper const
	}
}

//State Event handlers
func (this *meetingState) onCreated(ev *MeetingCreated) {

	this.Ref = ev.Ref
	this.What = ev.What

	this.Creator = ev.User

	// add creator as participant
	p := NewParticipant(this.NextParticipantId, ParticipantType_Admin, this.Creator, nil)
	this.Participants[p.Id] = p

	this.NextParticipantId++
}

func (this *meetingState) onTitleChanged(ev *MeetingTitleChanged) {

	this.Title = ev.Value
}

func (this *meetingState) onInvitationCreated(ev *InvitationCreated) {

	this.Invitations[ev.Id] = ev.Invitation
	this.NextParticipantId++
}

func (this *meetingState) onInvitationDeleted(ev *InvitationDeleted) {

	delete(this.Invitations, ev.Id)
}

func (this *meetingState) onInvitationAccepted(ev *InvitationAccepted) {

	invitation := this.Invitations[ev.Id]
	delete(this.Invitations, ev.Id)

	p := NewParticipant(ev.Id, invitation.Type, invitation.User, nil)
	this.Participants[ev.Id] = p
}

func (this *meetingState) onInvitationDeclined(ev *InvitationDeclined) {

	delete(this.Invitations, ev.Id)
}

func (this *meetingState) onSuggestionAdded(ev *SuggestionAdded) {

	// add suggestion
	s := NewSuggestion(ev.Id, ev.Data, ev.ParticipantId)
	this.Suggestions[s.Id] = s

	this.NextSuggestionId++
}

func (this *meetingState) onSuggestionRemoved(ev *SuggestionRemoved) {

	// add suggestion
	delete(this.Suggestions, ev.Id)
}

func (this *meetingState) onParticipantAdded(ev *ParticipantAdded) {

	// add participant
	p := NewParticipant(ev.Id, ev.Type, ev.User, ev.CompanionOf)
	this.Participants[p.Id] = p

	this.NextParticipantId++
}

func (this *meetingState) onParticipantRemoved(ev *ParticipantRemoved) {

	//clear votes
	for _, s := range this.Suggestions {

		if _, vOk := s.Voters[ev.Id]; vOk {
			delete(s.Voters, ev.Id)
		}
	}

	delete(this.Participants, ev.Id)
}

func (this *meetingState) onVoteAdded(ev *VoteAdded) {

	s := this.Suggestions[ev.SuggestionId]
	s.Voters[ev.ParticipantId] = ev.Value
}

func (this *meetingState) onVoteChanged(ev *VoteChanged) {

	s := this.Suggestions[ev.SuggestionId]
	s.Voters[ev.ParticipantId] = ev.New
}

func (this *meetingState) onVoteRemoved(ev *VoteRemoved) {

	s := this.Suggestions[ev.SuggestionId]
	delete(s.Voters, ev.ParticipantId)
}

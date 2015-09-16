package domain

type MeetingUser struct {
	Type       UserType
	ExternalId interface{}
}

type SuggestionValue struct {
	Type      SuggestionType
	ValueType SuggestionValueType
	Value     interface{}
}

//contains info for a invited user
//can be used either at event creation or after created
type Invitation struct {
	Type ParticipantType
	User *MeetingUser
}

type MeetingRef struct {
	ParticipantRef string
	AdminRef       string
}

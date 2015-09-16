package domain

const (
	What_Generic    What = 0
	What_Restaurant What = 1
	What_Bar        What = 2
	What_Trip       What = 3
	What_Movie      What = 4
	What_Meeting    What = 5
	What_Beach      What = 6
)

const (
	ParticipantType_Admin         ParticipantType = 1
	ParticipantType_User          ParticipantType = 2
	ParticipantType_UserCompanion ParticipantType = 3 //wife/husband of a Participant whose type is User
	//ParticipantType_Anonymous     ParticipantType = 4
)

const (
	ParticipantCommitment_Going     ParticipantCommitment = 1
	ParticipantCommitment_NotGoing  ParticipantCommitment = 2
	ParticipantCommitment_Maybe     ParticipantCommitment = 3
	ParticipantCommitment_Undefined ParticipantCommitment = 4
)

const (
	UserType_Google    UserType = 1
	UserType_Facebook  UserType = 2
	UserType_Email     UserType = 3
	UserType_Anonymous UserType = 4
)

const (
	SuggestionType_Where SuggestionType = 1
	SuggestionType_When  SuggestionType = 2
	SuggestionType_Which SuggestionType = 3
)

const (
	SuggestionValueType_Text   SuggestionValueType = 1
	SuggestionValueType_Map    SuggestionValueType = 2
	SuggestionValueType_Places SuggestionValueType = 3

	SuggestionValueType_Date      SuggestionValueType = 102
	SuggestionValueType_DateRange SuggestionValueType = 103
)

const (
	VoteType_Yes   VoteType = "yes"
	VoteType_Maybe VoteType = "maybe"
	VoteType_No    VoteType = "no"
)

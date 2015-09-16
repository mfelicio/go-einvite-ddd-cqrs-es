package model

import (
	"github.com/mfelicio/einvite/backend/cqrs/memory"
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
	. "github.com/mfelicio/einvite/backend/domain"
	. "github.com/mfelicio/einvite/backend/domain/services"
	check "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type DomainTests struct {
	framework sourcing.Framework

	user *MeetingUser
}

var _ = check.Suite(&DomainTests{})

func (t *DomainTests) SetUpSuite(c *check.C) {

}

func (t *DomainTests) SetUpTest(c *check.C) {
	t.framework = sourcing.NewFramework(memory.NewEventStore())
	t.user = &MeetingUser{
		Type:       UserType_Google,
		ExternalId: "me user",
	}

	BindEvents(t.framework, NewDummyMeetingService())
}

func (t *DomainTests) TearDownTest(c *check.C) {
	t.framework = nil
}

func (t *DomainTests) TearDownSuite(c *check.C) {

}

func (t *DomainTests) TestCreateMeeting(c *check.C) {

	var ref *MeetingRef

	evs, err :=
		t.framework.Update(AggregateType_Meeting, "id1", func(m interface{}) {
			ref = m.(*Meeting).Create(t.user, What_Restaurant, "", []*SuggestionValue{}, []*Invitation{})
		})

	c.Check(ref, check.NotNil)
	c.Check(err, check.IsNil)
	c.Check(evs, check.HasLen, 1)

	//should be idempotent, but duplicates must not produce state change events
	evs, err =
		t.framework.Update(AggregateType_Meeting, "id1", func(m interface{}) {
			ref = m.(*Meeting).Create(t.user, What_Restaurant, "", []*SuggestionValue{}, []*Invitation{})
		})

	c.Check(ref, check.NotNil)
	c.Check(err, check.IsNil)
	c.Check(evs, check.HasLen, 0)
}

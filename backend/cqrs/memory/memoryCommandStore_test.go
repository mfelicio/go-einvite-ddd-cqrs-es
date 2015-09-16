package memory

import (
	"fmt"
	. "github.com/mfelicio/einvite/backend/cqrs"
	check "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) {
	check.TestingT(t) // tests entry point
}

type mcs struct{}

var _ = check.Suite(&mcs{}) //SetupSuite, SetupTest, TearDownTest and TearDownSuite not needed

func (t *mcs) TestFullQueue(c *check.C) {

	store := NewCommandStore(1)
	defer close(store.input)

	//1 item in queue
	err := store.Push("id1", "cmd-name", "cmd1")

	c.Check(err, check.IsNil)

	err = store.Push("id2", "cmd-name", "cmd2") //cant enqueue because len(queue) == limit

	c.Check(err, check.NotNil)

	//lets pop and then push again

	_, err = store.Pop() //queue has 0 items

	err = store.Push("id2", "cmd-name", "cmd2") //queue has 1 item

	c.Check(err, check.IsNil)

	err = store.Push("id3", "cmd-name", "cmd3") // repeating just for sanity check

	c.Check(err, check.NotNil)
}

func (t *mcs) TestBasicBehavior(c *check.C) {

	store := NewCommandStore(1)
	defer close(store.input)

	var id CommandId = "id1"

	store.Push(id, "cmd-name", "cmd1")

	transaction, _ := store.Pop()

	c.Check(transaction.Command(), check.Equals, "cmd1")

	transaction.Commit("cmd1 result")

	result, err := store.Check(id)

	c.Check(err, check.IsNil)

	c.Check(result.Status, check.Equals, CommandStatus_Ok)
	c.Check(result.Data, check.Equals, "cmd1 result")
}

func (t *mcs) TestCanPopBeforePushing(c *check.C) {
	store := NewCommandStore(1)
	defer close(store.input)

	var tr CommandTransaction

	go func() {
		tr, _ = store.Pop()
	}()

	store.Push("id1", "name1", "cmd1")
	time.Sleep(10 * time.Millisecond)

	c.Check(tr, check.NotNil)
}

func (t *mcs) TestCheckNonExistent(c *check.C) {
	store := NewCommandStore(1)
	defer close(store.input)

	_, err := store.Check("non existent id") //should not be able to check non existent commands

	c.Check(err, check.NotNil)
}

func (t *mcs) TestExpiration(c *check.C) {
	store := NewCommandStore(1)
	defer close(store.input)

	store.expiry = 200 * time.Millisecond

	var id CommandId = "id1"

	store.Push(id, "cmd-name", "cmd1")
	tr, _ := store.Pop()

	tr.Failed(fmt.Errorf("failure")) // check can be done within 200ms only

	_, err := store.Check(id) // should be able to check

	c.Check(err, check.IsNil)

	time.Sleep(100 * time.Millisecond)

	_, err = store.Check(id)

	c.Check(err, check.IsNil)

	time.Sleep(150 * time.Millisecond) //250ms have passed

	_, err = store.Check(id)

	c.Check(err, check.NotNil)
}

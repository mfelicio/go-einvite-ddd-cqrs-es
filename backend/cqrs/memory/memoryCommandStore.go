package memory

import (
	"fmt"
	. "github.com/mfelicio/einvite/backend/cqrs"
	"log"
	"time"
)

type commandTransaction struct {
	input   chan (*qMessage)
	id      CommandId
	name    string
	command interface{}
}

func NewTransaction(input chan (*qMessage), id CommandId, name string, cmd interface{}) CommandTransaction {
	return &commandTransaction{input: input, id: id, name: name, command: cmd}
}

func (this *commandTransaction) Command() interface{} {
	return this.command
}

func (this *commandTransaction) Name() string {
	return this.name
}

func (this *commandTransaction) Id() CommandId {
	return this.id
}

func (this *commandTransaction) Commit(data interface{}) {

	this.input <- &qMessage{
		msgType: qMessage_commit,
		msg: &commitMessage{
			id:   this.id,
			data: data,
		},
	}
}

func (this *commandTransaction) Failed(err error) {

	this.input <- &qMessage{
		msgType: qMessage_failed,
		msg: &failedMessage{
			id:  this.id,
			err: err,
		},
	}
}

type command struct {
	Id      CommandId
	Name    string
	Payload interface{}

	*CommandResult
}

type commandStore struct {
	commands map[CommandId]*command
	queue    chan (CommandId)
	input    chan (*qMessage)

	expiry time.Duration
	limit  int
}

func NewCommandStore(limit int) *commandStore {

	store := &commandStore{
		commands: make(map[CommandId]*command),
		queue:    make(chan (CommandId), limit),
		input:    make(chan (*qMessage)),
		expiry:   1 * time.Hour,
		limit:    limit,
	}

	go store.processQueue()

	return store
}

func (this *commandStore) Push(id CommandId, name string, command interface{}) error {

	out := make(chan (error))
	defer close(out)

	this.input <- &qMessage{
		msgType: qMessage_push,
		msg: &pushMessage{
			id:      id,
			name:    name,
			command: command,
			out:     out,
		},
	}
	return <-out
}
func (this *commandStore) Pop() (CommandTransaction, error) {

	out := make(chan (CommandTransaction))
	defer close(out)

	this.input <- &qMessage{
		msgType: qMessage_pop,
		msg:     &popMessage{out},
	}

	transaction := <-out

	return transaction, nil
}

func (this *commandStore) Check(id CommandId) (CommandResult, error) {

	out := make(chan (*CommandResult))
	defer close(out)

	this.input <- &qMessage{
		msgType: qMessage_check,
		msg: &checkMessage{
			id:  id,
			out: out,
		},
	}

	result := <-out

	if result == nil {
		return CommandResult{}, fmt.Errorf("No result found for command id %s", id)
	}

	return *result, nil
}

func (this *commandStore) processQueue() {

	for {
		msg, ok := <-this.input

		if ok {
			switch msg.msgType {

			case qMessage_push:
				this.onPush(msg.msg.(*pushMessage))
			case qMessage_pop:
				this.onPop(msg.msg.(*popMessage))
			case qMessage_command:
				this.onCommand(msg.msg.(*commandMessage))
			case qMessage_commit:
				this.onCommit(msg.msg.(*commitMessage))
			case qMessage_failed:
				this.onFailed(msg.msg.(*failedMessage))
			case qMessage_check:
				this.onCheck(msg.msg.(*checkMessage))
			}
		} else {
			//queue closed?
			log.Println("MemoryCommandStore input queue closed")
			break
		}
	}
}

func (this *commandStore) onPush(msg *pushMessage) {

	if len(this.queue) == this.limit {
		msg.out <- fmt.Errorf("Queue is full at %d elements.", this.limit)
		return
	}

	cmd := &command{
		Id:      msg.id,
		Name:    msg.name,
		Payload: msg.command,
		CommandResult: &CommandResult{
			Status:  CommandStatus_Queued,
			Updated: time.Now(),
			Data:    nil,
		},
	}

	this.commands[cmd.Id] = cmd

	this.queue <- cmd.Id
	msg.out <- nil
}

func (this *commandStore) onPop(msg *popMessage) {

	go func() {

		id := <-this.queue

		this.input <- &qMessage{
			msgType: qMessage_command,
			msg: &commandMessage{
				id:  id,
				out: msg.out,
			},
		}
	}()
}

func (this *commandStore) onCommand(msg *commandMessage) {

	cmd := this.commands[msg.id]
	transaction := NewTransaction(this.input, cmd.Id, cmd.Name, cmd.Payload)

	msg.out <- transaction
}

func (this *commandStore) onCheck(msg *checkMessage) {

	if cmd, ok := this.commands[msg.id]; ok {
		msg.out <- cmd.CommandResult
	} else {
		msg.out <- nil
	}
}

func (this *commandStore) onCommit(msg *commitMessage) {

	cmd := this.commands[msg.id]
	cmd.Status = CommandStatus_Ok
	cmd.Updated = time.Now()
	cmd.Data = msg.data

	go this.expire(msg.id)
}

func (this *commandStore) onFailed(msg *failedMessage) {

	cmd := this.commands[msg.id]
	cmd.Status = CommandStatus_Error
	cmd.Updated = time.Now()
	cmd.Data = msg.err.Error()

	go this.expire(msg.id)
}

func (this *commandStore) expire(id CommandId) {

	//this should probably send a expired message
	time.Sleep(this.expiry)
	delete(this.commands, id)
}

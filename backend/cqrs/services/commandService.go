package services

import (
	"fmt"
	. "github.com/mfelicio/einvite/backend/cqrs"
	"log"
	"time"
)

//Create:
// Load aggregate and call method on Aggregate

//Reads:
// Get view from cache
//   If exists, compare cached version to aggregate version
//	 	If true, return cached view
//		Else load events since cacheVersion, update view and return it
//			Then update cached entry async
//	 Else
//		Load all events, create view with a TTL and return it
//			Then create cache entry async

type CommandServiceOptions struct {
	TransactionsPerSecond int
	BackOffDuration       time.Duration
}

type CommandHandler func(command interface{}) (interface{}, error)

type CommandService struct {
	options  CommandServiceOptions
	store    CommandStore
	handlers map[string]CommandHandler
}

func NewCommandService(options CommandServiceOptions, store CommandStore) *CommandService {

	return &CommandService{
		options:  options,
		store:    store,
		handlers: make(map[string]CommandHandler),
	}
}

func (this *CommandService) Bind(name string, handler CommandHandler) {

	if _, exists := this.handlers[name]; exists {
		panic(fmt.Sprintf("There is already a CommandHandler registered for %s command", name))
	}

	this.handlers[name] = handler
}

func (this *CommandService) Start() {

	log.Println("CommandService started")

	for {

		if trans, err := this.store.Pop(); err == nil {

			// process command transaction
			go this.process(trans)

		} else {

			log.Println("could not retrieve command, backing off for %s. Error: %s", this.options.BackOffDuration, err.Error())
			time.Sleep(this.options.BackOffDuration)
		}

	}
}

func (this *CommandService) process(trans CommandTransaction) {

	name := trans.Name()
	cmd := trans.Command()

	handler := this.handlers[name]

	data, err := handler(cmd)

	if err == nil {

		trans.Commit(data)

	} else {

		trans.Failed(err)
	}

}

func (this *CommandService) Stop() {

	log.Println("CommandService stopped")
}

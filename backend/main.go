package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/mfelicio/einvite/backend/api"
	"github.com/mfelicio/einvite/backend/cqrs/memory"
	"github.com/mfelicio/einvite/backend/cqrs/services"
	"github.com/mfelicio/einvite/backend/cqrs/sourcing"
	"github.com/mfelicio/einvite/backend/domain/commands"
	"github.com/mfelicio/einvite/backend/domain/model"
	domainServices "github.com/mfelicio/einvite/backend/domain/services"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {

	log.Println("Setting GOMAXPROCS = ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmdStore := memory.NewCommandStore(10000) // config: max queued commands is 10k
	eventBus := memory.NewEventBus(false)     // config: false means not async

	cmdServiceOptions := services.CommandServiceOptions{
		TransactionsPerSecond: 100,
		BackOffDuration:       1 * time.Second,
	}
	cmdService := services.NewCommandService(cmdServiceOptions, cmdStore)

	framework := sourcing.NewFramework(memory.NewEventStore())

	refSvc := domainServices.NewDummyMeetingService()
	model.BindEvents(framework, refSvc)

	commands.BindHandlers(framework, cmdService, eventBus)

	api.Register(cmdStore)

	go cmdService.Start()
	defer cmdService.Stop()

	initSwagger()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initSwagger() {
	config := swagger.Config{
		WebServices:    restful.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://einvite.cloudapp.net",   //should come from config
		ApiPath:        "/api-docs",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/docs/",
		SwaggerFilePath: "api/swagger-ui"}
	swagger.InstallSwaggerService(config)
}

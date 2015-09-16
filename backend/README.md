Backend
=======

The backend follows a distributed DDD / CQRS / ES architecture that allows horizontal scalability on each component.

The main components and responsibilities are:

* REST API - interface for mobile and web clients
    * Controllers
	* Security
* Command Services - updates domain model
    * Command Handlers
	* Domain model
	* Event Store
	    * Dynamodb
* Query Services - reads cached data
    * Query Handlers
	* Event Handlers
	* Query Store
	    * Redis
* Messaging
	* CommandBus - integration between REST API and CommandServices
	* EventBus - integration between CommandServices and QueryServices

These components can be hosted separately or on the same process, for simplicity, using in-memory implementations.


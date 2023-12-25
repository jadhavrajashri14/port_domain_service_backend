# port_domain_service_backend

##How to run the application locally on ubuntu machine :
------------------------------------------------------
after downloading the code from github,
run following commands in the folder where you copied the code

for e.g. if you copied to /home/rajashrijadhav/RajashriJadhavData/RAJASHRI_ASSIGNMENTS/port_domain_service_backend
then in that folder

1. run "make build" to create the application binary in /app/ folder

2. run "docker build -t port-domain-service ."  -- This will create the docker image
3. check the docker image with "docker images" command
4. run "docker run -it --rm -p 8080:8080 port-domain-service"  -- this will start the port_domain_service server
You will get the output below,
INFO[0000] initializing routes ------
INFO[0000] main goroutine waiting for api call -->       state="port_domain_service server"
INFO[0000] starting server on :8080                      state="port_domain_service server"

5. Then use POSTMAN to call the create and update endpoints of this server.
The screenshots of postman are attached in the mail.

  1. POSTMAN /create endpoint

	POST - http://127.0.0.1:8080/create
	-- upload the ports.json file in the Body with form-data
	-- jsonFile - location of ports.json file
	-- after executing this endpoint you get
	-- Status 201 Created

  2. POSTMAN /update endpoint

 	PATCH - http://127.0.0.1:8080/update
	-- Inside Body - raw - select JSON option
	supply the port to be updated as below
	-- {"USSEA":{"name":"Cape Romanzof","city":"Cape Romanzof","province":"Alaska","country":"United Arab Emirates","alias":["Tacoma"],"regions":[],"coordinates":[55.2756505, 25.284755],"timezone":"America/Anchorage","unlocs":["AEPRA"],"code":"3001"}}

	-- Status 200 OK

----------------------------
## Build Tools used

Framework used : net/http to create the endpoints
in-memory map is used to store the ports of the json file
testing - github.com/smartystreets/goconvey/convey package used
validation - done inside the code where endpoint functions are defined
signal handling - done in the cmd/port_domain_service.go file. This is the main() function of the server.
(e.g. a TERM or KILL signal should result in a graceful shutdown). -- DONE
Code structure according to hexagonal architecture  -- DONE
A Dockerfile should be used to contain and run the service (Hint: extra points for avoiding code building in docker). -- DONE
readme.md should explain how to run your program and test it. -- DONE

-------------------------------
## Build Tools used
-------------------------------
go-junit-report  - go-junit-report reads the go test verbose output from standard in and writes junit compatible XML to standard out.

gocov - Package gocov is a code coverage analysis tool for Go.

Swaggo  - this is not used in the code, but I know how to use this, used frequently in the endpoint which I developed.
Swaggo is a tool that creates Swagger documentation for Go APIs. It makes documenting API endpoints easier, helping developers understand and use the API.

Golangci-lint
Golangci-lint is a tool for checking Go code quality, finding issues, bugs, and style problems. It helps keep the code clean and maintainable.

-------------------------------
## Code sturcture
-------------------------------
Domain - internal/core/domain/model.go
this has PortDomains model struct.

Ports - internal/core/ports/ports.go
this has interfaces which define app business logic for PortDomainRepository and PortDomainsService. These interfaces willl be implemented by adapters.

Services - internal/core/services/services.go
the services module establishes connection between core and the outside world.  This has PortDomainsService struct that implements interfaces ofPortDomainsService in the ports module like CreatePortDomain and UpdatePortDomain - calls repo.Create and Update

Adapters - internal/adapters/handler/server.go
this defines the server which defines the routes of the application and keeps on listening on the port :8080

Adapters - internal/adapters/handler/http.go
this defines the endpoints of the port_domain_service application

repository - internal/adapters/repository/inmemdb.go
this defines the in memory map to be used for storing ports data and contains the repository functions to store and update the ports data.

Makefile and Dockerfile is also provided.


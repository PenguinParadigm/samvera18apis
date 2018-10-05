package main

//go:generate go run generate/generate.go

import (
	"log"
	"net/http"

	"github.com/PenguinParadigm/samvera18apis/taquito/aws_session"
	"github.com/PenguinParadigm/samvera18apis/taquito/config"
	"github.com/PenguinParadigm/samvera18apis/taquito/db"
	"github.com/PenguinParadigm/samvera18apis/taquito/generated/restapi"
	"github.com/PenguinParadigm/samvera18apis/taquito/generated/restapi/operations"
	"github.com/PenguinParadigm/samvera18apis/taquito/handlers"
	"github.com/PenguinParadigm/samvera18apis/taquito/identifier"
	"github.com/PenguinParadigm/samvera18apis/taquito/middleware"
	"github.com/justinas/alice"
)

func main() {
	// Initialize our global struct
	config := config.NewConfig()
	awsSession := aws_session.Connect(config.AwsDisableSSL)
	database := &db.DynamodbDatabase{
		Connection: db.Connect(awsSession, config.DynamodbEndpoint),
		Table:      config.ResourceTableName,
	}

	identifierService := identifier.NewService(config)
	server := createServer(database, identifierService, config.Port)
	defer server.Shutdown()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer(database db.Database, identifierService identifier.Service, port int) *restapi.Server {
	api := handlers.BuildAPI(database, identifierService)
	server := restapi.NewServer(api)
	server.SetHandler(BuildHandler(api))

	// set the port this service will be run on
	server.Port = port
	return server
}

// BuildHandler sets up the middleware that wraps the API
func BuildHandler(api *operations.TacoAPI) http.Handler {
	return alice.New(
		middleware.NewRecoveryMW(),
		middleware.NewRequestLoggerMW(),
	).Then(api.Serve(nil))
}

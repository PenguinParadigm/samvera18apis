package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/PenguinParadigm/samvera18apis/simpleDev/generated/models"
	"github.com/PenguinParadigm/samvera18apis/simpleDev/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// NewRetrieveResource will query DynamoDB with ID for Resource JSON
func NewRetrieveResource() operations.RetrieveResourceHandler {
	return &retrieveResource{}
}

// retrieveResource handles a request for finding & returning an entry
type retrieveResource struct{}

// Handle the retrieve resource request
func (d *retrieveResource) Handle(params operations.RetrieveResourceParams, agent *models.Agent) middleware.Responder {
	var dataResource models.Resource
	var jsonFile []byte
	jsonFile, err := ioutil.ReadFile("examples/update_request.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonFile, &dataResource)
	return operations.NewRetrieveResourceOK().WithPayload(&dataResource)
}

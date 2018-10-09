package handlers

import (
	"time"

	"github.com/PenguinParadigm/samvera18apis/simpleDev/generated/models"
	"github.com/PenguinParadigm/samvera18apis/simpleDev/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// NewDepositResource -- Accepts requests to create resource.
func NewDepositResource() operations.DepositResourceHandler {
	return &depositResource{}
}

type depositResource struct{}

// Handle the create resource request
func (d *depositResource) Handle(params operations.DepositResourceParams, agent *models.Agent) middleware.Responder {
	// Grab the data passed in from the call
	// The params is an `operations.DepositResourceParams`, which is defined in the
	// `generated/restapi/operations/deposit_resource_parameters.go` file
	resource := params.Payload

	// Mint an identifier (UUID)
	uuid, err := uuid.NewRandom()
	if err != nil {
		// Explode. I don't know what to do if you can't generate a random number
		panic(err)
	}

	// Add administrative metadata
	var identifier strfmt.UUID
	identifier = strfmt.UUID(uuid.String())
	resource.Identifier = identifier
	resource.Depositor = agent
	// Do a declaration ahead of time; instead of `var version int64; version = int64(1)`, we could do `version := int64(1)
	var version int64
	version = int64(1)
	resource.Version = version
	var datestamp strfmt.DateTime
	datestamp = strfmt.DateTime(time.Now().UTC())
	resource.Administrative.Created = datestamp
	resource.Administrative.LastUpdated = datestamp

	// Build the API response
	response := &models.ResourceResponse{Identifier: &identifier, Data: &resource}
	return operations.NewDepositResourceCreated().WithPayload(response)
}

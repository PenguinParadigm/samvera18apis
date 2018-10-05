package validators

import (
	"github.com/PenguinParadigm/samvera18apis/taquito/datautils"
	"github.com/PenguinParadigm/samvera18apis/taquito/generated/models"
)

// ResourceValidator is the interface for validators that check the resources format
type ResourceValidator interface {
	ValidateResource(resource *datautils.Resource) *models.ErrorResponseErrors
}

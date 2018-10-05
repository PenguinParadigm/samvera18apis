package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/PenguinParadigm/samvera18apis/taquito/datautils"
	"github.com/PenguinParadigm/samvera18apis/taquito/generated/models"
)

type mockValidator struct {
	returns bool
}

func (d *mockValidator) ValidateResource(*datautils.Resource) *models.ErrorResponseErrors {
	if d.returns {
		return nil
	}
	return &models.ErrorResponseErrors{}
}

func TestCompositeResourceIsInvalid(t *testing.T) {
	v1 := &mockValidator{returns: true}
	v2 := &mockValidator{returns: false}

	compositeValidator := NewCompositeResourceValidator([]ResourceValidator{v1, v2})
	err := compositeValidator.ValidateResource(testResource("bs646cd8717.json"))
	assert.NotNil(t, err)
}

func TestCompositeResourceIsValid(t *testing.T) {
	v1 := &mockValidator{returns: true}
	v2 := &mockValidator{returns: true}

	compositeValidator := NewCompositeResourceValidator([]ResourceValidator{v1, v2})
	err := compositeValidator.ValidateResource(testResource("bs646cd8717.json"))
	assert.Nil(t, err)
}

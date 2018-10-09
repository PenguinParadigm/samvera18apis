// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PenguinParadigm/samvera18apis/simpleDev/generated/models"
)

// NewUpdateResourceParams creates a new UpdateResourceParams object
// no default values defined in spec.
func NewUpdateResourceParams() UpdateResourceParams {

	return UpdateResourceParams{}
}

// UpdateResourceParams contains all the bound params for the update resource operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateResource
type UpdateResourceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*SDR Identifier for the Resource.
	  Required: true
	  In: path
	*/
	ID string
	/*JSON-LD Representation of the resource metadata required fields and only the fields you wish to update (identified via its TACO identifier). Needs to fit the SDR 3.0 MAP requirements.
	  Required: true
	  In: body
	*/
	Payload *models.Resource
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateResourceParams() beforehand.
func (o *UpdateResourceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("ID")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Resource
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("payload", "body"))
			} else {
				res = append(res, errors.NewParseError("payload", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Payload = &body
			}
		}
	} else {
		res = append(res, errors.Required("payload", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdateResourceParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

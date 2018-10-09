// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/PenguinParadigm/samvera18apis/simpleDev/generated/models"
)

// DepositResourceHandlerFunc turns a function with the right signature into a deposit resource handler
type DepositResourceHandlerFunc func(DepositResourceParams, *models.Agent) middleware.Responder

// Handle executing the request and returning a response
func (fn DepositResourceHandlerFunc) Handle(params DepositResourceParams, principal *models.Agent) middleware.Responder {
	return fn(params, principal)
}

// DepositResourceHandler interface for that can handle valid deposit resource params
type DepositResourceHandler interface {
	Handle(DepositResourceParams, *models.Agent) middleware.Responder
}

// NewDepositResource creates a new http.Handler for the deposit resource operation
func NewDepositResource(ctx *middleware.Context, handler DepositResourceHandler) *DepositResource {
	return &DepositResource{Context: ctx, Handler: handler}
}

/*DepositResource swagger:route POST /resource depositResource

Deposit New TAQUITO Resource.

Deposits a new resource into our digital repository. Returns the domain identifier for the resource.

*/
type DepositResource struct {
	Context *middleware.Context
	Handler DepositResourceHandler
}

func (o *DepositResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDepositResourceParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Agent
	if uprinc != nil {
		principal = uprinc.(*models.Agent) // this is really a models.Agent, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

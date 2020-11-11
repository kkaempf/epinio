// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllServiceBindingsHandlerFunc turns a function with the right signature into a list all service bindings handler
type ListAllServiceBindingsHandlerFunc func(ListAllServiceBindingsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllServiceBindingsHandlerFunc) Handle(params ListAllServiceBindingsParams) middleware.Responder {
	return fn(params)
}

// ListAllServiceBindingsHandler interface for that can handle valid list all service bindings params
type ListAllServiceBindingsHandler interface {
	Handle(ListAllServiceBindingsParams) middleware.Responder
}

// NewListAllServiceBindings creates a new http.Handler for the list all service bindings operation
func NewListAllServiceBindings(ctx *middleware.Context, handler ListAllServiceBindingsHandler) *ListAllServiceBindings {
	return &ListAllServiceBindings{Context: ctx, Handler: handler}
}

/*ListAllServiceBindings swagger:route GET /service_bindings serviceBindings listAllServiceBindings

List all Service Bindings

curl --insecure -i %s/v2/service_bindings -X GET -H 'Authorization: %s'

*/
type ListAllServiceBindings struct {
	Context *middleware.Context
	Handler ListAllServiceBindingsHandler
}

func (o *ListAllServiceBindings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllServiceBindingsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

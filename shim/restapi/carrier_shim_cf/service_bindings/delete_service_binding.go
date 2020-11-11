// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteServiceBindingHandlerFunc turns a function with the right signature into a delete service binding handler
type DeleteServiceBindingHandlerFunc func(DeleteServiceBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteServiceBindingHandlerFunc) Handle(params DeleteServiceBindingParams) middleware.Responder {
	return fn(params)
}

// DeleteServiceBindingHandler interface for that can handle valid delete service binding params
type DeleteServiceBindingHandler interface {
	Handle(DeleteServiceBindingParams) middleware.Responder
}

// NewDeleteServiceBinding creates a new http.Handler for the delete service binding operation
func NewDeleteServiceBinding(ctx *middleware.Context, handler DeleteServiceBindingHandler) *DeleteServiceBinding {
	return &DeleteServiceBinding{Context: ctx, Handler: handler}
}

/*DeleteServiceBinding swagger:route DELETE /service_bindings/{guid} serviceBindings deleteServiceBinding

Delete a Particular Service Binding

curl --insecure -i %s/v2/service_bindings/{guid} -X DELETE -H 'Authorization: %s'

*/
type DeleteServiceBinding struct {
	Context *middleware.Context
	Handler DeleteServiceBindingHandler
}

func (o *DeleteServiceBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteServiceBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

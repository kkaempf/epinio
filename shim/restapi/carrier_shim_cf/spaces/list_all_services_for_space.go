// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllServicesForSpaceHandlerFunc turns a function with the right signature into a list all services for space handler
type ListAllServicesForSpaceHandlerFunc func(ListAllServicesForSpaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllServicesForSpaceHandlerFunc) Handle(params ListAllServicesForSpaceParams) middleware.Responder {
	return fn(params)
}

// ListAllServicesForSpaceHandler interface for that can handle valid list all services for space params
type ListAllServicesForSpaceHandler interface {
	Handle(ListAllServicesForSpaceParams) middleware.Responder
}

// NewListAllServicesForSpace creates a new http.Handler for the list all services for space operation
func NewListAllServicesForSpace(ctx *middleware.Context, handler ListAllServicesForSpaceHandler) *ListAllServicesForSpace {
	return &ListAllServicesForSpace{Context: ctx, Handler: handler}
}

/*ListAllServicesForSpace swagger:route GET /spaces/{guid}/services spaces listAllServicesForSpace

List all Services for the Space

curl --insecure -i %s/v2/spaces/{guid}/services -X GET -H 'Authorization: %s'

*/
type ListAllServicesForSpace struct {
	Context *middleware.Context
	Handler ListAllServicesForSpaceHandler
}

func (o *ListAllServicesForSpace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllServicesForSpaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// Code generated by go-swagger; DO NOT EDIT.

package resource_match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllMatchingResourcesHandlerFunc turns a function with the right signature into a list all matching resources handler
type ListAllMatchingResourcesHandlerFunc func(ListAllMatchingResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllMatchingResourcesHandlerFunc) Handle(params ListAllMatchingResourcesParams) middleware.Responder {
	return fn(params)
}

// ListAllMatchingResourcesHandler interface for that can handle valid list all matching resources params
type ListAllMatchingResourcesHandler interface {
	Handle(ListAllMatchingResourcesParams) middleware.Responder
}

// NewListAllMatchingResources creates a new http.Handler for the list all matching resources operation
func NewListAllMatchingResources(ctx *middleware.Context, handler ListAllMatchingResourcesHandler) *ListAllMatchingResources {
	return &ListAllMatchingResources{Context: ctx, Handler: handler}
}

/*ListAllMatchingResources swagger:route PUT /resource_match resourceMatch listAllMatchingResources

List all matching resources

curl --insecure -i %s/v2/resource_match -X PUT -H 'Authorization: %s' -d '%s'

*/
type ListAllMatchingResources struct {
	Context *middleware.Context
	Handler ListAllMatchingResourcesHandler
}

func (o *ListAllMatchingResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllMatchingResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

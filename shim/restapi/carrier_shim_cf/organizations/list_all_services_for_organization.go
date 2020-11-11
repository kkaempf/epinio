// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllServicesForOrganizationHandlerFunc turns a function with the right signature into a list all services for organization handler
type ListAllServicesForOrganizationHandlerFunc func(ListAllServicesForOrganizationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllServicesForOrganizationHandlerFunc) Handle(params ListAllServicesForOrganizationParams) middleware.Responder {
	return fn(params)
}

// ListAllServicesForOrganizationHandler interface for that can handle valid list all services for organization params
type ListAllServicesForOrganizationHandler interface {
	Handle(ListAllServicesForOrganizationParams) middleware.Responder
}

// NewListAllServicesForOrganization creates a new http.Handler for the list all services for organization operation
func NewListAllServicesForOrganization(ctx *middleware.Context, handler ListAllServicesForOrganizationHandler) *ListAllServicesForOrganization {
	return &ListAllServicesForOrganization{Context: ctx, Handler: handler}
}

/*ListAllServicesForOrganization swagger:route GET /organizations/{guid}/services organizations listAllServicesForOrganization

List all Services for the Organization

curl --insecure -i %s/v2/organizations/{guid}/services -X GET -H 'Authorization: %s'

*/
type ListAllServicesForOrganization struct {
	Context *middleware.Context
	Handler ListAllServicesForOrganizationHandler
}

func (o *ListAllServicesForOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllServicesForOrganizationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

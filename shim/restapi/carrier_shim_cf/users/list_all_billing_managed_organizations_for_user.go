// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllBillingManagedOrganizationsForUserHandlerFunc turns a function with the right signature into a list all billing managed organizations for user handler
type ListAllBillingManagedOrganizationsForUserHandlerFunc func(ListAllBillingManagedOrganizationsForUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllBillingManagedOrganizationsForUserHandlerFunc) Handle(params ListAllBillingManagedOrganizationsForUserParams) middleware.Responder {
	return fn(params)
}

// ListAllBillingManagedOrganizationsForUserHandler interface for that can handle valid list all billing managed organizations for user params
type ListAllBillingManagedOrganizationsForUserHandler interface {
	Handle(ListAllBillingManagedOrganizationsForUserParams) middleware.Responder
}

// NewListAllBillingManagedOrganizationsForUser creates a new http.Handler for the list all billing managed organizations for user operation
func NewListAllBillingManagedOrganizationsForUser(ctx *middleware.Context, handler ListAllBillingManagedOrganizationsForUserHandler) *ListAllBillingManagedOrganizationsForUser {
	return &ListAllBillingManagedOrganizationsForUser{Context: ctx, Handler: handler}
}

/*ListAllBillingManagedOrganizationsForUser swagger:route GET /users/{guid}/billing_managed_organizations users listAllBillingManagedOrganizationsForUser

List all Billing Managed Organizations for the User

curl --insecure -i %s/v2/users/{guid}/billing_managed_organizations -X GET -H 'Authorization: %s'

*/
type ListAllBillingManagedOrganizationsForUser struct {
	Context *middleware.Context
	Handler ListAllBillingManagedOrganizationsForUserHandler
}

func (o *ListAllBillingManagedOrganizationsForUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllBillingManagedOrganizationsForUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

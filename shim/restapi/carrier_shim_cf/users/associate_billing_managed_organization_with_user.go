// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AssociateBillingManagedOrganizationWithUserHandlerFunc turns a function with the right signature into a associate billing managed organization with user handler
type AssociateBillingManagedOrganizationWithUserHandlerFunc func(AssociateBillingManagedOrganizationWithUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociateBillingManagedOrganizationWithUserHandlerFunc) Handle(params AssociateBillingManagedOrganizationWithUserParams) middleware.Responder {
	return fn(params)
}

// AssociateBillingManagedOrganizationWithUserHandler interface for that can handle valid associate billing managed organization with user params
type AssociateBillingManagedOrganizationWithUserHandler interface {
	Handle(AssociateBillingManagedOrganizationWithUserParams) middleware.Responder
}

// NewAssociateBillingManagedOrganizationWithUser creates a new http.Handler for the associate billing managed organization with user operation
func NewAssociateBillingManagedOrganizationWithUser(ctx *middleware.Context, handler AssociateBillingManagedOrganizationWithUserHandler) *AssociateBillingManagedOrganizationWithUser {
	return &AssociateBillingManagedOrganizationWithUser{Context: ctx, Handler: handler}
}

/*AssociateBillingManagedOrganizationWithUser swagger:route PUT /users/{guid}/billing_managed_organizations/{billing_managed_organization_guid} users associateBillingManagedOrganizationWithUser

Associate Billing Managed Organization with the User

curl --insecure -i %s/v2/users/{guid}/billing_managed_organizations/{billing_managed_organization_guid} -X PUT -H 'Authorization: %s'

*/
type AssociateBillingManagedOrganizationWithUser struct {
	Context *middleware.Context
	Handler AssociateBillingManagedOrganizationWithUserHandler
}

func (o *AssociateBillingManagedOrganizationWithUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociateBillingManagedOrganizationWithUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

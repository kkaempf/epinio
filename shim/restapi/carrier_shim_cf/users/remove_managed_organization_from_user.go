// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveManagedOrganizationFromUserHandlerFunc turns a function with the right signature into a remove managed organization from user handler
type RemoveManagedOrganizationFromUserHandlerFunc func(RemoveManagedOrganizationFromUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveManagedOrganizationFromUserHandlerFunc) Handle(params RemoveManagedOrganizationFromUserParams) middleware.Responder {
	return fn(params)
}

// RemoveManagedOrganizationFromUserHandler interface for that can handle valid remove managed organization from user params
type RemoveManagedOrganizationFromUserHandler interface {
	Handle(RemoveManagedOrganizationFromUserParams) middleware.Responder
}

// NewRemoveManagedOrganizationFromUser creates a new http.Handler for the remove managed organization from user operation
func NewRemoveManagedOrganizationFromUser(ctx *middleware.Context, handler RemoveManagedOrganizationFromUserHandler) *RemoveManagedOrganizationFromUser {
	return &RemoveManagedOrganizationFromUser{Context: ctx, Handler: handler}
}

/*RemoveManagedOrganizationFromUser swagger:route DELETE /users/{guid}/managed_organizations/{managed_organization_guid} users removeManagedOrganizationFromUser

Remove Managed Organization from the User

curl --insecure -i %s/v2/users/{guid}/managed_organizations/{managed_organization_guid} -X DELETE -H 'Authorization: %s'

*/
type RemoveManagedOrganizationFromUser struct {
	Context *middleware.Context
	Handler RemoveManagedOrganizationFromUserHandler
}

func (o *RemoveManagedOrganizationFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveManagedOrganizationFromUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

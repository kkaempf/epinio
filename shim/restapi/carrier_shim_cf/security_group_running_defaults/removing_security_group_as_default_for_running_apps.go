// Code generated by go-swagger; DO NOT EDIT.

package security_group_running_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemovingSecurityGroupAsDefaultForRunningAppsHandlerFunc turns a function with the right signature into a removing security group as default for running apps handler
type RemovingSecurityGroupAsDefaultForRunningAppsHandlerFunc func(RemovingSecurityGroupAsDefaultForRunningAppsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemovingSecurityGroupAsDefaultForRunningAppsHandlerFunc) Handle(params RemovingSecurityGroupAsDefaultForRunningAppsParams) middleware.Responder {
	return fn(params)
}

// RemovingSecurityGroupAsDefaultForRunningAppsHandler interface for that can handle valid removing security group as default for running apps params
type RemovingSecurityGroupAsDefaultForRunningAppsHandler interface {
	Handle(RemovingSecurityGroupAsDefaultForRunningAppsParams) middleware.Responder
}

// NewRemovingSecurityGroupAsDefaultForRunningApps creates a new http.Handler for the removing security group as default for running apps operation
func NewRemovingSecurityGroupAsDefaultForRunningApps(ctx *middleware.Context, handler RemovingSecurityGroupAsDefaultForRunningAppsHandler) *RemovingSecurityGroupAsDefaultForRunningApps {
	return &RemovingSecurityGroupAsDefaultForRunningApps{Context: ctx, Handler: handler}
}

/*RemovingSecurityGroupAsDefaultForRunningApps swagger:route DELETE /config/running_security_groups/{guid} securityGroupRunningDefaults removingSecurityGroupAsDefaultForRunningApps

Removing a Security Group as a default for running Apps

curl --insecure -i %s/v2/config/running_security_groups/{guid} -X DELETE -H 'Authorization: %s'

*/
type RemovingSecurityGroupAsDefaultForRunningApps struct {
	Context *middleware.Context
	Handler RemovingSecurityGroupAsDefaultForRunningAppsHandler
}

func (o *RemovingSecurityGroupAsDefaultForRunningApps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemovingSecurityGroupAsDefaultForRunningAppsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

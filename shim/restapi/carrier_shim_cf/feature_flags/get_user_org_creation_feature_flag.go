// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserOrgCreationFeatureFlagHandlerFunc turns a function with the right signature into a get user org creation feature flag handler
type GetUserOrgCreationFeatureFlagHandlerFunc func(GetUserOrgCreationFeatureFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserOrgCreationFeatureFlagHandlerFunc) Handle(params GetUserOrgCreationFeatureFlagParams) middleware.Responder {
	return fn(params)
}

// GetUserOrgCreationFeatureFlagHandler interface for that can handle valid get user org creation feature flag params
type GetUserOrgCreationFeatureFlagHandler interface {
	Handle(GetUserOrgCreationFeatureFlagParams) middleware.Responder
}

// NewGetUserOrgCreationFeatureFlag creates a new http.Handler for the get user org creation feature flag operation
func NewGetUserOrgCreationFeatureFlag(ctx *middleware.Context, handler GetUserOrgCreationFeatureFlagHandler) *GetUserOrgCreationFeatureFlag {
	return &GetUserOrgCreationFeatureFlag{Context: ctx, Handler: handler}
}

/*GetUserOrgCreationFeatureFlag swagger:route GET /config/feature_flags/user_org_creation featureFlags getUserOrgCreationFeatureFlag

Get the User Org Creation feature flag

curl --insecure -i %s/v2/config/feature_flags/user_org_creation -X GET -H 'Authorization: %s'

*/
type GetUserOrgCreationFeatureFlag struct {
	Context *middleware.Context
	Handler GetUserOrgCreationFeatureFlagHandler
}

func (o *GetUserOrgCreationFeatureFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserOrgCreationFeatureFlagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

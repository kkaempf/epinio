// Code generated by go-swagger; DO NOT EDIT.

package space_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveSpaceFromSpaceQuotaDefinitionHandlerFunc turns a function with the right signature into a remove space from space quota definition handler
type RemoveSpaceFromSpaceQuotaDefinitionHandlerFunc func(RemoveSpaceFromSpaceQuotaDefinitionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveSpaceFromSpaceQuotaDefinitionHandlerFunc) Handle(params RemoveSpaceFromSpaceQuotaDefinitionParams) middleware.Responder {
	return fn(params)
}

// RemoveSpaceFromSpaceQuotaDefinitionHandler interface for that can handle valid remove space from space quota definition params
type RemoveSpaceFromSpaceQuotaDefinitionHandler interface {
	Handle(RemoveSpaceFromSpaceQuotaDefinitionParams) middleware.Responder
}

// NewRemoveSpaceFromSpaceQuotaDefinition creates a new http.Handler for the remove space from space quota definition operation
func NewRemoveSpaceFromSpaceQuotaDefinition(ctx *middleware.Context, handler RemoveSpaceFromSpaceQuotaDefinitionHandler) *RemoveSpaceFromSpaceQuotaDefinition {
	return &RemoveSpaceFromSpaceQuotaDefinition{Context: ctx, Handler: handler}
}

/*RemoveSpaceFromSpaceQuotaDefinition swagger:route DELETE /space_quota_definitions/{guid}/spaces/{space_guid} spaceQuotaDefinitions removeSpaceFromSpaceQuotaDefinition

Remove Space from the Space Quota Definition

curl --insecure -i %s/v2/space_quota_definitions/{guid}/spaces/{space_guid} -X DELETE -H 'Authorization: %s'

*/
type RemoveSpaceFromSpaceQuotaDefinition struct {
	Context *middleware.Context
	Handler RemoveSpaceFromSpaceQuotaDefinitionHandler
}

func (o *RemoveSpaceFromSpaceQuotaDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveSpaceFromSpaceQuotaDefinitionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

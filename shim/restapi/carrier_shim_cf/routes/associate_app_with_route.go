// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AssociateAppWithRouteHandlerFunc turns a function with the right signature into a associate app with route handler
type AssociateAppWithRouteHandlerFunc func(AssociateAppWithRouteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociateAppWithRouteHandlerFunc) Handle(params AssociateAppWithRouteParams) middleware.Responder {
	return fn(params)
}

// AssociateAppWithRouteHandler interface for that can handle valid associate app with route params
type AssociateAppWithRouteHandler interface {
	Handle(AssociateAppWithRouteParams) middleware.Responder
}

// NewAssociateAppWithRoute creates a new http.Handler for the associate app with route operation
func NewAssociateAppWithRoute(ctx *middleware.Context, handler AssociateAppWithRouteHandler) *AssociateAppWithRoute {
	return &AssociateAppWithRoute{Context: ctx, Handler: handler}
}

/*AssociateAppWithRoute swagger:route PUT /routes/{guid}/apps/{app_guid} routes associateAppWithRoute

Associate App with the Route

curl --insecure -i %s/v2/routes/{guid}/apps/{app_guid} -X PUT -H 'Authorization: %s'

*/
type AssociateAppWithRoute struct {
	Context *middleware.Context
	Handler AssociateAppWithRouteHandler
}

func (o *AssociateAppWithRoute) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociateAppWithRouteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// Code generated by go-swagger; DO NOT EDIT.

package buildpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteBuildpackHandlerFunc turns a function with the right signature into a delete buildpack handler
type DeleteBuildpackHandlerFunc func(DeleteBuildpackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBuildpackHandlerFunc) Handle(params DeleteBuildpackParams) middleware.Responder {
	return fn(params)
}

// DeleteBuildpackHandler interface for that can handle valid delete buildpack params
type DeleteBuildpackHandler interface {
	Handle(DeleteBuildpackParams) middleware.Responder
}

// NewDeleteBuildpack creates a new http.Handler for the delete buildpack operation
func NewDeleteBuildpack(ctx *middleware.Context, handler DeleteBuildpackHandler) *DeleteBuildpack {
	return &DeleteBuildpack{Context: ctx, Handler: handler}
}

/*DeleteBuildpack swagger:route DELETE /buildpacks/{guid} buildpacks deleteBuildpack

Delete a Particular Buildpack

curl --insecure -i %s/v2/buildpacks/{guid} -X DELETE -H 'Authorization: %s'

*/
type DeleteBuildpack struct {
	Context *middleware.Context
	Handler DeleteBuildpackHandler
}

func (o *DeleteBuildpack) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBuildpackParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

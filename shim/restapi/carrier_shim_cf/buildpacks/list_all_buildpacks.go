// Code generated by go-swagger; DO NOT EDIT.

package buildpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllBuildpacksHandlerFunc turns a function with the right signature into a list all buildpacks handler
type ListAllBuildpacksHandlerFunc func(ListAllBuildpacksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllBuildpacksHandlerFunc) Handle(params ListAllBuildpacksParams) middleware.Responder {
	return fn(params)
}

// ListAllBuildpacksHandler interface for that can handle valid list all buildpacks params
type ListAllBuildpacksHandler interface {
	Handle(ListAllBuildpacksParams) middleware.Responder
}

// NewListAllBuildpacks creates a new http.Handler for the list all buildpacks operation
func NewListAllBuildpacks(ctx *middleware.Context, handler ListAllBuildpacksHandler) *ListAllBuildpacks {
	return &ListAllBuildpacks{Context: ctx, Handler: handler}
}

/*ListAllBuildpacks swagger:route GET /buildpacks buildpacks listAllBuildpacks

List all Buildpacks

curl --insecure -i %s/v2/buildpacks -X GET -H 'Authorization: %s'

*/
type ListAllBuildpacks struct {
	Context *middleware.Context
	Handler ListAllBuildpacksHandler
}

func (o *ListAllBuildpacks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllBuildpacksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

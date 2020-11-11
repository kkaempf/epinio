// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateSpaceHandlerFunc turns a function with the right signature into a update space handler
type UpdateSpaceHandlerFunc func(UpdateSpaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateSpaceHandlerFunc) Handle(params UpdateSpaceParams) middleware.Responder {
	return fn(params)
}

// UpdateSpaceHandler interface for that can handle valid update space params
type UpdateSpaceHandler interface {
	Handle(UpdateSpaceParams) middleware.Responder
}

// NewUpdateSpace creates a new http.Handler for the update space operation
func NewUpdateSpace(ctx *middleware.Context, handler UpdateSpaceHandler) *UpdateSpace {
	return &UpdateSpace{Context: ctx, Handler: handler}
}

/*UpdateSpace swagger:route PUT /spaces/{guid} spaces updateSpace

Update a Space

curl --insecure -i %s/v2/spaces/{guid} -X PUT -H 'Authorization: %s' -d '%s'

*/
type UpdateSpace struct {
	Context *middleware.Context
	Handler UpdateSpaceHandler
}

func (o *UpdateSpace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateSpaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

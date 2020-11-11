// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSpaceSummaryHandlerFunc turns a function with the right signature into a get space summary handler
type GetSpaceSummaryHandlerFunc func(GetSpaceSummaryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpaceSummaryHandlerFunc) Handle(params GetSpaceSummaryParams) middleware.Responder {
	return fn(params)
}

// GetSpaceSummaryHandler interface for that can handle valid get space summary params
type GetSpaceSummaryHandler interface {
	Handle(GetSpaceSummaryParams) middleware.Responder
}

// NewGetSpaceSummary creates a new http.Handler for the get space summary operation
func NewGetSpaceSummary(ctx *middleware.Context, handler GetSpaceSummaryHandler) *GetSpaceSummary {
	return &GetSpaceSummary{Context: ctx, Handler: handler}
}

/*GetSpaceSummary swagger:route GET /spaces/{guid}/summary spaces getSpaceSummary

Get Space summary

curl --insecure -i %s/v2/spaces/{guid}/summary -X GET -H 'Authorization: %s'

*/
type GetSpaceSummary struct {
	Context *middleware.Context
	Handler GetSpaceSummaryHandler
}

func (o *GetSpaceSummary) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSpaceSummaryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEnvForAppHandlerFunc turns a function with the right signature into a get env for app handler
type GetEnvForAppHandlerFunc func(GetEnvForAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEnvForAppHandlerFunc) Handle(params GetEnvForAppParams) middleware.Responder {
	return fn(params)
}

// GetEnvForAppHandler interface for that can handle valid get env for app params
type GetEnvForAppHandler interface {
	Handle(GetEnvForAppParams) middleware.Responder
}

// NewGetEnvForApp creates a new http.Handler for the get env for app operation
func NewGetEnvForApp(ctx *middleware.Context, handler GetEnvForAppHandler) *GetEnvForApp {
	return &GetEnvForApp{Context: ctx, Handler: handler}
}

/*GetEnvForApp swagger:route GET /apps/{guid}/env apps getEnvForApp

Get the env for an App

curl --insecure -i %s/v2/apps/{guid}/env -X GET -H 'Authorization: %s'

*/
type GetEnvForApp struct {
	Context *middleware.Context
	Handler GetEnvForAppHandler
}

func (o *GetEnvForApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEnvForAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

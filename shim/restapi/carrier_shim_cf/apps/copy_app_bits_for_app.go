// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CopyAppBitsForAppHandlerFunc turns a function with the right signature into a copy app bits for app handler
type CopyAppBitsForAppHandlerFunc func(CopyAppBitsForAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CopyAppBitsForAppHandlerFunc) Handle(params CopyAppBitsForAppParams) middleware.Responder {
	return fn(params)
}

// CopyAppBitsForAppHandler interface for that can handle valid copy app bits for app params
type CopyAppBitsForAppHandler interface {
	Handle(CopyAppBitsForAppParams) middleware.Responder
}

// NewCopyAppBitsForApp creates a new http.Handler for the copy app bits for app operation
func NewCopyAppBitsForApp(ctx *middleware.Context, handler CopyAppBitsForAppHandler) *CopyAppBitsForApp {
	return &CopyAppBitsForApp{Context: ctx, Handler: handler}
}

/*CopyAppBitsForApp swagger:route POST /apps/{guid}/copy_bits apps copyAppBitsForApp

Copy the app bits for an App

curl --insecure -i %s/v2/apps/{guid}/copy_bits -X POST -H 'Authorization: %s' -d '%s'

*/
type CopyAppBitsForApp struct {
	Context *middleware.Context
	Handler CopyAppBitsForAppHandler
}

func (o *CopyAppBitsForApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCopyAppBitsForAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

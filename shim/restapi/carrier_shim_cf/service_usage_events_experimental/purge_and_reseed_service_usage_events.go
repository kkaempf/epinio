// Code generated by go-swagger; DO NOT EDIT.

package service_usage_events_experimental

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PurgeAndReseedServiceUsageEventsHandlerFunc turns a function with the right signature into a purge and reseed service usage events handler
type PurgeAndReseedServiceUsageEventsHandlerFunc func(PurgeAndReseedServiceUsageEventsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PurgeAndReseedServiceUsageEventsHandlerFunc) Handle(params PurgeAndReseedServiceUsageEventsParams) middleware.Responder {
	return fn(params)
}

// PurgeAndReseedServiceUsageEventsHandler interface for that can handle valid purge and reseed service usage events params
type PurgeAndReseedServiceUsageEventsHandler interface {
	Handle(PurgeAndReseedServiceUsageEventsParams) middleware.Responder
}

// NewPurgeAndReseedServiceUsageEvents creates a new http.Handler for the purge and reseed service usage events operation
func NewPurgeAndReseedServiceUsageEvents(ctx *middleware.Context, handler PurgeAndReseedServiceUsageEventsHandler) *PurgeAndReseedServiceUsageEvents {
	return &PurgeAndReseedServiceUsageEvents{Context: ctx, Handler: handler}
}

/*PurgeAndReseedServiceUsageEvents swagger:route POST /service_usage_events/destructively_purge_all_and_reseed_existing_instances serviceUsageEventsExperimental purgeAndReseedServiceUsageEvents

Purge and reseed Service Usage Events

curl --insecure -i %s/v2/service_usage_events/destructively_purge_all_and_reseed_existing_instances -X POST -H 'Authorization: %s'

*/
type PurgeAndReseedServiceUsageEvents struct {
	Context *middleware.Context
	Handler PurgeAndReseedServiceUsageEventsHandler
}

func (o *PurgeAndReseedServiceUsageEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPurgeAndReseedServiceUsageEventsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

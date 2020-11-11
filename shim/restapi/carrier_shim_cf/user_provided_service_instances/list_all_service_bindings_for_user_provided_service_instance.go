// Code generated by go-swagger; DO NOT EDIT.

package user_provided_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllServiceBindingsForUserProvidedServiceInstanceHandlerFunc turns a function with the right signature into a list all service bindings for user provided service instance handler
type ListAllServiceBindingsForUserProvidedServiceInstanceHandlerFunc func(ListAllServiceBindingsForUserProvidedServiceInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllServiceBindingsForUserProvidedServiceInstanceHandlerFunc) Handle(params ListAllServiceBindingsForUserProvidedServiceInstanceParams) middleware.Responder {
	return fn(params)
}

// ListAllServiceBindingsForUserProvidedServiceInstanceHandler interface for that can handle valid list all service bindings for user provided service instance params
type ListAllServiceBindingsForUserProvidedServiceInstanceHandler interface {
	Handle(ListAllServiceBindingsForUserProvidedServiceInstanceParams) middleware.Responder
}

// NewListAllServiceBindingsForUserProvidedServiceInstance creates a new http.Handler for the list all service bindings for user provided service instance operation
func NewListAllServiceBindingsForUserProvidedServiceInstance(ctx *middleware.Context, handler ListAllServiceBindingsForUserProvidedServiceInstanceHandler) *ListAllServiceBindingsForUserProvidedServiceInstance {
	return &ListAllServiceBindingsForUserProvidedServiceInstance{Context: ctx, Handler: handler}
}

/*ListAllServiceBindingsForUserProvidedServiceInstance swagger:route GET /user_provided_service_instances/{guid}/service_bindings userProvidedServiceInstances listAllServiceBindingsForUserProvidedServiceInstance

List all Service Bindings for the User Provided Service Instance

curl --insecure -i %s/v2/user_provided_service_instances/{guid}/service_bindings -X GET -H 'Authorization: %s'

*/
type ListAllServiceBindingsForUserProvidedServiceInstance struct {
	Context *middleware.Context
	Handler ListAllServiceBindingsForUserProvidedServiceInstanceHandler
}

func (o *ListAllServiceBindingsForUserProvidedServiceInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllServiceBindingsForUserProvidedServiceInstanceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

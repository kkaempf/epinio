// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewCheckRouteExistsParams creates a new CheckRouteExistsParams object
// no default values defined in spec.
func NewCheckRouteExistsParams() CheckRouteExistsParams {

	return CheckRouteExistsParams{}
}

// CheckRouteExistsParams contains all the bound params for the check route exists operation
// typically these are obtained from a http.Request
//
// swagger:parameters checkRouteExists
type CheckRouteExistsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The domain_guid parameter is used as a part of the request URL: '/v2/routes/reserved/domain/:domain_guid/host/:host'
	  Required: true
	  In: path
	*/
	DomainGUID string
	/*The host parameter is used as a part of the request URL: '/v2/routes/reserved/domain/:domain_guid/host/:host'
	  Required: true
	  In: path
	*/
	Host string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCheckRouteExistsParams() beforehand.
func (o *CheckRouteExistsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDomainGUID, rhkDomainGUID, _ := route.Params.GetOK("domain_guid")
	if err := o.bindDomainGUID(rDomainGUID, rhkDomainGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rHost, rhkHost, _ := route.Params.GetOK("host")
	if err := o.bindHost(rHost, rhkHost, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDomainGUID binds and validates parameter DomainGUID from path.
func (o *CheckRouteExistsParams) bindDomainGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.DomainGUID = raw

	return nil
}

// bindHost binds and validates parameter Host from path.
func (o *CheckRouteExistsParams) bindHost(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Host = raw

	return nil
}

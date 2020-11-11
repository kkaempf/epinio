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

// NewRemoveAppFromRouteParams creates a new RemoveAppFromRouteParams object
// no default values defined in spec.
func NewRemoveAppFromRouteParams() RemoveAppFromRouteParams {

	return RemoveAppFromRouteParams{}
}

// RemoveAppFromRouteParams contains all the bound params for the remove app from route operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeAppFromRoute
type RemoveAppFromRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The app_guid parameter is used as a part of the request URL: '/v2/routes/:guid/apps/:app_guid'
	  Required: true
	  In: path
	*/
	AppGUID string
	/*The guid parameter is used as a part of the request URL: '/v2/routes/:guid/apps/:app_guid'
	  Required: true
	  In: path
	*/
	GUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveAppFromRouteParams() beforehand.
func (o *RemoveAppFromRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAppGUID, rhkAppGUID, _ := route.Params.GetOK("app_guid")
	if err := o.bindAppGUID(rAppGUID, rhkAppGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAppGUID binds and validates parameter AppGUID from path.
func (o *RemoveAppFromRouteParams) bindAppGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AppGUID = raw

	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *RemoveAppFromRouteParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

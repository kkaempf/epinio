// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRemoveManagerFromSpaceParams creates a new RemoveManagerFromSpaceParams object
// no default values defined in spec.
func NewRemoveManagerFromSpaceParams() RemoveManagerFromSpaceParams {

	return RemoveManagerFromSpaceParams{}
}

// RemoveManagerFromSpaceParams contains all the bound params for the remove manager from space operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeManagerFromSpace
type RemoveManagerFromSpaceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The guid parameter is used as a part of the request URL: '/v2/spaces/:guid/managers/:manager_guid'
	  Required: true
	  In: path
	*/
	GUID string
	/*The manager_guid parameter is used as a part of the request URL: '/v2/spaces/:guid/managers/:manager_guid'
	  Required: true
	  In: path
	*/
	ManagerGUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveManagerFromSpaceParams() beforehand.
func (o *RemoveManagerFromSpaceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rManagerGUID, rhkManagerGUID, _ := route.Params.GetOK("manager_guid")
	if err := o.bindManagerGUID(rManagerGUID, rhkManagerGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *RemoveManagerFromSpaceParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

// bindManagerGUID binds and validates parameter ManagerGUID from path.
func (o *RemoveManagerFromSpaceParams) bindManagerGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ManagerGUID = raw

	return nil
}

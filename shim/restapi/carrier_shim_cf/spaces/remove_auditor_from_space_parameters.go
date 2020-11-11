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

// NewRemoveAuditorFromSpaceParams creates a new RemoveAuditorFromSpaceParams object
// no default values defined in spec.
func NewRemoveAuditorFromSpaceParams() RemoveAuditorFromSpaceParams {

	return RemoveAuditorFromSpaceParams{}
}

// RemoveAuditorFromSpaceParams contains all the bound params for the remove auditor from space operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeAuditorFromSpace
type RemoveAuditorFromSpaceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The auditor_guid parameter is used as a part of the request URL: '/v2/spaces/:guid/auditors/:auditor_guid'
	  Required: true
	  In: path
	*/
	AuditorGUID string
	/*The guid parameter is used as a part of the request URL: '/v2/spaces/:guid/auditors/:auditor_guid'
	  Required: true
	  In: path
	*/
	GUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveAuditorFromSpaceParams() beforehand.
func (o *RemoveAuditorFromSpaceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAuditorGUID, rhkAuditorGUID, _ := route.Params.GetOK("auditor_guid")
	if err := o.bindAuditorGUID(rAuditorGUID, rhkAuditorGUID, route.Formats); err != nil {
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

// bindAuditorGUID binds and validates parameter AuditorGUID from path.
func (o *RemoveAuditorFromSpaceParams) bindAuditorGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AuditorGUID = raw

	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *RemoveAuditorFromSpaceParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

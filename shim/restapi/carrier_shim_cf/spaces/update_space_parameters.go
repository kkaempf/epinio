// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/suse/carrier/shim/models"
)

// NewUpdateSpaceParams creates a new UpdateSpaceParams object
// no default values defined in spec.
func NewUpdateSpaceParams() UpdateSpaceParams {

	return UpdateSpaceParams{}
}

// UpdateSpaceParams contains all the bound params for the update space operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateSpace
type UpdateSpaceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The guid parameter is used as a part of the request URL: '/v2/spaces/:guid'
	  Required: true
	  In: path
	*/
	GUID string
	/*An object instance that is serialized and sent as the request body.
	  Required: true
	  In: body
	*/
	Value *models.UpdateSpaceRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateSpaceParams() beforehand.
func (o *UpdateSpaceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UpdateSpaceRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("value", "body", ""))
			} else {
				res = append(res, errors.NewParseError("value", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Value = &body
			}
		}
	} else {
		res = append(res, errors.Required("value", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *UpdateSpaceParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

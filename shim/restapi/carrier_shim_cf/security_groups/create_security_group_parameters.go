// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/suse/carrier/shim/models"
)

// NewCreateSecurityGroupParams creates a new CreateSecurityGroupParams object
// no default values defined in spec.
func NewCreateSecurityGroupParams() CreateSecurityGroupParams {

	return CreateSecurityGroupParams{}
}

// CreateSecurityGroupParams contains all the bound params for the create security group operation
// typically these are obtained from a http.Request
//
// swagger:parameters createSecurityGroup
type CreateSecurityGroupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*An object instance that is serialized and sent as the request body.
	  Required: true
	  In: body
	*/
	Value *models.CreateSecurityGroupRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateSecurityGroupParams() beforehand.
func (o *CreateSecurityGroupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CreateSecurityGroupRequest
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

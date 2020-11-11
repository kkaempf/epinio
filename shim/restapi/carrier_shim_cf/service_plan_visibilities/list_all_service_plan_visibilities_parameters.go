// Code generated by go-swagger; DO NOT EDIT.

package service_plan_visibilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewListAllServicePlanVisibilitiesParams creates a new ListAllServicePlanVisibilitiesParams object
// no default values defined in spec.
func NewListAllServicePlanVisibilitiesParams() ListAllServicePlanVisibilitiesParams {

	return ListAllServicePlanVisibilitiesParams{}
}

// ListAllServicePlanVisibilitiesParams contains all the bound params for the list all service plan visibilities operation
// typically these are obtained from a http.Request
//
// swagger:parameters listAllServicePlanVisibilities
type ListAllServicePlanVisibilitiesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListAllServicePlanVisibilitiesParams() beforehand.
func (o *ListAllServicePlanVisibilitiesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

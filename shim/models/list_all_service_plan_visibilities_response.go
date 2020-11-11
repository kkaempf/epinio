// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllServicePlanVisibilitiesResponse list all service plan visibilities response
//
// swagger:model listAllServicePlanVisibilitiesResponse
type ListAllServicePlanVisibilitiesResponse struct {

	// The organization Guid
	OrganizationGUID string `json:"organization_guid,omitempty"`

	// The organization Url
	OrganizationURL string `json:"organization_url,omitempty"`

	// The service Plan Guid
	ServicePlanGUID string `json:"service_plan_guid,omitempty"`

	// The service Plan Url
	ServicePlanURL string `json:"service_plan_url,omitempty"`
}

// Validate validates this list all service plan visibilities response
func (m *ListAllServicePlanVisibilitiesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllServicePlanVisibilitiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllServicePlanVisibilitiesResponse) UnmarshalBinary(b []byte) error {
	var res ListAllServicePlanVisibilitiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

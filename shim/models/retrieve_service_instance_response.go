// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrieveServiceInstanceResponse retrieve service instance response
//
// swagger:model retrieveServiceInstanceResponse
type RetrieveServiceInstanceResponse struct {

	// The credentials
	Credentials GenericObject `json:"credentials,omitempty"`

	// The dashboard Url
	DashboardURL GenericObject `json:"dashboard_url,omitempty"`

	// The gateway Data
	GatewayData GenericObject `json:"gateway_data,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The service Bindings Url
	ServiceBindingsURL string `json:"service_bindings_url,omitempty"`

	// The service Plan Guid
	ServicePlanGUID string `json:"service_plan_guid,omitempty"`

	// The service Plan Url
	ServicePlanURL string `json:"service_plan_url,omitempty"`

	// The space Guid
	SpaceGUID string `json:"space_guid,omitempty"`

	// The space Url
	SpaceURL string `json:"space_url,omitempty"`

	// The type
	Type string `json:"type,omitempty"`
}

// Validate validates this retrieve service instance response
func (m *RetrieveServiceInstanceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveServiceInstanceResponse) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if err := m.Credentials.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentials")
		}
		return err
	}

	return nil
}

func (m *RetrieveServiceInstanceResponse) validateDashboardURL(formats strfmt.Registry) error {

	if swag.IsZero(m.DashboardURL) { // not required
		return nil
	}

	if err := m.DashboardURL.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dashboard_url")
		}
		return err
	}

	return nil
}

func (m *RetrieveServiceInstanceResponse) validateGatewayData(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayData) { // not required
		return nil
	}

	if err := m.GatewayData.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gateway_data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetrieveServiceInstanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveServiceInstanceResponse) UnmarshalBinary(b []byte) error {
	var res RetrieveServiceInstanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

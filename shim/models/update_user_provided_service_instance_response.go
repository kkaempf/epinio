// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateUserProvidedServiceInstanceResponse update user provided service instance response
//
// swagger:model updateUserProvidedServiceInstanceResponse
type UpdateUserProvidedServiceInstanceResponse struct {

	// The credentials
	Credentials GenericObject `json:"credentials,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The service Bindings Url
	ServiceBindingsURL string `json:"service_bindings_url,omitempty"`

	// The space Guid
	SpaceGUID string `json:"space_guid,omitempty"`

	// The space Url
	SpaceURL string `json:"space_url,omitempty"`

	// The syslog Drain Url
	SyslogDrainURL string `json:"syslog_drain_url,omitempty"`

	// The type
	Type string `json:"type,omitempty"`
}

// Validate validates this update user provided service instance response
func (m *UpdateUserProvidedServiceInstanceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserProvidedServiceInstanceResponse) validateCredentials(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *UpdateUserProvidedServiceInstanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserProvidedServiceInstanceResponse) UnmarshalBinary(b []byte) error {
	var res UpdateUserProvidedServiceInstanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

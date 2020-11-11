// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetRouteCreationFeatureFlagResponse get route creation feature flag response
//
// swagger:model getRouteCreationFeatureFlagResponse
type GetRouteCreationFeatureFlagResponse struct {

	// The enabled
	Enabled bool `json:"enabled,omitempty"`

	// The error Message
	ErrorMessage GenericObject `json:"error_message,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The url
	URL string `json:"url,omitempty"`
}

// Validate validates this get route creation feature flag response
func (m *GetRouteCreationFeatureFlagResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetRouteCreationFeatureFlagResponse) validateErrorMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorMessage) { // not required
		return nil
	}

	if err := m.ErrorMessage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("error_message")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetRouteCreationFeatureFlagResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRouteCreationFeatureFlagResponse) UnmarshalBinary(b []byte) error {
	var res GetRouteCreationFeatureFlagResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

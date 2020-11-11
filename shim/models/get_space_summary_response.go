// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSpaceSummaryResponse get space summary response
//
// swagger:model getSpaceSummaryResponse
type GetSpaceSummaryResponse struct {

	// The apps
	Apps []GenericObject `json:"apps"`

	// The guid
	GUID string `json:"guid,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The services
	Services []GenericObject `json:"services"`
}

// Validate validates this get space summary response
func (m *GetSpaceSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSpaceSummaryResponse) validateApps(formats strfmt.Registry) error {

	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	for i := 0; i < len(m.Apps); i++ {

		if err := m.Apps[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GetSpaceSummaryResponse) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {

		if err := m.Services[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSpaceSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSpaceSummaryResponse) UnmarshalBinary(b []byte) error {
	var res GetSpaceSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

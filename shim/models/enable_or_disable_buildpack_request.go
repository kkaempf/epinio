// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnableOrDisableBuildpackRequest enable or disable buildpack request
//
// swagger:model enableOrDisableBuildpackRequest
type EnableOrDisableBuildpackRequest struct {

	// Whether or not the buildpack will be used for staging
	Enabled bool `json:"enabled,omitempty"`

	// The name of the uploaded buildpack file
	Filename GenericObject `json:"filename,omitempty"`

	// Whether or not the buildpack is locked to prevent updates
	Locked GenericObject `json:"locked,omitempty"`

	// The name of the buildpack. To be used by app buildpack field. (only alphanumeric characters)
	Name string `json:"name,omitempty"`

	// The order in which the buildpacks are checked during buildpack auto-detection.
	Position GenericObject `json:"position,omitempty"`
}

// Validate validates this enable or disable buildpack request
func (m *EnableOrDisableBuildpackRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnableOrDisableBuildpackRequest) validateFilename(formats strfmt.Registry) error {

	if swag.IsZero(m.Filename) { // not required
		return nil
	}

	if err := m.Filename.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("filename")
		}
		return err
	}

	return nil
}

func (m *EnableOrDisableBuildpackRequest) validateLocked(formats strfmt.Registry) error {

	if swag.IsZero(m.Locked) { // not required
		return nil
	}

	if err := m.Locked.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locked")
		}
		return err
	}

	return nil
}

func (m *EnableOrDisableBuildpackRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := m.Position.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("position")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnableOrDisableBuildpackRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnableOrDisableBuildpackRequest) UnmarshalBinary(b []byte) error {
	var res EnableOrDisableBuildpackRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

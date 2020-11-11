// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnableOrDisableBuildpackResponse enable or disable buildpack response
//
// swagger:model enableOrDisableBuildpackResponse
type EnableOrDisableBuildpackResponse struct {

	// The enabled
	Enabled bool `json:"enabled,omitempty"`

	// The filename
	Filename string `json:"filename,omitempty"`

	// The locked
	Locked bool `json:"locked,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The position
	Position int64 `json:"position,omitempty"`
}

// Validate validates this enable or disable buildpack response
func (m *EnableOrDisableBuildpackResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnableOrDisableBuildpackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnableOrDisableBuildpackResponse) UnmarshalBinary(b []byte) error {
	var res EnableOrDisableBuildpackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

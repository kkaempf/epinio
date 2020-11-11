// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllBuildpacksResponse list all buildpacks response
//
// swagger:model listAllBuildpacksResponse
type ListAllBuildpacksResponse struct {

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

// Validate validates this list all buildpacks response
func (m *ListAllBuildpacksResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllBuildpacksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllBuildpacksResponse) UnmarshalBinary(b []byte) error {
	var res ListAllBuildpacksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

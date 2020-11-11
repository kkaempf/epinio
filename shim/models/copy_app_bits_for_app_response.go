// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CopyAppBitsForAppResponse copy app bits for app response
//
// swagger:model copyAppBitsForAppResponse
type CopyAppBitsForAppResponse struct {

	// The guid
	GUID string `json:"guid,omitempty"`

	// The status
	Status string `json:"status,omitempty"`
}

// Validate validates this copy app bits for app response
func (m *CopyAppBitsForAppResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CopyAppBitsForAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyAppBitsForAppResponse) UnmarshalBinary(b []byte) error {
	var res CopyAppBitsForAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

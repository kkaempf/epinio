// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllSharedDomainsResponse list all shared domains response
//
// swagger:model listAllSharedDomainsResponse
type ListAllSharedDomainsResponse struct {

	// The name
	Name string `json:"name,omitempty"`
}

// Validate validates this list all shared domains response
func (m *ListAllSharedDomainsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllSharedDomainsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllSharedDomainsResponse) UnmarshalBinary(b []byte) error {
	var res ListAllSharedDomainsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

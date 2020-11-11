// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateSharedDomainRequest create shared domain request
//
// swagger:model createSharedDomainRequest
type CreateSharedDomainRequest struct {

	// The guid of the domain.
	GUID string `json:"guid,omitempty"`

	// The name of the domain.
	Name string `json:"name,omitempty"`
}

// Validate validates this create shared domain request
func (m *CreateSharedDomainRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSharedDomainRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSharedDomainRequest) UnmarshalBinary(b []byte) error {
	var res CreateSharedDomainRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

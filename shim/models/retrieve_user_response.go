// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrieveUserResponse retrieve user response
//
// swagger:model retrieveUserResponse
type RetrieveUserResponse struct {

	// The active
	Active bool `json:"active,omitempty"`

	// The admin
	Admin bool `json:"admin,omitempty"`

	// The audited Organizations Url
	AuditedOrganizationsURL string `json:"audited_organizations_url,omitempty"`

	// The audited Spaces Url
	AuditedSpacesURL string `json:"audited_spaces_url,omitempty"`

	// The billing Managed Organizations Url
	BillingManagedOrganizationsURL string `json:"billing_managed_organizations_url,omitempty"`

	// The default Space Guid
	DefaultSpaceGUID string `json:"default_space_guid,omitempty"`

	// The default Space Url
	DefaultSpaceURL string `json:"default_space_url,omitempty"`

	// The managed Organizations Url
	ManagedOrganizationsURL string `json:"managed_organizations_url,omitempty"`

	// The managed Spaces Url
	ManagedSpacesURL string `json:"managed_spaces_url,omitempty"`

	// The organizations Url
	OrganizationsURL string `json:"organizations_url,omitempty"`

	// The spaces Url
	SpacesURL string `json:"spaces_url,omitempty"`
}

// Validate validates this retrieve user response
func (m *RetrieveUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetrieveUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveUserResponse) UnmarshalBinary(b []byte) error {
	var res RetrieveUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

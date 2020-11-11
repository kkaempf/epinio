// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllRoutesForAppResponse list all routes for app response
//
// swagger:model listAllRoutesForAppResponse
type ListAllRoutesForAppResponse struct {

	// The apps Url
	AppsURL string `json:"apps_url,omitempty"`

	// The domain Guid
	DomainGUID string `json:"domain_guid,omitempty"`

	// The domain Url
	DomainURL string `json:"domain_url,omitempty"`

	// The host
	Host string `json:"host,omitempty"`

	// The space Guid
	SpaceGUID string `json:"space_guid,omitempty"`

	// The space Url
	SpaceURL string `json:"space_url,omitempty"`
}

// Validate validates this list all routes for app response
func (m *ListAllRoutesForAppResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllRoutesForAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllRoutesForAppResponse) UnmarshalBinary(b []byte) error {
	var res ListAllRoutesForAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

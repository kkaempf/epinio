// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetInfoResponse get info response
//
// swagger:model getInfoResponse
type GetInfoResponse struct {

	// The api Version
	APIVersion string `json:"api_version,omitempty"`

	// The authorization Endpoint
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`

	// The build
	Build string `json:"build,omitempty"`

	// The description
	Description string `json:"description,omitempty"`

	// The logging Endpoint
	LoggingEndpoint string `json:"logging_endpoint,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The support
	Support string `json:"support,omitempty"`

	// The token Endpoint
	TokenEndpoint string `json:"token_endpoint,omitempty"`

	// The version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this get info response
func (m *GetInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInfoResponse) UnmarshalBinary(b []byte) error {
	var res GetInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

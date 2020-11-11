// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetSecurityGroupAsDefaultForRunningAppsResponseResource set security group as default for running apps response resource
//
// swagger:model setSecurityGroupAsDefaultForRunningAppsResponseResource
type SetSecurityGroupAsDefaultForRunningAppsResponseResource struct {

	// entity
	Entity *SetSecurityGroupAsDefaultForRunningAppsResponse `json:"entity,omitempty"`

	// metadata
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

// Validate validates this set security group as default for running apps response resource
func (m *SetSecurityGroupAsDefaultForRunningAppsResponseResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetSecurityGroupAsDefaultForRunningAppsResponseResource) validateEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *SetSecurityGroupAsDefaultForRunningAppsResponseResource) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetSecurityGroupAsDefaultForRunningAppsResponseResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetSecurityGroupAsDefaultForRunningAppsResponseResource) UnmarshalBinary(b []byte) error {
	var res SetSecurityGroupAsDefaultForRunningAppsResponseResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

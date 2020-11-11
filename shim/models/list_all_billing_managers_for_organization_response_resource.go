// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllBillingManagersForOrganizationResponseResource list all billing managers for organization response resource
//
// swagger:model listAllBillingManagersForOrganizationResponseResource
type ListAllBillingManagersForOrganizationResponseResource struct {

	// entity
	Entity *ListAllBillingManagersForOrganizationResponse `json:"entity,omitempty"`

	// metadata
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

// Validate validates this list all billing managers for organization response resource
func (m *ListAllBillingManagersForOrganizationResponseResource) Validate(formats strfmt.Registry) error {
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

func (m *ListAllBillingManagersForOrganizationResponseResource) validateEntity(formats strfmt.Registry) error {

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

func (m *ListAllBillingManagersForOrganizationResponseResource) validateMetadata(formats strfmt.Registry) error {

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
func (m *ListAllBillingManagersForOrganizationResponseResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllBillingManagersForOrganizationResponseResource) UnmarshalBinary(b []byte) error {
	var res ListAllBillingManagersForOrganizationResponseResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

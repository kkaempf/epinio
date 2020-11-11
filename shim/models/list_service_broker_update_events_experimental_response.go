// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListServiceBrokerUpdateEventsExperimentalResponse list service broker update events experimental response
//
// swagger:model listServiceBrokerUpdateEventsExperimentalResponse
type ListServiceBrokerUpdateEventsExperimentalResponse struct {

	// The actee
	Actee string `json:"actee,omitempty"`

	// The actee Name
	ActeeName string `json:"actee_name,omitempty"`

	// The actee Type
	ActeeType string `json:"actee_type,omitempty"`

	// The actor
	Actor string `json:"actor,omitempty"`

	// The actor Name
	ActorName string `json:"actor_name,omitempty"`

	// The actor Type
	ActorType string `json:"actor_type,omitempty"`

	// The metadata
	Metadata GenericObject `json:"metadata,omitempty"`

	// The organization Guid
	OrganizationGUID string `json:"organization_guid,omitempty"`

	// The space Guid
	SpaceGUID string `json:"space_guid,omitempty"`

	// The timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// The type
	Type string `json:"type,omitempty"`
}

// Validate validates this list service broker update events experimental response
func (m *ListServiceBrokerUpdateEventsExperimentalResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListServiceBrokerUpdateEventsExperimentalResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := m.Metadata.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metadata")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListServiceBrokerUpdateEventsExperimentalResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListServiceBrokerUpdateEventsExperimentalResponse) UnmarshalBinary(b []byte) error {
	var res ListServiceBrokerUpdateEventsExperimentalResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

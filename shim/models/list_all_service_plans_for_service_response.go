// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllServicePlansForServiceResponse list all service plans for service response
//
// swagger:model listAllServicePlansForServiceResponse
type ListAllServicePlansForServiceResponse struct {

	// The active
	Active bool `json:"active,omitempty"`

	// The description
	Description string `json:"description,omitempty"`

	// The extra
	Extra GenericObject `json:"extra,omitempty"`

	// The free
	Free bool `json:"free,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The public
	Public bool `json:"public,omitempty"`

	// The service Guid
	ServiceGUID string `json:"service_guid,omitempty"`

	// The service Instances Url
	ServiceInstancesURL string `json:"service_instances_url,omitempty"`

	// The service Url
	ServiceURL string `json:"service_url,omitempty"`

	// The unique Id
	UniqueID string `json:"unique_id,omitempty"`
}

// Validate validates this list all service plans for service response
func (m *ListAllServicePlansForServiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtra(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAllServicePlansForServiceResponse) validateExtra(formats strfmt.Registry) error {

	if swag.IsZero(m.Extra) { // not required
		return nil
	}

	if err := m.Extra.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("extra")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListAllServicePlansForServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllServicePlansForServiceResponse) UnmarshalBinary(b []byte) error {
	var res ListAllServicePlansForServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

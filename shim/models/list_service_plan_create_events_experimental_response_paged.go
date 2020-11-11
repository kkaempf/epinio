// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListServicePlanCreateEventsExperimentalResponsePaged list service plan create events experimental response paged
//
// swagger:model listServicePlanCreateEventsExperimentalResponsePaged
type ListServicePlanCreateEventsExperimentalResponsePaged struct {

	// next url
	NextURL string `json:"next_url,omitempty"`

	// prev url
	PrevURL string `json:"prev_url,omitempty"`

	// resources
	Resources []*ListServicePlanCreateEventsExperimentalResponseResource `json:"resources"`

	// total pages
	TotalPages int64 `json:"total_pages,omitempty"`

	// total results
	TotalResults int64 `json:"total_results,omitempty"`
}

// Validate validates this list service plan create events experimental response paged
func (m *ListServicePlanCreateEventsExperimentalResponsePaged) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListServicePlanCreateEventsExperimentalResponsePaged) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListServicePlanCreateEventsExperimentalResponsePaged) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListServicePlanCreateEventsExperimentalResponsePaged) UnmarshalBinary(b []byte) error {
	var res ListServicePlanCreateEventsExperimentalResponsePaged
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

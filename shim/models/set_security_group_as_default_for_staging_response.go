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

// SetSecurityGroupAsDefaultForStagingResponse set security group as default for staging response
//
// swagger:model setSecurityGroupAsDefaultForStagingResponse
type SetSecurityGroupAsDefaultForStagingResponse struct {

	// The name
	Name string `json:"name,omitempty"`

	// The rules
	Rules []GenericObject `json:"rules"`

	// The running Default
	RunningDefault bool `json:"running_default,omitempty"`

	// The staging Default
	StagingDefault bool `json:"staging_default,omitempty"`
}

// Validate validates this set security group as default for staging response
func (m *SetSecurityGroupAsDefaultForStagingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetSecurityGroupAsDefaultForStagingResponse) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {

		if err := m.Rules[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rules" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetSecurityGroupAsDefaultForStagingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetSecurityGroupAsDefaultForStagingResponse) UnmarshalBinary(b []byte) error {
	var res SetSecurityGroupAsDefaultForStagingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

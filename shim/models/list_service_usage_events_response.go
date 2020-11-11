// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListServiceUsageEventsResponse list service usage events response
//
// swagger:model listServiceUsageEventsResponse
type ListServiceUsageEventsResponse struct {

	// The org Guid
	OrgGUID string `json:"org_guid,omitempty"`

	// The service Guid
	ServiceGUID string `json:"service_guid,omitempty"`

	// The service Instance Guid
	ServiceInstanceGUID string `json:"service_instance_guid,omitempty"`

	// The service Instance Name
	ServiceInstanceName string `json:"service_instance_name,omitempty"`

	// The service Instance Type
	ServiceInstanceType string `json:"service_instance_type,omitempty"`

	// The service Label
	ServiceLabel string `json:"service_label,omitempty"`

	// The service Plan Guid
	ServicePlanGUID string `json:"service_plan_guid,omitempty"`

	// The service Plan Name
	ServicePlanName string `json:"service_plan_name,omitempty"`

	// The space Guid
	SpaceGUID string `json:"space_guid,omitempty"`

	// The space Name
	SpaceName string `json:"space_name,omitempty"`

	// The state
	State string `json:"state,omitempty"`
}

// Validate validates this list service usage events response
func (m *ListServiceUsageEventsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListServiceUsageEventsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListServiceUsageEventsResponse) UnmarshalBinary(b []byte) error {
	var res ListServiceUsageEventsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

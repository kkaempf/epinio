// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllSpacesForSpaceQuotaDefinitionResponse list all spaces for space quota definition response
//
// swagger:model listAllSpacesForSpaceQuotaDefinitionResponse
type ListAllSpacesForSpaceQuotaDefinitionResponse struct {

	// The app Events Url
	AppEventsURL string `json:"app_events_url,omitempty"`

	// The apps Url
	AppsURL string `json:"apps_url,omitempty"`

	// The auditors Url
	AuditorsURL string `json:"auditors_url,omitempty"`

	// The developers Url
	DevelopersURL string `json:"developers_url,omitempty"`

	// The domains Url
	DomainsURL string `json:"domains_url,omitempty"`

	// The events Url
	EventsURL string `json:"events_url,omitempty"`

	// The managers Url
	ManagersURL string `json:"managers_url,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The organization Guid
	OrganizationGUID string `json:"organization_guid,omitempty"`

	// The organization Url
	OrganizationURL string `json:"organization_url,omitempty"`

	// The routes Url
	RoutesURL string `json:"routes_url,omitempty"`

	// The security Groups Url
	SecurityGroupsURL string `json:"security_groups_url,omitempty"`

	// The service Instances Url
	ServiceInstancesURL string `json:"service_instances_url,omitempty"`

	// The space Quota Definition Guid
	SpaceQuotaDefinitionGUID string `json:"space_quota_definition_guid,omitempty"`

	// The space Quota Definition Url
	SpaceQuotaDefinitionURL string `json:"space_quota_definition_url,omitempty"`
}

// Validate validates this list all spaces for space quota definition response
func (m *ListAllSpacesForSpaceQuotaDefinitionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllSpacesForSpaceQuotaDefinitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllSpacesForSpaceQuotaDefinitionResponse) UnmarshalBinary(b []byte) error {
	var res ListAllSpacesForSpaceQuotaDefinitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

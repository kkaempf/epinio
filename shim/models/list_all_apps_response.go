// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllAppsResponse list all apps response
//
// swagger:model listAllAppsResponse
type ListAllAppsResponse struct {

	// The buildpack
	Buildpack GenericObject `json:"buildpack,omitempty"`

	// The command
	Command GenericObject `json:"command,omitempty"`

	// The console
	Console bool `json:"console,omitempty"`

	// The debug
	Debug GenericObject `json:"debug,omitempty"`

	// The detected Buildpack
	DetectedBuildpack GenericObject `json:"detected_buildpack,omitempty"`

	// The detected Start Command
	DetectedStartCommand string `json:"detected_start_command,omitempty"`

	// The disk Quota
	DiskQuota int64 `json:"disk_quota,omitempty"`

	// The docker Image
	DockerImage GenericObject `json:"docker_image,omitempty"`

	// The environment Json
	EnvironmentJSON GenericObject `json:"environment_json,omitempty"`

	// The events Url
	EventsURL string `json:"events_url,omitempty"`

	// The health Check Timeout
	HealthCheckTimeout GenericObject `json:"health_check_timeout,omitempty"`

	// The health Check Type
	HealthCheckType string `json:"health_check_type,omitempty"`

	// The instances
	Instances int64 `json:"instances,omitempty"`

	// The memory
	Memory int64 `json:"memory,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The package State
	PackageState string `json:"package_state,omitempty"`

	// The package Updated At
	PackageUpdatedAt string `json:"package_updated_at,omitempty"`

	// The production
	Production bool `json:"production,omitempty"`

	// The routes Url
	RoutesURL string `json:"routes_url,omitempty"`

	// The service Bindings Url
	ServiceBindingsURL string `json:"service_bindings_url,omitempty"`

	// The space Guid
	SpaceGUID string `json:"space_guid,omitempty"`

	// The space Url
	SpaceURL string `json:"space_url,omitempty"`

	// The stack Guid
	StackGUID string `json:"stack_guid,omitempty"`

	// The stack Url
	StackURL string `json:"stack_url,omitempty"`

	// The staging Failed Reason
	StagingFailedReason GenericObject `json:"staging_failed_reason,omitempty"`

	// The staging Task Id
	StagingTaskID GenericObject `json:"staging_task_id,omitempty"`

	// The state
	State string `json:"state,omitempty"`

	// The version
	Version string `json:"version,omitempty"`
}

// Validate validates this list all apps response
func (m *ListAllAppsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildpack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetectedBuildpack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthCheckTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStagingFailedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStagingTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAllAppsResponse) validateBuildpack(formats strfmt.Registry) error {

	if swag.IsZero(m.Buildpack) { // not required
		return nil
	}

	if err := m.Buildpack.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("buildpack")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	if err := m.Command.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("command")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateDebug(formats strfmt.Registry) error {

	if swag.IsZero(m.Debug) { // not required
		return nil
	}

	if err := m.Debug.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debug")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateDetectedBuildpack(formats strfmt.Registry) error {

	if swag.IsZero(m.DetectedBuildpack) { // not required
		return nil
	}

	if err := m.DetectedBuildpack.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("detected_buildpack")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateDockerImage(formats strfmt.Registry) error {

	if swag.IsZero(m.DockerImage) { // not required
		return nil
	}

	if err := m.DockerImage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("docker_image")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateEnvironmentJSON(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentJSON) { // not required
		return nil
	}

	if err := m.EnvironmentJSON.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("environment_json")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateHealthCheckTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthCheckTimeout) { // not required
		return nil
	}

	if err := m.HealthCheckTimeout.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("health_check_timeout")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateStagingFailedReason(formats strfmt.Registry) error {

	if swag.IsZero(m.StagingFailedReason) { // not required
		return nil
	}

	if err := m.StagingFailedReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staging_failed_reason")
		}
		return err
	}

	return nil
}

func (m *ListAllAppsResponse) validateStagingTaskID(formats strfmt.Registry) error {

	if swag.IsZero(m.StagingTaskID) { // not required
		return nil
	}

	if err := m.StagingTaskID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staging_task_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListAllAppsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllAppsResponse) UnmarshalBinary(b []byte) error {
	var res ListAllAppsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package environment_variable_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// UpdateContentsOfRunningEnvironmentVariableGroupURL generates an URL for the update contents of running environment variable group operation
type UpdateContentsOfRunningEnvironmentVariableGroupURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) WithBasePath(bp string) *UpdateContentsOfRunningEnvironmentVariableGroupURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/config/environment_variable_groups/running"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdateContentsOfRunningEnvironmentVariableGroupURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdateContentsOfRunningEnvironmentVariableGroupURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UpdateContentsOfRunningEnvironmentVariableGroupURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

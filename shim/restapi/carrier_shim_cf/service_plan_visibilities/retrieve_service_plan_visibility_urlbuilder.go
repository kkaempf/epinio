// Code generated by go-swagger; DO NOT EDIT.

package service_plan_visibilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RetrieveServicePlanVisibilityURL generates an URL for the retrieve service plan visibility operation
type RetrieveServicePlanVisibilityURL struct {
	GUID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveServicePlanVisibilityURL) WithBasePath(bp string) *RetrieveServicePlanVisibilityURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveServicePlanVisibilityURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RetrieveServicePlanVisibilityURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/service_plan_visibilities/{guid}"

	guid := o.GUID
	if guid != "" {
		_path = strings.Replace(_path, "{guid}", guid, -1)
	} else {
		return nil, errors.New("guid is required on RetrieveServicePlanVisibilityURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RetrieveServicePlanVisibilityURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RetrieveServicePlanVisibilityURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RetrieveServicePlanVisibilityURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RetrieveServicePlanVisibilityURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RetrieveServicePlanVisibilityURL")
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
func (o *RetrieveServicePlanVisibilityURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

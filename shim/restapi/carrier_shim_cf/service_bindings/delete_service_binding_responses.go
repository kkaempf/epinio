// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteServiceBindingNoContentCode is the HTTP code returned for type DeleteServiceBindingNoContent
const DeleteServiceBindingNoContentCode int = 204

/*DeleteServiceBindingNoContent successful response

swagger:response deleteServiceBindingNoContent
*/
type DeleteServiceBindingNoContent struct {
}

// NewDeleteServiceBindingNoContent creates DeleteServiceBindingNoContent with default headers values
func NewDeleteServiceBindingNoContent() *DeleteServiceBindingNoContent {

	return &DeleteServiceBindingNoContent{}
}

// WriteResponse to the client
func (o *DeleteServiceBindingNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

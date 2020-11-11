// Code generated by go-swagger; DO NOT EDIT.

package service_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteServicePlansNoContentCode is the HTTP code returned for type DeleteServicePlansNoContent
const DeleteServicePlansNoContentCode int = 204

/*DeleteServicePlansNoContent successful response

swagger:response deleteServicePlansNoContent
*/
type DeleteServicePlansNoContent struct {
}

// NewDeleteServicePlansNoContent creates DeleteServicePlansNoContent with default headers values
func NewDeleteServicePlansNoContent() *DeleteServicePlansNoContent {

	return &DeleteServicePlansNoContent{}
}

// WriteResponse to the client
func (o *DeleteServicePlansNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

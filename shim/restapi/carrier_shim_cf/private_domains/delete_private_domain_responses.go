// Code generated by go-swagger; DO NOT EDIT.

package private_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeletePrivateDomainNoContentCode is the HTTP code returned for type DeletePrivateDomainNoContent
const DeletePrivateDomainNoContentCode int = 204

/*DeletePrivateDomainNoContent successful response

swagger:response deletePrivateDomainNoContent
*/
type DeletePrivateDomainNoContent struct {
}

// NewDeletePrivateDomainNoContent creates DeletePrivateDomainNoContent with default headers values
func NewDeletePrivateDomainNoContent() *DeletePrivateDomainNoContent {

	return &DeletePrivateDomainNoContent{}
}

// WriteResponse to the client
func (o *DeletePrivateDomainNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

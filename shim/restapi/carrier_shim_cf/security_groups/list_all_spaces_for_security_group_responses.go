// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllSpacesForSecurityGroupOKCode is the HTTP code returned for type ListAllSpacesForSecurityGroupOK
const ListAllSpacesForSecurityGroupOKCode int = 200

/*ListAllSpacesForSecurityGroupOK successful response

swagger:response listAllSpacesForSecurityGroupOK
*/
type ListAllSpacesForSecurityGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllSpacesForSecurityGroupResponsePaged `json:"body,omitempty"`
}

// NewListAllSpacesForSecurityGroupOK creates ListAllSpacesForSecurityGroupOK with default headers values
func NewListAllSpacesForSecurityGroupOK() *ListAllSpacesForSecurityGroupOK {

	return &ListAllSpacesForSecurityGroupOK{}
}

// WithPayload adds the payload to the list all spaces for security group o k response
func (o *ListAllSpacesForSecurityGroupOK) WithPayload(payload *models.ListAllSpacesForSecurityGroupResponsePaged) *ListAllSpacesForSecurityGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all spaces for security group o k response
func (o *ListAllSpacesForSecurityGroupOK) SetPayload(payload *models.ListAllSpacesForSecurityGroupResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllSpacesForSecurityGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

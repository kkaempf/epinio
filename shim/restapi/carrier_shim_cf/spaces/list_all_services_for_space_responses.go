// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllServicesForSpaceOKCode is the HTTP code returned for type ListAllServicesForSpaceOK
const ListAllServicesForSpaceOKCode int = 200

/*ListAllServicesForSpaceOK successful response

swagger:response listAllServicesForSpaceOK
*/
type ListAllServicesForSpaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllServicesForSpaceResponsePaged `json:"body,omitempty"`
}

// NewListAllServicesForSpaceOK creates ListAllServicesForSpaceOK with default headers values
func NewListAllServicesForSpaceOK() *ListAllServicesForSpaceOK {

	return &ListAllServicesForSpaceOK{}
}

// WithPayload adds the payload to the list all services for space o k response
func (o *ListAllServicesForSpaceOK) WithPayload(payload *models.ListAllServicesForSpaceResponsePaged) *ListAllServicesForSpaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all services for space o k response
func (o *ListAllServicesForSpaceOK) SetPayload(payload *models.ListAllServicesForSpaceResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllServicesForSpaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllDevelopersForSpaceOKCode is the HTTP code returned for type ListAllDevelopersForSpaceOK
const ListAllDevelopersForSpaceOKCode int = 200

/*ListAllDevelopersForSpaceOK successful response

swagger:response listAllDevelopersForSpaceOK
*/
type ListAllDevelopersForSpaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllDevelopersForSpaceResponsePaged `json:"body,omitempty"`
}

// NewListAllDevelopersForSpaceOK creates ListAllDevelopersForSpaceOK with default headers values
func NewListAllDevelopersForSpaceOK() *ListAllDevelopersForSpaceOK {

	return &ListAllDevelopersForSpaceOK{}
}

// WithPayload adds the payload to the list all developers for space o k response
func (o *ListAllDevelopersForSpaceOK) WithPayload(payload *models.ListAllDevelopersForSpaceResponsePaged) *ListAllDevelopersForSpaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all developers for space o k response
func (o *ListAllDevelopersForSpaceOK) SetPayload(payload *models.ListAllDevelopersForSpaceResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllDevelopersForSpaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveStackOKCode is the HTTP code returned for type RetrieveStackOK
const RetrieveStackOKCode int = 200

/*RetrieveStackOK successful response

swagger:response retrieveStackOK
*/
type RetrieveStackOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveStackResponse `json:"body,omitempty"`
}

// NewRetrieveStackOK creates RetrieveStackOK with default headers values
func NewRetrieveStackOK() *RetrieveStackOK {

	return &RetrieveStackOK{}
}

// WithPayload adds the payload to the retrieve stack o k response
func (o *RetrieveStackOK) WithPayload(payload *models.RetrieveStackResponse) *RetrieveStackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve stack o k response
func (o *RetrieveStackOK) SetPayload(payload *models.RetrieveStackResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveStackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

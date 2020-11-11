// Code generated by go-swagger; DO NOT EDIT.

package user_provided_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveUserProvidedServiceInstanceOKCode is the HTTP code returned for type RetrieveUserProvidedServiceInstanceOK
const RetrieveUserProvidedServiceInstanceOKCode int = 200

/*RetrieveUserProvidedServiceInstanceOK successful response

swagger:response retrieveUserProvidedServiceInstanceOK
*/
type RetrieveUserProvidedServiceInstanceOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveUserProvidedServiceInstanceResponse `json:"body,omitempty"`
}

// NewRetrieveUserProvidedServiceInstanceOK creates RetrieveUserProvidedServiceInstanceOK with default headers values
func NewRetrieveUserProvidedServiceInstanceOK() *RetrieveUserProvidedServiceInstanceOK {

	return &RetrieveUserProvidedServiceInstanceOK{}
}

// WithPayload adds the payload to the retrieve user provided service instance o k response
func (o *RetrieveUserProvidedServiceInstanceOK) WithPayload(payload *models.RetrieveUserProvidedServiceInstanceResponse) *RetrieveUserProvidedServiceInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve user provided service instance o k response
func (o *RetrieveUserProvidedServiceInstanceOK) SetPayload(payload *models.RetrieveUserProvidedServiceInstanceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveUserProvidedServiceInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package service_plan_visibilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllServicePlanVisibilitiesOKCode is the HTTP code returned for type ListAllServicePlanVisibilitiesOK
const ListAllServicePlanVisibilitiesOKCode int = 200

/*ListAllServicePlanVisibilitiesOK successful response

swagger:response listAllServicePlanVisibilitiesOK
*/
type ListAllServicePlanVisibilitiesOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllServicePlanVisibilitiesResponsePaged `json:"body,omitempty"`
}

// NewListAllServicePlanVisibilitiesOK creates ListAllServicePlanVisibilitiesOK with default headers values
func NewListAllServicePlanVisibilitiesOK() *ListAllServicePlanVisibilitiesOK {

	return &ListAllServicePlanVisibilitiesOK{}
}

// WithPayload adds the payload to the list all service plan visibilities o k response
func (o *ListAllServicePlanVisibilitiesOK) WithPayload(payload *models.ListAllServicePlanVisibilitiesResponsePaged) *ListAllServicePlanVisibilitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all service plan visibilities o k response
func (o *ListAllServicePlanVisibilitiesOK) SetPayload(payload *models.ListAllServicePlanVisibilitiesResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllServicePlanVisibilitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package service_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// UpdateServicePlanDeprecatedCreatedCode is the HTTP code returned for type UpdateServicePlanDeprecatedCreated
const UpdateServicePlanDeprecatedCreatedCode int = 201

/*UpdateServicePlanDeprecatedCreated successful response

swagger:response updateServicePlanDeprecatedCreated
*/
type UpdateServicePlanDeprecatedCreated struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateServicePlanDeprecatedResponse `json:"body,omitempty"`
}

// NewUpdateServicePlanDeprecatedCreated creates UpdateServicePlanDeprecatedCreated with default headers values
func NewUpdateServicePlanDeprecatedCreated() *UpdateServicePlanDeprecatedCreated {

	return &UpdateServicePlanDeprecatedCreated{}
}

// WithPayload adds the payload to the update service plan deprecated created response
func (o *UpdateServicePlanDeprecatedCreated) WithPayload(payload *models.UpdateServicePlanDeprecatedResponse) *UpdateServicePlanDeprecatedCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service plan deprecated created response
func (o *UpdateServicePlanDeprecatedCreated) SetPayload(payload *models.UpdateServicePlanDeprecatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServicePlanDeprecatedCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

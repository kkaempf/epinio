// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveBillingManagerFromOrganizationCreatedCode is the HTTP code returned for type RemoveBillingManagerFromOrganizationCreated
const RemoveBillingManagerFromOrganizationCreatedCode int = 201

/*RemoveBillingManagerFromOrganizationCreated successful response

swagger:response removeBillingManagerFromOrganizationCreated
*/
type RemoveBillingManagerFromOrganizationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveBillingManagerFromOrganizationResponse `json:"body,omitempty"`
}

// NewRemoveBillingManagerFromOrganizationCreated creates RemoveBillingManagerFromOrganizationCreated with default headers values
func NewRemoveBillingManagerFromOrganizationCreated() *RemoveBillingManagerFromOrganizationCreated {

	return &RemoveBillingManagerFromOrganizationCreated{}
}

// WithPayload adds the payload to the remove billing manager from organization created response
func (o *RemoveBillingManagerFromOrganizationCreated) WithPayload(payload *models.RemoveBillingManagerFromOrganizationResponse) *RemoveBillingManagerFromOrganizationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove billing manager from organization created response
func (o *RemoveBillingManagerFromOrganizationCreated) SetPayload(payload *models.RemoveBillingManagerFromOrganizationResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveBillingManagerFromOrganizationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

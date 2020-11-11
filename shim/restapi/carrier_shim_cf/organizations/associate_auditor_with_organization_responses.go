// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// AssociateAuditorWithOrganizationCreatedCode is the HTTP code returned for type AssociateAuditorWithOrganizationCreated
const AssociateAuditorWithOrganizationCreatedCode int = 201

/*AssociateAuditorWithOrganizationCreated successful response

swagger:response associateAuditorWithOrganizationCreated
*/
type AssociateAuditorWithOrganizationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AssociateAuditorWithOrganizationResponse `json:"body,omitempty"`
}

// NewAssociateAuditorWithOrganizationCreated creates AssociateAuditorWithOrganizationCreated with default headers values
func NewAssociateAuditorWithOrganizationCreated() *AssociateAuditorWithOrganizationCreated {

	return &AssociateAuditorWithOrganizationCreated{}
}

// WithPayload adds the payload to the associate auditor with organization created response
func (o *AssociateAuditorWithOrganizationCreated) WithPayload(payload *models.AssociateAuditorWithOrganizationResponse) *AssociateAuditorWithOrganizationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate auditor with organization created response
func (o *AssociateAuditorWithOrganizationCreated) SetPayload(payload *models.AssociateAuditorWithOrganizationResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateAuditorWithOrganizationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

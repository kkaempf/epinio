// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// AssociateAuditedOrganizationWithUserCreatedCode is the HTTP code returned for type AssociateAuditedOrganizationWithUserCreated
const AssociateAuditedOrganizationWithUserCreatedCode int = 201

/*AssociateAuditedOrganizationWithUserCreated successful response

swagger:response associateAuditedOrganizationWithUserCreated
*/
type AssociateAuditedOrganizationWithUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AssociateAuditedOrganizationWithUserResponse `json:"body,omitempty"`
}

// NewAssociateAuditedOrganizationWithUserCreated creates AssociateAuditedOrganizationWithUserCreated with default headers values
func NewAssociateAuditedOrganizationWithUserCreated() *AssociateAuditedOrganizationWithUserCreated {

	return &AssociateAuditedOrganizationWithUserCreated{}
}

// WithPayload adds the payload to the associate audited organization with user created response
func (o *AssociateAuditedOrganizationWithUserCreated) WithPayload(payload *models.AssociateAuditedOrganizationWithUserResponse) *AssociateAuditedOrganizationWithUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate audited organization with user created response
func (o *AssociateAuditedOrganizationWithUserCreated) SetPayload(payload *models.AssociateAuditedOrganizationWithUserResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateAuditedOrganizationWithUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

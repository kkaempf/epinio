// Code generated by go-swagger; DO NOT EDIT.

package organization_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveOrganizationQuotaDefinitionOKCode is the HTTP code returned for type RetrieveOrganizationQuotaDefinitionOK
const RetrieveOrganizationQuotaDefinitionOKCode int = 200

/*RetrieveOrganizationQuotaDefinitionOK successful response

swagger:response retrieveOrganizationQuotaDefinitionOK
*/
type RetrieveOrganizationQuotaDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveOrganizationQuotaDefinitionResponse `json:"body,omitempty"`
}

// NewRetrieveOrganizationQuotaDefinitionOK creates RetrieveOrganizationQuotaDefinitionOK with default headers values
func NewRetrieveOrganizationQuotaDefinitionOK() *RetrieveOrganizationQuotaDefinitionOK {

	return &RetrieveOrganizationQuotaDefinitionOK{}
}

// WithPayload adds the payload to the retrieve organization quota definition o k response
func (o *RetrieveOrganizationQuotaDefinitionOK) WithPayload(payload *models.RetrieveOrganizationQuotaDefinitionResponse) *RetrieveOrganizationQuotaDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve organization quota definition o k response
func (o *RetrieveOrganizationQuotaDefinitionOK) SetPayload(payload *models.RetrieveOrganizationQuotaDefinitionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveOrganizationQuotaDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

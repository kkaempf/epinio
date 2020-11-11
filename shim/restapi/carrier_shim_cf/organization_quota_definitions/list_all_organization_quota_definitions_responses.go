// Code generated by go-swagger; DO NOT EDIT.

package organization_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllOrganizationQuotaDefinitionsOKCode is the HTTP code returned for type ListAllOrganizationQuotaDefinitionsOK
const ListAllOrganizationQuotaDefinitionsOKCode int = 200

/*ListAllOrganizationQuotaDefinitionsOK successful response

swagger:response listAllOrganizationQuotaDefinitionsOK
*/
type ListAllOrganizationQuotaDefinitionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllOrganizationQuotaDefinitionsResponsePaged `json:"body,omitempty"`
}

// NewListAllOrganizationQuotaDefinitionsOK creates ListAllOrganizationQuotaDefinitionsOK with default headers values
func NewListAllOrganizationQuotaDefinitionsOK() *ListAllOrganizationQuotaDefinitionsOK {

	return &ListAllOrganizationQuotaDefinitionsOK{}
}

// WithPayload adds the payload to the list all organization quota definitions o k response
func (o *ListAllOrganizationQuotaDefinitionsOK) WithPayload(payload *models.ListAllOrganizationQuotaDefinitionsResponsePaged) *ListAllOrganizationQuotaDefinitionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all organization quota definitions o k response
func (o *ListAllOrganizationQuotaDefinitionsOK) SetPayload(payload *models.ListAllOrganizationQuotaDefinitionsResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllOrganizationQuotaDefinitionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

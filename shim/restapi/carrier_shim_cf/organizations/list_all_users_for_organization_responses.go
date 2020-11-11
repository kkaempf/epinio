// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllUsersForOrganizationOKCode is the HTTP code returned for type ListAllUsersForOrganizationOK
const ListAllUsersForOrganizationOKCode int = 200

/*ListAllUsersForOrganizationOK successful response

swagger:response listAllUsersForOrganizationOK
*/
type ListAllUsersForOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllUsersForOrganizationResponsePaged `json:"body,omitempty"`
}

// NewListAllUsersForOrganizationOK creates ListAllUsersForOrganizationOK with default headers values
func NewListAllUsersForOrganizationOK() *ListAllUsersForOrganizationOK {

	return &ListAllUsersForOrganizationOK{}
}

// WithPayload adds the payload to the list all users for organization o k response
func (o *ListAllUsersForOrganizationOK) WithPayload(payload *models.ListAllUsersForOrganizationResponsePaged) *ListAllUsersForOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all users for organization o k response
func (o *ListAllUsersForOrganizationOK) SetPayload(payload *models.ListAllUsersForOrganizationResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllUsersForOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

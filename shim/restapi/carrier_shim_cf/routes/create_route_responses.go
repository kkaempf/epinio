// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// CreateRouteCreatedCode is the HTTP code returned for type CreateRouteCreated
const CreateRouteCreatedCode int = 201

/*CreateRouteCreated successful response

swagger:response createRouteCreated
*/
type CreateRouteCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreateRouteResponse `json:"body,omitempty"`
}

// NewCreateRouteCreated creates CreateRouteCreated with default headers values
func NewCreateRouteCreated() *CreateRouteCreated {

	return &CreateRouteCreated{}
}

// WithPayload adds the payload to the create route created response
func (o *CreateRouteCreated) WithPayload(payload *models.CreateRouteResponse) *CreateRouteCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create route created response
func (o *CreateRouteCreated) SetPayload(payload *models.CreateRouteResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRouteCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

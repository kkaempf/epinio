// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// GetInstanceInformationForStartedAppOKCode is the HTTP code returned for type GetInstanceInformationForStartedAppOK
const GetInstanceInformationForStartedAppOKCode int = 200

/*GetInstanceInformationForStartedAppOK successful response

swagger:response getInstanceInformationForStartedAppOK
*/
type GetInstanceInformationForStartedAppOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetInstanceInformationForStartedAppResponse `json:"body,omitempty"`
}

// NewGetInstanceInformationForStartedAppOK creates GetInstanceInformationForStartedAppOK with default headers values
func NewGetInstanceInformationForStartedAppOK() *GetInstanceInformationForStartedAppOK {

	return &GetInstanceInformationForStartedAppOK{}
}

// WithPayload adds the payload to the get instance information for started app o k response
func (o *GetInstanceInformationForStartedAppOK) WithPayload(payload *models.GetInstanceInformationForStartedAppResponse) *GetInstanceInformationForStartedAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance information for started app o k response
func (o *GetInstanceInformationForStartedAppOK) SetPayload(payload *models.GetInstanceInformationForStartedAppResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceInformationForStartedAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SignUpOKCode is the HTTP code returned for type SignUpOK
const SignUpOKCode int = 200

/*
SignUpOK OK

swagger:response signUpOK
*/
type SignUpOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSignUpOK creates SignUpOK with default headers values
func NewSignUpOK() *SignUpOK {

	return &SignUpOK{}
}

// WithPayload adds the payload to the sign up o k response
func (o *SignUpOK) WithPayload(payload string) *SignUpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign up o k response
func (o *SignUpOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignUpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

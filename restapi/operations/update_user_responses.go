// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateUserOKCode is the HTTP code returned for type UpdateUserOK
const UpdateUserOKCode int = 200

/*
UpdateUserOK OK

swagger:response updateUserOK
*/
type UpdateUserOK struct {
}

// NewUpdateUserOK creates UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {

	return &UpdateUserOK{}
}

// WriteResponse to the client
func (o *UpdateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
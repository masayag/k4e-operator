// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostControlMessageForDeviceOKCode is the HTTP code returned for type PostControlMessageForDeviceOK
const PostControlMessageForDeviceOKCode int = 200

/*PostControlMessageForDeviceOK Success

swagger:response postControlMessageForDeviceOK
*/
type PostControlMessageForDeviceOK struct {
}

// NewPostControlMessageForDeviceOK creates PostControlMessageForDeviceOK with default headers values
func NewPostControlMessageForDeviceOK() *PostControlMessageForDeviceOK {

	return &PostControlMessageForDeviceOK{}
}

// WriteResponse to the client
func (o *PostControlMessageForDeviceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostControlMessageForDeviceUnauthorizedCode is the HTTP code returned for type PostControlMessageForDeviceUnauthorized
const PostControlMessageForDeviceUnauthorizedCode int = 401

/*PostControlMessageForDeviceUnauthorized Unauthorized

swagger:response postControlMessageForDeviceUnauthorized
*/
type PostControlMessageForDeviceUnauthorized struct {
}

// NewPostControlMessageForDeviceUnauthorized creates PostControlMessageForDeviceUnauthorized with default headers values
func NewPostControlMessageForDeviceUnauthorized() *PostControlMessageForDeviceUnauthorized {

	return &PostControlMessageForDeviceUnauthorized{}
}

// WriteResponse to the client
func (o *PostControlMessageForDeviceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostControlMessageForDeviceForbiddenCode is the HTTP code returned for type PostControlMessageForDeviceForbidden
const PostControlMessageForDeviceForbiddenCode int = 403

/*PostControlMessageForDeviceForbidden Forbidden

swagger:response postControlMessageForDeviceForbidden
*/
type PostControlMessageForDeviceForbidden struct {
}

// NewPostControlMessageForDeviceForbidden creates PostControlMessageForDeviceForbidden with default headers values
func NewPostControlMessageForDeviceForbidden() *PostControlMessageForDeviceForbidden {

	return &PostControlMessageForDeviceForbidden{}
}

// WriteResponse to the client
func (o *PostControlMessageForDeviceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostControlMessageForDeviceNotFoundCode is the HTTP code returned for type PostControlMessageForDeviceNotFound
const PostControlMessageForDeviceNotFoundCode int = 404

/*PostControlMessageForDeviceNotFound Error

swagger:response postControlMessageForDeviceNotFound
*/
type PostControlMessageForDeviceNotFound struct {
}

// NewPostControlMessageForDeviceNotFound creates PostControlMessageForDeviceNotFound with default headers values
func NewPostControlMessageForDeviceNotFound() *PostControlMessageForDeviceNotFound {

	return &PostControlMessageForDeviceNotFound{}
}

// WriteResponse to the client
func (o *PostControlMessageForDeviceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostControlMessageForDeviceInternalServerErrorCode is the HTTP code returned for type PostControlMessageForDeviceInternalServerError
const PostControlMessageForDeviceInternalServerErrorCode int = 500

/*PostControlMessageForDeviceInternalServerError Error

swagger:response postControlMessageForDeviceInternalServerError
*/
type PostControlMessageForDeviceInternalServerError struct {
}

// NewPostControlMessageForDeviceInternalServerError creates PostControlMessageForDeviceInternalServerError with default headers values
func NewPostControlMessageForDeviceInternalServerError() *PostControlMessageForDeviceInternalServerError {

	return &PostControlMessageForDeviceInternalServerError{}
}

// WriteResponse to the client
func (o *PostControlMessageForDeviceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

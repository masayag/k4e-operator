// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakub-dzon/k4e-operator/models"
)

// GetControlMessageForDeviceReader is a Reader for the GetControlMessageForDevice structure.
type GetControlMessageForDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControlMessageForDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetControlMessageForDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetControlMessageForDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetControlMessageForDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetControlMessageForDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetControlMessageForDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetControlMessageForDeviceOK creates a GetControlMessageForDeviceOK with default headers values
func NewGetControlMessageForDeviceOK() *GetControlMessageForDeviceOK {
	return &GetControlMessageForDeviceOK{}
}

/*GetControlMessageForDeviceOK handles this case with default header values.

Success
*/
type GetControlMessageForDeviceOK struct {
	Payload *models.Message
}

func (o *GetControlMessageForDeviceOK) Error() string {
	return fmt.Sprintf("[GET /control/{device_id}/in][%d] getControlMessageForDeviceOK  %+v", 200, o.Payload)
}

func (o *GetControlMessageForDeviceOK) GetPayload() *models.Message {
	return o.Payload
}

func (o *GetControlMessageForDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControlMessageForDeviceUnauthorized creates a GetControlMessageForDeviceUnauthorized with default headers values
func NewGetControlMessageForDeviceUnauthorized() *GetControlMessageForDeviceUnauthorized {
	return &GetControlMessageForDeviceUnauthorized{}
}

/*GetControlMessageForDeviceUnauthorized handles this case with default header values.

Unauthorized
*/
type GetControlMessageForDeviceUnauthorized struct {
}

func (o *GetControlMessageForDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /control/{device_id}/in][%d] getControlMessageForDeviceUnauthorized ", 401)
}

func (o *GetControlMessageForDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControlMessageForDeviceForbidden creates a GetControlMessageForDeviceForbidden with default headers values
func NewGetControlMessageForDeviceForbidden() *GetControlMessageForDeviceForbidden {
	return &GetControlMessageForDeviceForbidden{}
}

/*GetControlMessageForDeviceForbidden handles this case with default header values.

Forbidden
*/
type GetControlMessageForDeviceForbidden struct {
}

func (o *GetControlMessageForDeviceForbidden) Error() string {
	return fmt.Sprintf("[GET /control/{device_id}/in][%d] getControlMessageForDeviceForbidden ", 403)
}

func (o *GetControlMessageForDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControlMessageForDeviceNotFound creates a GetControlMessageForDeviceNotFound with default headers values
func NewGetControlMessageForDeviceNotFound() *GetControlMessageForDeviceNotFound {
	return &GetControlMessageForDeviceNotFound{}
}

/*GetControlMessageForDeviceNotFound handles this case with default header values.

Error
*/
type GetControlMessageForDeviceNotFound struct {
}

func (o *GetControlMessageForDeviceNotFound) Error() string {
	return fmt.Sprintf("[GET /control/{device_id}/in][%d] getControlMessageForDeviceNotFound ", 404)
}

func (o *GetControlMessageForDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControlMessageForDeviceInternalServerError creates a GetControlMessageForDeviceInternalServerError with default headers values
func NewGetControlMessageForDeviceInternalServerError() *GetControlMessageForDeviceInternalServerError {
	return &GetControlMessageForDeviceInternalServerError{}
}

/*GetControlMessageForDeviceInternalServerError handles this case with default header values.

Error
*/
type GetControlMessageForDeviceInternalServerError struct {
}

func (o *GetControlMessageForDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /control/{device_id}/in][%d] getControlMessageForDeviceInternalServerError ", 500)
}

func (o *GetControlMessageForDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fnproject/fn_go/modelsv2"
)

// UpdateAppReader is a Reader for the UpdateApp structure.
type UpdateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAppOK creates a UpdateAppOK with default headers values
func NewUpdateAppOK() *UpdateAppOK {
	return &UpdateAppOK{}
}

/*UpdateAppOK handles this case with default header values.

Application details and stats.
*/
type UpdateAppOK struct {
	Payload *modelsv2.App
}

func (o *UpdateAppOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appID}][%d] updateAppOK  %+v", 200, o.Payload)
}

func (o *UpdateAppOK) GetPayload() *modelsv2.App {
	return o.Payload
}

func (o *UpdateAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppBadRequest creates a UpdateAppBadRequest with default headers values
func NewUpdateAppBadRequest() *UpdateAppBadRequest {
	return &UpdateAppBadRequest{}
}

/*UpdateAppBadRequest handles this case with default header values.

Parameters are missing or invalid.
*/
type UpdateAppBadRequest struct {
	Payload *modelsv2.Error
}

func (o *UpdateAppBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appID}][%d] updateAppBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAppBadRequest) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *UpdateAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppNotFound creates a UpdateAppNotFound with default headers values
func NewUpdateAppNotFound() *UpdateAppNotFound {
	return &UpdateAppNotFound{}
}

/*UpdateAppNotFound handles this case with default header values.

The Application does not exist.
*/
type UpdateAppNotFound struct {
	Payload *modelsv2.Error
}

func (o *UpdateAppNotFound) Error() string {
	return fmt.Sprintf("[PUT /apps/{appID}][%d] updateAppNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAppNotFound) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *UpdateAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppDefault creates a UpdateAppDefault with default headers values
func NewUpdateAppDefault(code int) *UpdateAppDefault {
	return &UpdateAppDefault{
		_statusCode: code,
	}
}

/*UpdateAppDefault handles this case with default header values.

An unexpected error occurred.
*/
type UpdateAppDefault struct {
	_statusCode int

	Payload *modelsv2.Error
}

// Code gets the status code for the update app default response
func (o *UpdateAppDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAppDefault) Error() string {
	return fmt.Sprintf("[PUT /apps/{appID}][%d] UpdateApp default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAppDefault) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *UpdateAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
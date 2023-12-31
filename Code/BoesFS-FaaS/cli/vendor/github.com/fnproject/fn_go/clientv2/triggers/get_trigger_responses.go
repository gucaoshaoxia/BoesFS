// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fnproject/fn_go/modelsv2"
)

// GetTriggerReader is a Reader for the GetTrigger structure.
type GetTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTriggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTriggerOK creates a GetTriggerOK with default headers values
func NewGetTriggerOK() *GetTriggerOK {
	return &GetTriggerOK{}
}

/*GetTriggerOK handles this case with default header values.

Trigger information
*/
type GetTriggerOK struct {
	Payload *modelsv2.Trigger
}

func (o *GetTriggerOK) Error() string {
	return fmt.Sprintf("[GET /triggers/{triggerID}][%d] getTriggerOK  %+v", 200, o.Payload)
}

func (o *GetTriggerOK) GetPayload() *modelsv2.Trigger {
	return o.Payload
}

func (o *GetTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerNotFound creates a GetTriggerNotFound with default headers values
func NewGetTriggerNotFound() *GetTriggerNotFound {
	return &GetTriggerNotFound{}
}

/*GetTriggerNotFound handles this case with default header values.

The Trigger does not exist.
*/
type GetTriggerNotFound struct {
	Payload *modelsv2.Error
}

func (o *GetTriggerNotFound) Error() string {
	return fmt.Sprintf("[GET /triggers/{triggerID}][%d] getTriggerNotFound  %+v", 404, o.Payload)
}

func (o *GetTriggerNotFound) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *GetTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerDefault creates a GetTriggerDefault with default headers values
func NewGetTriggerDefault(code int) *GetTriggerDefault {
	return &GetTriggerDefault{
		_statusCode: code,
	}
}

/*GetTriggerDefault handles this case with default header values.

An unexpected error occurred.
*/
type GetTriggerDefault struct {
	_statusCode int

	Payload *modelsv2.Error
}

// Code gets the status code for the get trigger default response
func (o *GetTriggerDefault) Code() int {
	return o._statusCode
}

func (o *GetTriggerDefault) Error() string {
	return fmt.Sprintf("[GET /triggers/{triggerID}][%d] GetTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *GetTriggerDefault) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *GetTriggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

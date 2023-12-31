// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fnproject/fn_go/modelsv2"
)

// NewUpdateTriggerParams creates a new UpdateTriggerParams object
// with the default values initialized.
func NewUpdateTriggerParams() *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTriggerParamsWithTimeout creates a new UpdateTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTriggerParamsWithTimeout(timeout time.Duration) *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{

		timeout: timeout,
	}
}

// NewUpdateTriggerParamsWithContext creates a new UpdateTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTriggerParamsWithContext(ctx context.Context) *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{

		Context: ctx,
	}
}

// NewUpdateTriggerParamsWithHTTPClient creates a new UpdateTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTriggerParamsWithHTTPClient(client *http.Client) *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{
		HTTPClient: client,
	}
}

/*UpdateTriggerParams contains all the parameters to send to the API endpoint
for the update trigger operation typically these are written to a http.Request
*/
type UpdateTriggerParams struct {

	/*Body
	  Trigger data to merge into current value.

	*/
	Body *modelsv2.Trigger
	/*TriggerID
	  Opaque, unique Trigger ID.

	*/
	TriggerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update trigger params
func (o *UpdateTriggerParams) WithTimeout(timeout time.Duration) *UpdateTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update trigger params
func (o *UpdateTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update trigger params
func (o *UpdateTriggerParams) WithContext(ctx context.Context) *UpdateTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update trigger params
func (o *UpdateTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update trigger params
func (o *UpdateTriggerParams) WithHTTPClient(client *http.Client) *UpdateTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update trigger params
func (o *UpdateTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update trigger params
func (o *UpdateTriggerParams) WithBody(body *modelsv2.Trigger) *UpdateTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update trigger params
func (o *UpdateTriggerParams) SetBody(body *modelsv2.Trigger) {
	o.Body = body
}

// WithTriggerID adds the triggerID to the update trigger params
func (o *UpdateTriggerParams) WithTriggerID(triggerID string) *UpdateTriggerParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the update trigger params
func (o *UpdateTriggerParams) SetTriggerID(triggerID string) {
	o.TriggerID = triggerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param triggerID
	if err := r.SetPathParam("triggerID", o.TriggerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package fns

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

// NewUpdateFnParams creates a new UpdateFnParams object
// with the default values initialized.
func NewUpdateFnParams() *UpdateFnParams {
	var ()
	return &UpdateFnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFnParamsWithTimeout creates a new UpdateFnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFnParamsWithTimeout(timeout time.Duration) *UpdateFnParams {
	var ()
	return &UpdateFnParams{

		timeout: timeout,
	}
}

// NewUpdateFnParamsWithContext creates a new UpdateFnParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFnParamsWithContext(ctx context.Context) *UpdateFnParams {
	var ()
	return &UpdateFnParams{

		Context: ctx,
	}
}

// NewUpdateFnParamsWithHTTPClient creates a new UpdateFnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFnParamsWithHTTPClient(client *http.Client) *UpdateFnParams {
	var ()
	return &UpdateFnParams{
		HTTPClient: client,
	}
}

/*UpdateFnParams contains all the parameters to send to the API endpoint
for the update fn operation typically these are written to a http.Request
*/
type UpdateFnParams struct {

	/*Body
	  Function data to merge with current values.

	*/
	Body *modelsv2.Fn
	/*FnID
	  Opaque, unique Function ID.

	*/
	FnID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update fn params
func (o *UpdateFnParams) WithTimeout(timeout time.Duration) *UpdateFnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update fn params
func (o *UpdateFnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update fn params
func (o *UpdateFnParams) WithContext(ctx context.Context) *UpdateFnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update fn params
func (o *UpdateFnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update fn params
func (o *UpdateFnParams) WithHTTPClient(client *http.Client) *UpdateFnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update fn params
func (o *UpdateFnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update fn params
func (o *UpdateFnParams) WithBody(body *modelsv2.Fn) *UpdateFnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update fn params
func (o *UpdateFnParams) SetBody(body *modelsv2.Fn) {
	o.Body = body
}

// WithFnID adds the fnID to the update fn params
func (o *UpdateFnParams) WithFnID(fnID string) *UpdateFnParams {
	o.SetFnID(fnID)
	return o
}

// SetFnID adds the fnId to the update fn params
func (o *UpdateFnParams) SetFnID(fnID string) {
	o.FnID = fnID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param fnID
	if err := r.SetPathParam("fnID", o.FnID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

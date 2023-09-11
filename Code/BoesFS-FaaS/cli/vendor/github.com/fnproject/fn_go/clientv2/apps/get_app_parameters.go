// Code generated by go-swagger; DO NOT EDIT.

package apps

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
)

// NewGetAppParams creates a new GetAppParams object
// with the default values initialized.
func NewGetAppParams() *GetAppParams {
	var ()
	return &GetAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppParamsWithTimeout creates a new GetAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppParamsWithTimeout(timeout time.Duration) *GetAppParams {
	var ()
	return &GetAppParams{

		timeout: timeout,
	}
}

// NewGetAppParamsWithContext creates a new GetAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppParamsWithContext(ctx context.Context) *GetAppParams {
	var ()
	return &GetAppParams{

		Context: ctx,
	}
}

// NewGetAppParamsWithHTTPClient creates a new GetAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppParamsWithHTTPClient(client *http.Client) *GetAppParams {
	var ()
	return &GetAppParams{
		HTTPClient: client,
	}
}

/*GetAppParams contains all the parameters to send to the API endpoint
for the get app operation typically these are written to a http.Request
*/
type GetAppParams struct {

	/*AppID
	  Opaque, unique Application ID.

	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get app params
func (o *GetAppParams) WithTimeout(timeout time.Duration) *GetAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app params
func (o *GetAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app params
func (o *GetAppParams) WithContext(ctx context.Context) *GetAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app params
func (o *GetAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app params
func (o *GetAppParams) WithHTTPClient(client *http.Client) *GetAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app params
func (o *GetAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app params
func (o *GetAppParams) WithAppID(appID string) *GetAppParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app params
func (o *GetAppParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appID
	if err := r.SetPathParam("appID", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

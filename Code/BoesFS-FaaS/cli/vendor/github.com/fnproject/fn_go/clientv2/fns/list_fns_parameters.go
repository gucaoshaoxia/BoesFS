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
	"github.com/go-openapi/swag"
)

// NewListFnsParams creates a new ListFnsParams object
// with the default values initialized.
func NewListFnsParams() *ListFnsParams {
	var ()
	return &ListFnsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFnsParamsWithTimeout creates a new ListFnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFnsParamsWithTimeout(timeout time.Duration) *ListFnsParams {
	var ()
	return &ListFnsParams{

		timeout: timeout,
	}
}

// NewListFnsParamsWithContext creates a new ListFnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFnsParamsWithContext(ctx context.Context) *ListFnsParams {
	var ()
	return &ListFnsParams{

		Context: ctx,
	}
}

// NewListFnsParamsWithHTTPClient creates a new ListFnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFnsParamsWithHTTPClient(client *http.Client) *ListFnsParams {
	var ()
	return &ListFnsParams{
		HTTPClient: client,
	}
}

/*ListFnsParams contains all the parameters to send to the API endpoint
for the list fns operation typically these are written to a http.Request
*/
type ListFnsParams struct {

	/*AppID
	  Application ID.

	*/
	AppID *string
	/*Cursor
	  Cursor from previous response.next_cursor to begin results after, if any.

	*/
	Cursor *string
	/*Name
	  Function name to filter by

	*/
	Name *string
	/*PerPage
	  Number of results to return, defaults to 30. Max of 100.

	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list fns params
func (o *ListFnsParams) WithTimeout(timeout time.Duration) *ListFnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list fns params
func (o *ListFnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list fns params
func (o *ListFnsParams) WithContext(ctx context.Context) *ListFnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list fns params
func (o *ListFnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list fns params
func (o *ListFnsParams) WithHTTPClient(client *http.Client) *ListFnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list fns params
func (o *ListFnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the list fns params
func (o *ListFnsParams) WithAppID(appID *string) *ListFnsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the list fns params
func (o *ListFnsParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WithCursor adds the cursor to the list fns params
func (o *ListFnsParams) WithCursor(cursor *string) *ListFnsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list fns params
func (o *ListFnsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithName adds the name to the list fns params
func (o *ListFnsParams) WithName(name *string) *ListFnsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list fns params
func (o *ListFnsParams) SetName(name *string) {
	o.Name = name
}

// WithPerPage adds the perPage to the list fns params
func (o *ListFnsParams) WithPerPage(perPage *int64) *ListFnsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the list fns params
func (o *ListFnsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *ListFnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppID != nil {

		// query param app_id
		var qrAppID string
		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {
			if err := r.SetQueryParam("app_id", qAppID); err != nil {
				return err
			}
		}

	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

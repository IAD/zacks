// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTickerHistoryParams creates a new GetTickerHistoryParams object
// with the default values initialized.
func NewGetTickerHistoryParams() *GetTickerHistoryParams {
	var ()
	return &GetTickerHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTickerHistoryParamsWithTimeout creates a new GetTickerHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTickerHistoryParamsWithTimeout(timeout time.Duration) *GetTickerHistoryParams {
	var ()
	return &GetTickerHistoryParams{

		timeout: timeout,
	}
}

// NewGetTickerHistoryParamsWithContext creates a new GetTickerHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTickerHistoryParamsWithContext(ctx context.Context) *GetTickerHistoryParams {
	var ()
	return &GetTickerHistoryParams{

		Context: ctx,
	}
}

// NewGetTickerHistoryParamsWithHTTPClient creates a new GetTickerHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTickerHistoryParamsWithHTTPClient(client *http.Client) *GetTickerHistoryParams {
	var ()
	return &GetTickerHistoryParams{
		HTTPClient: client,
	}
}

/*GetTickerHistoryParams contains all the parameters to send to the API endpoint
for the get ticker history operation typically these are written to a http.Request
*/
type GetTickerHistoryParams struct {

	/*Ticker*/
	Ticker string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ticker history params
func (o *GetTickerHistoryParams) WithTimeout(timeout time.Duration) *GetTickerHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ticker history params
func (o *GetTickerHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ticker history params
func (o *GetTickerHistoryParams) WithContext(ctx context.Context) *GetTickerHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ticker history params
func (o *GetTickerHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ticker history params
func (o *GetTickerHistoryParams) WithHTTPClient(client *http.Client) *GetTickerHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ticker history params
func (o *GetTickerHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTicker adds the ticker to the get ticker history params
func (o *GetTickerHistoryParams) WithTicker(ticker string) *GetTickerHistoryParams {
	o.SetTicker(ticker)
	return o
}

// SetTicker adds the ticker to the get ticker history params
func (o *GetTickerHistoryParams) SetTicker(ticker string) {
	o.Ticker = ticker
}

// WriteToRequest writes these params to a swagger request
func (o *GetTickerHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ticker
	if err := r.SetPathParam("ticker", o.Ticker); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

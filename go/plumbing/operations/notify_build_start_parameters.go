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
	"github.com/go-openapi/strfmt"
)

// NewNotifyBuildStartParams creates a new NotifyBuildStartParams object
// with the default values initialized.
func NewNotifyBuildStartParams() *NotifyBuildStartParams {
	var ()
	return &NotifyBuildStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotifyBuildStartParamsWithTimeout creates a new NotifyBuildStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotifyBuildStartParamsWithTimeout(timeout time.Duration) *NotifyBuildStartParams {
	var ()
	return &NotifyBuildStartParams{

		timeout: timeout,
	}
}

// NewNotifyBuildStartParamsWithContext creates a new NotifyBuildStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotifyBuildStartParamsWithContext(ctx context.Context) *NotifyBuildStartParams {
	var ()
	return &NotifyBuildStartParams{

		Context: ctx,
	}
}

// NewNotifyBuildStartParamsWithHTTPClient creates a new NotifyBuildStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotifyBuildStartParamsWithHTTPClient(client *http.Client) *NotifyBuildStartParams {
	var ()
	return &NotifyBuildStartParams{
		HTTPClient: client,
	}
}

/*
NotifyBuildStartParams contains all the parameters to send to the API endpoint
for the notify build start operation typically these are written to a http.Request
*/
type NotifyBuildStartParams struct {

	/*BuildID*/
	BuildID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notify build start params
func (o *NotifyBuildStartParams) WithTimeout(timeout time.Duration) *NotifyBuildStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notify build start params
func (o *NotifyBuildStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notify build start params
func (o *NotifyBuildStartParams) WithContext(ctx context.Context) *NotifyBuildStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notify build start params
func (o *NotifyBuildStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notify build start params
func (o *NotifyBuildStartParams) WithHTTPClient(client *http.Client) *NotifyBuildStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notify build start params
func (o *NotifyBuildStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the notify build start params
func (o *NotifyBuildStartParams) WithBuildID(buildID string) *NotifyBuildStartParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the notify build start params
func (o *NotifyBuildStartParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WriteToRequest writes these params to a swagger request
func (o *NotifyBuildStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

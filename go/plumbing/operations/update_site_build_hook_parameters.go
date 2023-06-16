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

	"github.com/netlify/open-api/v2/go/models"
)

// NewUpdateSiteBuildHookParams creates a new UpdateSiteBuildHookParams object
// with the default values initialized.
func NewUpdateSiteBuildHookParams() *UpdateSiteBuildHookParams {
	var ()
	return &UpdateSiteBuildHookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSiteBuildHookParamsWithTimeout creates a new UpdateSiteBuildHookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSiteBuildHookParamsWithTimeout(timeout time.Duration) *UpdateSiteBuildHookParams {
	var ()
	return &UpdateSiteBuildHookParams{

		timeout: timeout,
	}
}

// NewUpdateSiteBuildHookParamsWithContext creates a new UpdateSiteBuildHookParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSiteBuildHookParamsWithContext(ctx context.Context) *UpdateSiteBuildHookParams {
	var ()
	return &UpdateSiteBuildHookParams{

		Context: ctx,
	}
}

// NewUpdateSiteBuildHookParamsWithHTTPClient creates a new UpdateSiteBuildHookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSiteBuildHookParamsWithHTTPClient(client *http.Client) *UpdateSiteBuildHookParams {
	var ()
	return &UpdateSiteBuildHookParams{
		HTTPClient: client,
	}
}

/*UpdateSiteBuildHookParams contains all the parameters to send to the API endpoint
for the update site build hook operation typically these are written to a http.Request
*/
type UpdateSiteBuildHookParams struct {

	/*BuildHook*/
	BuildHook *models.BuildHookSetup
	/*ID*/
	ID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update site build hook params
func (o *UpdateSiteBuildHookParams) WithTimeout(timeout time.Duration) *UpdateSiteBuildHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update site build hook params
func (o *UpdateSiteBuildHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update site build hook params
func (o *UpdateSiteBuildHookParams) WithContext(ctx context.Context) *UpdateSiteBuildHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update site build hook params
func (o *UpdateSiteBuildHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update site build hook params
func (o *UpdateSiteBuildHookParams) WithHTTPClient(client *http.Client) *UpdateSiteBuildHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update site build hook params
func (o *UpdateSiteBuildHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildHook adds the buildHook to the update site build hook params
func (o *UpdateSiteBuildHookParams) WithBuildHook(buildHook *models.BuildHookSetup) *UpdateSiteBuildHookParams {
	o.SetBuildHook(buildHook)
	return o
}

// SetBuildHook adds the buildHook to the update site build hook params
func (o *UpdateSiteBuildHookParams) SetBuildHook(buildHook *models.BuildHookSetup) {
	o.BuildHook = buildHook
}

// WithID adds the id to the update site build hook params
func (o *UpdateSiteBuildHookParams) WithID(id string) *UpdateSiteBuildHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update site build hook params
func (o *UpdateSiteBuildHookParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the update site build hook params
func (o *UpdateSiteBuildHookParams) WithSiteID(siteID string) *UpdateSiteBuildHookParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update site build hook params
func (o *UpdateSiteBuildHookParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSiteBuildHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BuildHook != nil {
		if err := r.SetBodyParam(o.BuildHook); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

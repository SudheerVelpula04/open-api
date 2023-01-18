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
	"github.com/go-openapi/swag"
)

// NewListSiteDeploysParams creates a new ListSiteDeploysParams object
// with the default values initialized.
func NewListSiteDeploysParams() *ListSiteDeploysParams {
	var ()
	return &ListSiteDeploysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSiteDeploysParamsWithTimeout creates a new ListSiteDeploysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSiteDeploysParamsWithTimeout(timeout time.Duration) *ListSiteDeploysParams {
	var ()
	return &ListSiteDeploysParams{

		timeout: timeout,
	}
}

// NewListSiteDeploysParamsWithContext creates a new ListSiteDeploysParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSiteDeploysParamsWithContext(ctx context.Context) *ListSiteDeploysParams {
	var ()
	return &ListSiteDeploysParams{

		Context: ctx,
	}
}

// NewListSiteDeploysParamsWithHTTPClient creates a new ListSiteDeploysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSiteDeploysParamsWithHTTPClient(client *http.Client) *ListSiteDeploysParams {
	var ()
	return &ListSiteDeploysParams{
		HTTPClient: client,
	}
}

/*ListSiteDeploysParams contains all the parameters to send to the API endpoint
for the list site deploys operation typically these are written to a http.Request
*/
type ListSiteDeploysParams struct {

	/*Branch*/
	Branch *string
	/*DeployPreviews*/
	DeployPreviews *bool
	/*LatestPublished*/
	LatestPublished *bool
	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32
	/*Production*/
	Production *bool
	/*SiteID*/
	SiteID string
	/*State*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list site deploys params
func (o *ListSiteDeploysParams) WithTimeout(timeout time.Duration) *ListSiteDeploysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list site deploys params
func (o *ListSiteDeploysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list site deploys params
func (o *ListSiteDeploysParams) WithContext(ctx context.Context) *ListSiteDeploysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list site deploys params
func (o *ListSiteDeploysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list site deploys params
func (o *ListSiteDeploysParams) WithHTTPClient(client *http.Client) *ListSiteDeploysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list site deploys params
func (o *ListSiteDeploysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranch adds the branch to the list site deploys params
func (o *ListSiteDeploysParams) WithBranch(branch *string) *ListSiteDeploysParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the list site deploys params
func (o *ListSiteDeploysParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithDeployPreviews adds the deployPreviews to the list site deploys params
func (o *ListSiteDeploysParams) WithDeployPreviews(deployPreviews *bool) *ListSiteDeploysParams {
	o.SetDeployPreviews(deployPreviews)
	return o
}

// SetDeployPreviews adds the deployPreviews to the list site deploys params
func (o *ListSiteDeploysParams) SetDeployPreviews(deployPreviews *bool) {
	o.DeployPreviews = deployPreviews
}

// WithLatestPublished adds the latestPublished to the list site deploys params
func (o *ListSiteDeploysParams) WithLatestPublished(latestPublished *bool) *ListSiteDeploysParams {
	o.SetLatestPublished(latestPublished)
	return o
}

// SetLatestPublished adds the latestPublished to the list site deploys params
func (o *ListSiteDeploysParams) SetLatestPublished(latestPublished *bool) {
	o.LatestPublished = latestPublished
}

// WithPage adds the page to the list site deploys params
func (o *ListSiteDeploysParams) WithPage(page *int32) *ListSiteDeploysParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list site deploys params
func (o *ListSiteDeploysParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the list site deploys params
func (o *ListSiteDeploysParams) WithPerPage(perPage *int32) *ListSiteDeploysParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the list site deploys params
func (o *ListSiteDeploysParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithProduction adds the production to the list site deploys params
func (o *ListSiteDeploysParams) WithProduction(production *bool) *ListSiteDeploysParams {
	o.SetProduction(production)
	return o
}

// SetProduction adds the production to the list site deploys params
func (o *ListSiteDeploysParams) SetProduction(production *bool) {
	o.Production = production
}

// WithSiteID adds the siteID to the list site deploys params
func (o *ListSiteDeploysParams) WithSiteID(siteID string) *ListSiteDeploysParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the list site deploys params
func (o *ListSiteDeploysParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithState adds the state to the list site deploys params
func (o *ListSiteDeploysParams) WithState(state *string) *ListSiteDeploysParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the list site deploys params
func (o *ListSiteDeploysParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteDeploysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Branch != nil {

		// query param branch
		var qrBranch string
		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {
			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}

	}

	if o.DeployPreviews != nil {

		// query param deploy-previews
		var qrDeployPreviews bool
		if o.DeployPreviews != nil {
			qrDeployPreviews = *o.DeployPreviews
		}
		qDeployPreviews := swag.FormatBool(qrDeployPreviews)
		if qDeployPreviews != "" {
			if err := r.SetQueryParam("deploy-previews", qDeployPreviews); err != nil {
				return err
			}
		}

	}

	if o.LatestPublished != nil {

		// query param latest-published
		var qrLatestPublished bool
		if o.LatestPublished != nil {
			qrLatestPublished = *o.LatestPublished
		}
		qLatestPublished := swag.FormatBool(qrLatestPublished)
		if qLatestPublished != "" {
			if err := r.SetQueryParam("latest-published", qLatestPublished); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.Production != nil {

		// query param production
		var qrProduction bool
		if o.Production != nil {
			qrProduction = *o.Production
		}
		qProduction := swag.FormatBool(qrProduction)
		if qProduction != "" {
			if err := r.SetQueryParam("production", qProduction); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

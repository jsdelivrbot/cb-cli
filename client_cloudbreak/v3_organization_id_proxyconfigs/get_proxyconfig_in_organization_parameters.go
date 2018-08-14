// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProxyconfigInOrganizationParams creates a new GetProxyconfigInOrganizationParams object
// with the default values initialized.
func NewGetProxyconfigInOrganizationParams() *GetProxyconfigInOrganizationParams {
	var ()
	return &GetProxyconfigInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyconfigInOrganizationParamsWithTimeout creates a new GetProxyconfigInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProxyconfigInOrganizationParamsWithTimeout(timeout time.Duration) *GetProxyconfigInOrganizationParams {
	var ()
	return &GetProxyconfigInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetProxyconfigInOrganizationParamsWithContext creates a new GetProxyconfigInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProxyconfigInOrganizationParamsWithContext(ctx context.Context) *GetProxyconfigInOrganizationParams {
	var ()
	return &GetProxyconfigInOrganizationParams{

		Context: ctx,
	}
}

// NewGetProxyconfigInOrganizationParamsWithHTTPClient creates a new GetProxyconfigInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProxyconfigInOrganizationParamsWithHTTPClient(client *http.Client) *GetProxyconfigInOrganizationParams {
	var ()
	return &GetProxyconfigInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetProxyconfigInOrganizationParams contains all the parameters to send to the API endpoint
for the get proxyconfig in organization operation typically these are written to a http.Request
*/
type GetProxyconfigInOrganizationParams struct {

	/*Name*/
	Name string
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) WithTimeout(timeout time.Duration) *GetProxyconfigInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) WithContext(ctx context.Context) *GetProxyconfigInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) WithHTTPClient(client *http.Client) *GetProxyconfigInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) WithName(name string) *GetProxyconfigInOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) WithOrganizationID(organizationID int64) *GetProxyconfigInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get proxyconfig in organization params
func (o *GetProxyconfigInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyconfigInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

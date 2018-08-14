// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_imagecatalogs

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

// NewGetImageCatalogRequestFromNameInOrganizationParams creates a new GetImageCatalogRequestFromNameInOrganizationParams object
// with the default values initialized.
func NewGetImageCatalogRequestFromNameInOrganizationParams() *GetImageCatalogRequestFromNameInOrganizationParams {
	var ()
	return &GetImageCatalogRequestFromNameInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageCatalogRequestFromNameInOrganizationParamsWithTimeout creates a new GetImageCatalogRequestFromNameInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageCatalogRequestFromNameInOrganizationParamsWithTimeout(timeout time.Duration) *GetImageCatalogRequestFromNameInOrganizationParams {
	var ()
	return &GetImageCatalogRequestFromNameInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetImageCatalogRequestFromNameInOrganizationParamsWithContext creates a new GetImageCatalogRequestFromNameInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageCatalogRequestFromNameInOrganizationParamsWithContext(ctx context.Context) *GetImageCatalogRequestFromNameInOrganizationParams {
	var ()
	return &GetImageCatalogRequestFromNameInOrganizationParams{

		Context: ctx,
	}
}

// NewGetImageCatalogRequestFromNameInOrganizationParamsWithHTTPClient creates a new GetImageCatalogRequestFromNameInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageCatalogRequestFromNameInOrganizationParamsWithHTTPClient(client *http.Client) *GetImageCatalogRequestFromNameInOrganizationParams {
	var ()
	return &GetImageCatalogRequestFromNameInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetImageCatalogRequestFromNameInOrganizationParams contains all the parameters to send to the API endpoint
for the get image catalog request from name in organization operation typically these are written to a http.Request
*/
type GetImageCatalogRequestFromNameInOrganizationParams struct {

	/*Name*/
	Name string
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) WithTimeout(timeout time.Duration) *GetImageCatalogRequestFromNameInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) WithContext(ctx context.Context) *GetImageCatalogRequestFromNameInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) WithHTTPClient(client *http.Client) *GetImageCatalogRequestFromNameInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) WithName(name string) *GetImageCatalogRequestFromNameInOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) WithOrganizationID(organizationID int64) *GetImageCatalogRequestFromNameInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get image catalog request from name in organization params
func (o *GetImageCatalogRequestFromNameInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageCatalogRequestFromNameInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

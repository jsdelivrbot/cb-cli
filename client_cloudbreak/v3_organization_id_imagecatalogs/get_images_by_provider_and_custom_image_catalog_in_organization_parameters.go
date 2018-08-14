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

// NewGetImagesByProviderAndCustomImageCatalogInOrganizationParams creates a new GetImagesByProviderAndCustomImageCatalogInOrganizationParams object
// with the default values initialized.
func NewGetImagesByProviderAndCustomImageCatalogInOrganizationParams() *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	var ()
	return &GetImagesByProviderAndCustomImageCatalogInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesByProviderAndCustomImageCatalogInOrganizationParamsWithTimeout creates a new GetImagesByProviderAndCustomImageCatalogInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesByProviderAndCustomImageCatalogInOrganizationParamsWithTimeout(timeout time.Duration) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	var ()
	return &GetImagesByProviderAndCustomImageCatalogInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetImagesByProviderAndCustomImageCatalogInOrganizationParamsWithContext creates a new GetImagesByProviderAndCustomImageCatalogInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesByProviderAndCustomImageCatalogInOrganizationParamsWithContext(ctx context.Context) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	var ()
	return &GetImagesByProviderAndCustomImageCatalogInOrganizationParams{

		Context: ctx,
	}
}

// NewGetImagesByProviderAndCustomImageCatalogInOrganizationParamsWithHTTPClient creates a new GetImagesByProviderAndCustomImageCatalogInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImagesByProviderAndCustomImageCatalogInOrganizationParamsWithHTTPClient(client *http.Client) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	var ()
	return &GetImagesByProviderAndCustomImageCatalogInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetImagesByProviderAndCustomImageCatalogInOrganizationParams contains all the parameters to send to the API endpoint
for the get images by provider and custom image catalog in organization operation typically these are written to a http.Request
*/
type GetImagesByProviderAndCustomImageCatalogInOrganizationParams struct {

	/*Name*/
	Name string
	/*OrganizationID*/
	OrganizationID int64
	/*Platform*/
	Platform string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WithTimeout(timeout time.Duration) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WithContext(ctx context.Context) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WithHTTPClient(client *http.Client) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WithName(name string) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WithOrganizationID(organizationID int64) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WithPlatform adds the platform to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WithPlatform(platform string) *GetImagesByProviderAndCustomImageCatalogInOrganizationParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the get images by provider and custom image catalog in organization params
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) SetPlatform(platform string) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesByProviderAndCustomImageCatalogInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param platform
	if err := r.SetPathParam("platform", o.Platform); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

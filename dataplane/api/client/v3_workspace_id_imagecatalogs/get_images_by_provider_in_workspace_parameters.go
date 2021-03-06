// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_imagecatalogs

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

// NewGetImagesByProviderInWorkspaceParams creates a new GetImagesByProviderInWorkspaceParams object
// with the default values initialized.
func NewGetImagesByProviderInWorkspaceParams() *GetImagesByProviderInWorkspaceParams {
	var ()
	return &GetImagesByProviderInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesByProviderInWorkspaceParamsWithTimeout creates a new GetImagesByProviderInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesByProviderInWorkspaceParamsWithTimeout(timeout time.Duration) *GetImagesByProviderInWorkspaceParams {
	var ()
	return &GetImagesByProviderInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetImagesByProviderInWorkspaceParamsWithContext creates a new GetImagesByProviderInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesByProviderInWorkspaceParamsWithContext(ctx context.Context) *GetImagesByProviderInWorkspaceParams {
	var ()
	return &GetImagesByProviderInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetImagesByProviderInWorkspaceParamsWithHTTPClient creates a new GetImagesByProviderInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImagesByProviderInWorkspaceParamsWithHTTPClient(client *http.Client) *GetImagesByProviderInWorkspaceParams {
	var ()
	return &GetImagesByProviderInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetImagesByProviderInWorkspaceParams contains all the parameters to send to the API endpoint
for the get images by provider in workspace operation typically these are written to a http.Request
*/
type GetImagesByProviderInWorkspaceParams struct {

	/*Platform*/
	Platform string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) WithTimeout(timeout time.Duration) *GetImagesByProviderInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) WithContext(ctx context.Context) *GetImagesByProviderInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) WithHTTPClient(client *http.Client) *GetImagesByProviderInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlatform adds the platform to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) WithPlatform(platform string) *GetImagesByProviderInWorkspaceParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) SetPlatform(platform string) {
	o.Platform = platform
}

// WithWorkspaceID adds the workspaceID to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetImagesByProviderInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get images by provider in workspace params
func (o *GetImagesByProviderInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesByProviderInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param platform
	if err := r.SetPathParam("platform", o.Platform); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

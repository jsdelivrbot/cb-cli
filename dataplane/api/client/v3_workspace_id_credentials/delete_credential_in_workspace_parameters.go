// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_credentials

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

// NewDeleteCredentialInWorkspaceParams creates a new DeleteCredentialInWorkspaceParams object
// with the default values initialized.
func NewDeleteCredentialInWorkspaceParams() *DeleteCredentialInWorkspaceParams {
	var ()
	return &DeleteCredentialInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCredentialInWorkspaceParamsWithTimeout creates a new DeleteCredentialInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCredentialInWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteCredentialInWorkspaceParams {
	var ()
	return &DeleteCredentialInWorkspaceParams{

		timeout: timeout,
	}
}

// NewDeleteCredentialInWorkspaceParamsWithContext creates a new DeleteCredentialInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCredentialInWorkspaceParamsWithContext(ctx context.Context) *DeleteCredentialInWorkspaceParams {
	var ()
	return &DeleteCredentialInWorkspaceParams{

		Context: ctx,
	}
}

// NewDeleteCredentialInWorkspaceParamsWithHTTPClient creates a new DeleteCredentialInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCredentialInWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteCredentialInWorkspaceParams {
	var ()
	return &DeleteCredentialInWorkspaceParams{
		HTTPClient: client,
	}
}

/*DeleteCredentialInWorkspaceParams contains all the parameters to send to the API endpoint
for the delete credential in workspace operation typically these are written to a http.Request
*/
type DeleteCredentialInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteCredentialInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) WithContext(ctx context.Context) *DeleteCredentialInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteCredentialInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) WithName(name string) *DeleteCredentialInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) WithWorkspaceID(workspaceID int64) *DeleteCredentialInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete credential in workspace params
func (o *DeleteCredentialInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCredentialInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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

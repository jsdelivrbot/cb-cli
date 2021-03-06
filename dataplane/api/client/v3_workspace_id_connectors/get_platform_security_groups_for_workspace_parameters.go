// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_connectors

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewGetPlatformSecurityGroupsForWorkspaceParams creates a new GetPlatformSecurityGroupsForWorkspaceParams object
// with the default values initialized.
func NewGetPlatformSecurityGroupsForWorkspaceParams() *GetPlatformSecurityGroupsForWorkspaceParams {
	var ()
	return &GetPlatformSecurityGroupsForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformSecurityGroupsForWorkspaceParamsWithTimeout creates a new GetPlatformSecurityGroupsForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformSecurityGroupsForWorkspaceParamsWithTimeout(timeout time.Duration) *GetPlatformSecurityGroupsForWorkspaceParams {
	var ()
	return &GetPlatformSecurityGroupsForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetPlatformSecurityGroupsForWorkspaceParamsWithContext creates a new GetPlatformSecurityGroupsForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformSecurityGroupsForWorkspaceParamsWithContext(ctx context.Context) *GetPlatformSecurityGroupsForWorkspaceParams {
	var ()
	return &GetPlatformSecurityGroupsForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetPlatformSecurityGroupsForWorkspaceParamsWithHTTPClient creates a new GetPlatformSecurityGroupsForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformSecurityGroupsForWorkspaceParamsWithHTTPClient(client *http.Client) *GetPlatformSecurityGroupsForWorkspaceParams {
	var ()
	return &GetPlatformSecurityGroupsForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetPlatformSecurityGroupsForWorkspaceParams contains all the parameters to send to the API endpoint
for the get platform security groups for workspace operation typically these are written to a http.Request
*/
type GetPlatformSecurityGroupsForWorkspaceParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) WithTimeout(timeout time.Duration) *GetPlatformSecurityGroupsForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) WithContext(ctx context.Context) *GetPlatformSecurityGroupsForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) WithHTTPClient(client *http.Client) *GetPlatformSecurityGroupsForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) WithBody(body *model.PlatformResourceRequestJSON) *GetPlatformSecurityGroupsForWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetPlatformSecurityGroupsForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get platform security groups for workspace params
func (o *GetPlatformSecurityGroupsForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformSecurityGroupsForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

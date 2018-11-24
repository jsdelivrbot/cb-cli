// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

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

	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
)

// NewChangeImageV3Params creates a new ChangeImageV3Params object
// with the default values initialized.
func NewChangeImageV3Params() *ChangeImageV3Params {
	var ()
	return &ChangeImageV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeImageV3ParamsWithTimeout creates a new ChangeImageV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeImageV3ParamsWithTimeout(timeout time.Duration) *ChangeImageV3Params {
	var ()
	return &ChangeImageV3Params{

		timeout: timeout,
	}
}

// NewChangeImageV3ParamsWithContext creates a new ChangeImageV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewChangeImageV3ParamsWithContext(ctx context.Context) *ChangeImageV3Params {
	var ()
	return &ChangeImageV3Params{

		Context: ctx,
	}
}

// NewChangeImageV3ParamsWithHTTPClient creates a new ChangeImageV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeImageV3ParamsWithHTTPClient(client *http.Client) *ChangeImageV3Params {
	var ()
	return &ChangeImageV3Params{
		HTTPClient: client,
	}
}

/*ChangeImageV3Params contains all the parameters to send to the API endpoint
for the change image v3 operation typically these are written to a http.Request
*/
type ChangeImageV3Params struct {

	/*Body*/
	Body *model.StackImageChangeRequest
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change image v3 params
func (o *ChangeImageV3Params) WithTimeout(timeout time.Duration) *ChangeImageV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change image v3 params
func (o *ChangeImageV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change image v3 params
func (o *ChangeImageV3Params) WithContext(ctx context.Context) *ChangeImageV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change image v3 params
func (o *ChangeImageV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change image v3 params
func (o *ChangeImageV3Params) WithHTTPClient(client *http.Client) *ChangeImageV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change image v3 params
func (o *ChangeImageV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change image v3 params
func (o *ChangeImageV3Params) WithBody(body *model.StackImageChangeRequest) *ChangeImageV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change image v3 params
func (o *ChangeImageV3Params) SetBody(body *model.StackImageChangeRequest) {
	o.Body = body
}

// WithName adds the name to the change image v3 params
func (o *ChangeImageV3Params) WithName(name string) *ChangeImageV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the change image v3 params
func (o *ChangeImageV3Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the change image v3 params
func (o *ChangeImageV3Params) WithWorkspaceID(workspaceID int64) *ChangeImageV3Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the change image v3 params
func (o *ChangeImageV3Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeImageV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(model.StackImageChangeRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

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
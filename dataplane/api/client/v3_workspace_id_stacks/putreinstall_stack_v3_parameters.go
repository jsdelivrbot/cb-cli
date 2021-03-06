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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPutreinstallStackV3Params creates a new PutreinstallStackV3Params object
// with the default values initialized.
func NewPutreinstallStackV3Params() *PutreinstallStackV3Params {
	var ()
	return &PutreinstallStackV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutreinstallStackV3ParamsWithTimeout creates a new PutreinstallStackV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutreinstallStackV3ParamsWithTimeout(timeout time.Duration) *PutreinstallStackV3Params {
	var ()
	return &PutreinstallStackV3Params{

		timeout: timeout,
	}
}

// NewPutreinstallStackV3ParamsWithContext creates a new PutreinstallStackV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutreinstallStackV3ParamsWithContext(ctx context.Context) *PutreinstallStackV3Params {
	var ()
	return &PutreinstallStackV3Params{

		Context: ctx,
	}
}

// NewPutreinstallStackV3ParamsWithHTTPClient creates a new PutreinstallStackV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutreinstallStackV3ParamsWithHTTPClient(client *http.Client) *PutreinstallStackV3Params {
	var ()
	return &PutreinstallStackV3Params{
		HTTPClient: client,
	}
}

/*PutreinstallStackV3Params contains all the parameters to send to the API endpoint
for the putreinstall stack v3 operation typically these are written to a http.Request
*/
type PutreinstallStackV3Params struct {

	/*Body*/
	Body *model.ReinstallRequestV2
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) WithTimeout(timeout time.Duration) *PutreinstallStackV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) WithContext(ctx context.Context) *PutreinstallStackV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) WithHTTPClient(client *http.Client) *PutreinstallStackV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) WithBody(body *model.ReinstallRequestV2) *PutreinstallStackV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) SetBody(body *model.ReinstallRequestV2) {
	o.Body = body
}

// WithName adds the name to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) WithName(name string) *PutreinstallStackV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) WithWorkspaceID(workspaceID int64) *PutreinstallStackV3Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the putreinstall stack v3 params
func (o *PutreinstallStackV3Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutreinstallStackV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

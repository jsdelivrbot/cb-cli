// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_flexsubscriptions

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

// NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParams creates a new PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams object
// with the default values initialized.
func NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParams() *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	var ()
	return &PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParamsWithTimeout creates a new PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParamsWithTimeout(timeout time.Duration) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	var ()
	return &PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams{

		timeout: timeout,
	}
}

// NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParamsWithContext creates a new PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParamsWithContext(ctx context.Context) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	var ()
	return &PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams{

		Context: ctx,
	}
}

// NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParamsWithHTTPClient creates a new PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUsedForControllerFlexSubscriptionByNameInWorkspaceParamsWithHTTPClient(client *http.Client) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	var ()
	return &PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams{
		HTTPClient: client,
	}
}

/*PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams contains all the parameters to send to the API endpoint
for the put used for controller flex subscription by name in workspace operation typically these are written to a http.Request
*/
type PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) WithTimeout(timeout time.Duration) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) WithContext(ctx context.Context) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) WithHTTPClient(client *http.Client) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) WithName(name string) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) WithWorkspaceID(workspaceID int64) *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the put used for controller flex subscription by name in workspace params
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutUsedForControllerFlexSubscriptionByNameInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

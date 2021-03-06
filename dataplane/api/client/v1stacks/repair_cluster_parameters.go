// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

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

// NewRepairClusterParams creates a new RepairClusterParams object
// with the default values initialized.
func NewRepairClusterParams() *RepairClusterParams {
	var ()
	return &RepairClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepairClusterParamsWithTimeout creates a new RepairClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepairClusterParamsWithTimeout(timeout time.Duration) *RepairClusterParams {
	var ()
	return &RepairClusterParams{

		timeout: timeout,
	}
}

// NewRepairClusterParamsWithContext creates a new RepairClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepairClusterParamsWithContext(ctx context.Context) *RepairClusterParams {
	var ()
	return &RepairClusterParams{

		Context: ctx,
	}
}

// NewRepairClusterParamsWithHTTPClient creates a new RepairClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepairClusterParamsWithHTTPClient(client *http.Client) *RepairClusterParams {
	var ()
	return &RepairClusterParams{
		HTTPClient: client,
	}
}

/*RepairClusterParams contains all the parameters to send to the API endpoint
for the repair cluster operation typically these are written to a http.Request
*/
type RepairClusterParams struct {

	/*Body*/
	Body *model.ClusterRepairRequest
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repair cluster params
func (o *RepairClusterParams) WithTimeout(timeout time.Duration) *RepairClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repair cluster params
func (o *RepairClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repair cluster params
func (o *RepairClusterParams) WithContext(ctx context.Context) *RepairClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repair cluster params
func (o *RepairClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repair cluster params
func (o *RepairClusterParams) WithHTTPClient(client *http.Client) *RepairClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repair cluster params
func (o *RepairClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repair cluster params
func (o *RepairClusterParams) WithBody(body *model.ClusterRepairRequest) *RepairClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repair cluster params
func (o *RepairClusterParams) SetBody(body *model.ClusterRepairRequest) {
	o.Body = body
}

// WithID adds the id to the repair cluster params
func (o *RepairClusterParams) WithID(id int64) *RepairClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repair cluster params
func (o *RepairClusterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RepairClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

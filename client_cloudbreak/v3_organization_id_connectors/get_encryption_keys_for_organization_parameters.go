// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_connectors

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewGetEncryptionKeysForOrganizationParams creates a new GetEncryptionKeysForOrganizationParams object
// with the default values initialized.
func NewGetEncryptionKeysForOrganizationParams() *GetEncryptionKeysForOrganizationParams {
	var ()
	return &GetEncryptionKeysForOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEncryptionKeysForOrganizationParamsWithTimeout creates a new GetEncryptionKeysForOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEncryptionKeysForOrganizationParamsWithTimeout(timeout time.Duration) *GetEncryptionKeysForOrganizationParams {
	var ()
	return &GetEncryptionKeysForOrganizationParams{

		timeout: timeout,
	}
}

// NewGetEncryptionKeysForOrganizationParamsWithContext creates a new GetEncryptionKeysForOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEncryptionKeysForOrganizationParamsWithContext(ctx context.Context) *GetEncryptionKeysForOrganizationParams {
	var ()
	return &GetEncryptionKeysForOrganizationParams{

		Context: ctx,
	}
}

// NewGetEncryptionKeysForOrganizationParamsWithHTTPClient creates a new GetEncryptionKeysForOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEncryptionKeysForOrganizationParamsWithHTTPClient(client *http.Client) *GetEncryptionKeysForOrganizationParams {
	var ()
	return &GetEncryptionKeysForOrganizationParams{
		HTTPClient: client,
	}
}

/*GetEncryptionKeysForOrganizationParams contains all the parameters to send to the API endpoint
for the get encryption keys for organization operation typically these are written to a http.Request
*/
type GetEncryptionKeysForOrganizationParams struct {

	/*Body*/
	Body *models_cloudbreak.PlatformResourceRequestJSON
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) WithTimeout(timeout time.Duration) *GetEncryptionKeysForOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) WithContext(ctx context.Context) *GetEncryptionKeysForOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) WithHTTPClient(client *http.Client) *GetEncryptionKeysForOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) WithBody(body *models_cloudbreak.PlatformResourceRequestJSON) *GetEncryptionKeysForOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) SetBody(body *models_cloudbreak.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) WithOrganizationID(organizationID int64) *GetEncryptionKeysForOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get encryption keys for organization params
func (o *GetEncryptionKeysForOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEncryptionKeysForOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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
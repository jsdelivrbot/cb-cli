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

// NewCreateRecommendationForOrganizationParams creates a new CreateRecommendationForOrganizationParams object
// with the default values initialized.
func NewCreateRecommendationForOrganizationParams() *CreateRecommendationForOrganizationParams {
	var ()
	return &CreateRecommendationForOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRecommendationForOrganizationParamsWithTimeout creates a new CreateRecommendationForOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRecommendationForOrganizationParamsWithTimeout(timeout time.Duration) *CreateRecommendationForOrganizationParams {
	var ()
	return &CreateRecommendationForOrganizationParams{

		timeout: timeout,
	}
}

// NewCreateRecommendationForOrganizationParamsWithContext creates a new CreateRecommendationForOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRecommendationForOrganizationParamsWithContext(ctx context.Context) *CreateRecommendationForOrganizationParams {
	var ()
	return &CreateRecommendationForOrganizationParams{

		Context: ctx,
	}
}

// NewCreateRecommendationForOrganizationParamsWithHTTPClient creates a new CreateRecommendationForOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRecommendationForOrganizationParamsWithHTTPClient(client *http.Client) *CreateRecommendationForOrganizationParams {
	var ()
	return &CreateRecommendationForOrganizationParams{
		HTTPClient: client,
	}
}

/*CreateRecommendationForOrganizationParams contains all the parameters to send to the API endpoint
for the create recommendation for organization operation typically these are written to a http.Request
*/
type CreateRecommendationForOrganizationParams struct {

	/*Body*/
	Body *models_cloudbreak.RecommendationRequestJSON
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) WithTimeout(timeout time.Duration) *CreateRecommendationForOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) WithContext(ctx context.Context) *CreateRecommendationForOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) WithHTTPClient(client *http.Client) *CreateRecommendationForOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) WithBody(body *models_cloudbreak.RecommendationRequestJSON) *CreateRecommendationForOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) SetBody(body *models_cloudbreak.RecommendationRequestJSON) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) WithOrganizationID(organizationID int64) *CreateRecommendationForOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create recommendation for organization params
func (o *CreateRecommendationForOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRecommendationForOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.RecommendationRequestJSON)
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
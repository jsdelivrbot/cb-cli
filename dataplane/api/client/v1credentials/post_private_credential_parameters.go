// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPostPrivateCredentialParams creates a new PostPrivateCredentialParams object
// with the default values initialized.
func NewPostPrivateCredentialParams() *PostPrivateCredentialParams {
	var ()
	return &PostPrivateCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateCredentialParamsWithTimeout creates a new PostPrivateCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateCredentialParamsWithTimeout(timeout time.Duration) *PostPrivateCredentialParams {
	var ()
	return &PostPrivateCredentialParams{

		timeout: timeout,
	}
}

// NewPostPrivateCredentialParamsWithContext creates a new PostPrivateCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateCredentialParamsWithContext(ctx context.Context) *PostPrivateCredentialParams {
	var ()
	return &PostPrivateCredentialParams{

		Context: ctx,
	}
}

// NewPostPrivateCredentialParamsWithHTTPClient creates a new PostPrivateCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateCredentialParamsWithHTTPClient(client *http.Client) *PostPrivateCredentialParams {
	var ()
	return &PostPrivateCredentialParams{
		HTTPClient: client,
	}
}

/*PostPrivateCredentialParams contains all the parameters to send to the API endpoint
for the post private credential operation typically these are written to a http.Request
*/
type PostPrivateCredentialParams struct {

	/*Body*/
	Body *model.CredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private credential params
func (o *PostPrivateCredentialParams) WithTimeout(timeout time.Duration) *PostPrivateCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private credential params
func (o *PostPrivateCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private credential params
func (o *PostPrivateCredentialParams) WithContext(ctx context.Context) *PostPrivateCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private credential params
func (o *PostPrivateCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private credential params
func (o *PostPrivateCredentialParams) WithHTTPClient(client *http.Client) *PostPrivateCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private credential params
func (o *PostPrivateCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private credential params
func (o *PostPrivateCredentialParams) WithBody(body *model.CredentialRequest) *PostPrivateCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private credential params
func (o *PostPrivateCredentialParams) SetBody(body *model.CredentialRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

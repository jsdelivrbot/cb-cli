package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPrivateTemplateParams creates a new GetPrivateTemplateParams object
// with the default values initialized.
func NewGetPrivateTemplateParams() *GetPrivateTemplateParams {
	var ()
	return &GetPrivateTemplateParams{}
}

/*GetPrivateTemplateParams contains all the parameters to send to the API endpoint
for the get private template operation typically these are written to a http.Request
*/
type GetPrivateTemplateParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the get private template params
func (o *GetPrivateTemplateParams) WithName(name string) *GetPrivateTemplateParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateTemplateParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
package sssd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPrivateSssdParams creates a new GetPrivateSssdParams object
// with the default values initialized.
func NewGetPrivateSssdParams() *GetPrivateSssdParams {
	var ()
	return &GetPrivateSssdParams{}
}

/*GetPrivateSssdParams contains all the parameters to send to the API endpoint
for the get private sssd operation typically these are written to a http.Request
*/
type GetPrivateSssdParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the get private sssd params
func (o *GetPrivateSssdParams) WithName(name string) *GetPrivateSssdParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSssdParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
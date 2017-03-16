package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetClusterParams creates a new GetClusterParams object
// with the default values initialized.
func NewGetClusterParams() *GetClusterParams {
	var ()
	return &GetClusterParams{}
}

/*GetClusterParams contains all the parameters to send to the API endpoint
for the get cluster operation typically these are written to a http.Request
*/
type GetClusterParams struct {

	/*ClusterID*/
	ClusterID int64
}

// WithClusterID adds the clusterId to the get cluster params
func (o *GetClusterParams) WithClusterID(clusterId int64) *GetClusterParams {
	o.ClusterID = clusterId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
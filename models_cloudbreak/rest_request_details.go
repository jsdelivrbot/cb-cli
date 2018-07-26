// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RestRequestDetails rest request details
// swagger:model RestRequestDetails

type RestRequestDetails struct {

	// body
	Body string `json:"body,omitempty"`

	// cookies
	Cookies map[string]string `json:"cookies,omitempty"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// media type
	MediaType string `json:"mediaType,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// request Uri
	RequestURI string `json:"requestUri,omitempty"`
}

/* polymorph RestRequestDetails body false */

/* polymorph RestRequestDetails cookies false */

/* polymorph RestRequestDetails headers false */

/* polymorph RestRequestDetails mediaType false */

/* polymorph RestRequestDetails method false */

/* polymorph RestRequestDetails requestUri false */

// Validate validates this rest request details
func (m *RestRequestDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RestRequestDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestRequestDetails) UnmarshalBinary(b []byte) error {
	var res RestRequestDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

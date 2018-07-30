// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ParametersQueryResponse parameters query response
// swagger:model ParametersQueryResponse

type ParametersQueryResponse struct {

	// Custom parameters as a json
	// Required: true
	Custom map[string]string `json:"custom"`
}

/* polymorph ParametersQueryResponse custom false */

// Validate validates this parameters query response
func (m *ParametersQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParametersQueryResponse) validateCustom(formats strfmt.Registry) error {

	if swag.IsZero(m.Custom) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParametersQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParametersQueryResponse) UnmarshalBinary(b []byte) error {
	var res ParametersQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
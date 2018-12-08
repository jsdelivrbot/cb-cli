// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackMatrix stack matrix
// swagger:model StackMatrix
type StackMatrix struct {

	// hdf
	Hdf map[string]StackDescriptor `json:"hdf,omitempty"`

	// hdp
	Hdp map[string]StackDescriptor `json:"hdp,omitempty"`
}

// Validate validates this stack matrix
func (m *StackMatrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHdf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackMatrix) validateHdf(formats strfmt.Registry) error {

	if swag.IsZero(m.Hdf) { // not required
		return nil
	}

	for k := range m.Hdf {

		if err := validate.Required("hdf"+"."+k, "body", m.Hdf[k]); err != nil {
			return err
		}
		if val, ok := m.Hdf[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *StackMatrix) validateHdp(formats strfmt.Registry) error {

	if swag.IsZero(m.Hdp) { // not required
		return nil
	}

	for k := range m.Hdp {

		if err := validate.Required("hdp"+"."+k, "body", m.Hdp[k]); err != nil {
			return err
		}
		if val, ok := m.Hdp[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackMatrix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackMatrix) UnmarshalBinary(b []byte) error {
	var res StackMatrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
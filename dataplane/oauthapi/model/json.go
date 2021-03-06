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

// JSON Json
// swagger:model Json

type JSON struct {

	// array
	// Required: true
	Array *bool `json:"array"`

	// boolean
	// Required: true
	Boolean *bool `json:"boolean"`

	// null
	// Required: true
	Null *bool `json:"null"`

	// number
	// Required: true
	Number *bool `json:"number"`

	// object
	// Required: true
	Object *bool `json:"object"`

	// string
	// Required: true
	String *bool `json:"string"`
}

/* polymorph Json array false */

/* polymorph Json boolean false */

/* polymorph Json null false */

/* polymorph Json number false */

/* polymorph Json object false */

/* polymorph Json string false */

// Validate validates this Json
func (m *JSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArray(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBoolean(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNull(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateString(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSON) validateArray(formats strfmt.Registry) error {

	if err := validate.Required("array", "body", m.Array); err != nil {
		return err
	}

	return nil
}

func (m *JSON) validateBoolean(formats strfmt.Registry) error {

	if err := validate.Required("boolean", "body", m.Boolean); err != nil {
		return err
	}

	return nil
}

func (m *JSON) validateNull(formats strfmt.Registry) error {

	if err := validate.Required("null", "body", m.Null); err != nil {
		return err
	}

	return nil
}

func (m *JSON) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *JSON) validateObject(formats strfmt.Registry) error {

	if err := validate.Required("object", "body", m.Object); err != nil {
		return err
	}

	return nil
}

func (m *JSON) validateString(formats strfmt.Registry) error {

	if err := validate.Required("string", "body", m.String); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSON) UnmarshalBinary(b []byte) error {
	var res JSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BlueprintInput blueprint input
// swagger:model BlueprintInput
type BlueprintInput struct {

	// name
	Name string `json:"name,omitempty"`

	// property value
	PropertyValue string `json:"propertyValue,omitempty"`
}

// Validate validates this blueprint input
func (m *BlueprintInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintInput) UnmarshalBinary(b []byte) error {
	var res BlueprintInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

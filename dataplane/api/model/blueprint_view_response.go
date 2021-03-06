// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BlueprintViewResponse blueprint view response
// swagger:model BlueprintViewResponse
type BlueprintViewResponse struct {

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// number of host groups
	HostGroupCount int32 `json:"hostGroupCount,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[^;\/%]*$
	Name *string `json:"name"`

	// The type of the stack: for example HDP or HDF
	StackType string `json:"stackType,omitempty"`

	// The version of the stack
	StackVersion string `json:"stackVersion,omitempty"`

	// status of the blueprint
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED]
	Status string `json:"status,omitempty"`

	// user defined tags for blueprint
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// Validate validates this blueprint view response
func (m *BlueprintViewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintViewResponse) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintViewResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[^;\/%]*$`); err != nil {
		return err
	}

	return nil
}

var blueprintViewResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blueprintViewResponseTypeStatusPropEnum = append(blueprintViewResponseTypeStatusPropEnum, v)
	}
}

const (

	// BlueprintViewResponseStatusDEFAULT captures enum value "DEFAULT"
	BlueprintViewResponseStatusDEFAULT string = "DEFAULT"

	// BlueprintViewResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	BlueprintViewResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// BlueprintViewResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	BlueprintViewResponseStatusUSERMANAGED string = "USER_MANAGED"
)

// prop value enum
func (m *BlueprintViewResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, blueprintViewResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BlueprintViewResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintViewResponse) UnmarshalBinary(b []byte) error {
	var res BlueprintViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

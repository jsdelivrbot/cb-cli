// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceGroupResponse instance group response
// swagger:model InstanceGroupResponse
type InstanceGroupResponse struct {

	// name of the instance group
	// Required: true
	Group *string `json:"group"`

	// id of the resource
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// metadata of instances
	// Read Only: true
	// Unique: true
	Metadata []*InstanceMetaData `json:"metadata"`

	// number of nodes
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	NodeCount *int32 `json:"nodeCount"`

	// cloud specific parameters for instance group
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// instancegroup related securitygroup
	SecurityGroup *SecurityGroupResponse `json:"securityGroup,omitempty"`

	// security group resource id for the instance group
	SecurityGroupID int64 `json:"securityGroupId,omitempty"`

	// instancegroup related template
	Template *TemplateResponse `json:"template,omitempty"`

	// referenced template id
	TemplateID int64 `json:"templateId,omitempty"`

	// type of the instance group
	// Enum: [GATEWAY CORE]
	Type string `json:"type,omitempty"`
}

// Validate validates this instance group response
func (m *InstanceGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupResponse) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := validate.UniqueItems("metadata", "body", m.Metadata); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroupResponse) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("nodeCount", "body", int64(*m.NodeCount), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nodeCount", "body", int64(*m.NodeCount), 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupResponse) validateSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupResponse) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

var instanceGroupResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GATEWAY","CORE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupResponseTypeTypePropEnum = append(instanceGroupResponseTypeTypePropEnum, v)
	}
}

const (

	// InstanceGroupResponseTypeGATEWAY captures enum value "GATEWAY"
	InstanceGroupResponseTypeGATEWAY string = "GATEWAY"

	// InstanceGroupResponseTypeCORE captures enum value "CORE"
	InstanceGroupResponseTypeCORE string = "CORE"
)

// prop value enum
func (m *InstanceGroupResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupResponse) UnmarshalBinary(b []byte) error {
	var res InstanceGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

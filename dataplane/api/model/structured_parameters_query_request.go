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

// StructuredParametersQueryRequest structured parameters query request
// swagger:model StructuredParametersQueryRequest
type StructuredParametersQueryRequest struct {

	// Account name of the path
	AccountName string `json:"accountName,omitempty"`

	// Attached cluster
	// Required: true
	AttachedCluster bool `json:"attachedCluster"`

	// gathered from blueprintName field from the blueprint JSON
	// Required: true
	BlueprintName *string `json:"blueprintName"`

	// name of the stack
	// Required: true
	ClusterName *string `json:"clusterName"`

	// Type of filesystem
	// Required: true
	FileSystemType *string `json:"fileSystemType"`

	// Storage name of the path
	// Required: true
	StorageName *string `json:"storageName"`
}

// Validate validates this structured parameters query request
func (m *StructuredParametersQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlueprintName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileSystemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StructuredParametersQueryRequest) validateAttachedCluster(formats strfmt.Registry) error {

	if err := validate.Required("attachedCluster", "body", bool(m.AttachedCluster)); err != nil {
		return err
	}

	return nil
}

func (m *StructuredParametersQueryRequest) validateBlueprintName(formats strfmt.Registry) error {

	if err := validate.Required("blueprintName", "body", m.BlueprintName); err != nil {
		return err
	}

	return nil
}

func (m *StructuredParametersQueryRequest) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *StructuredParametersQueryRequest) validateFileSystemType(formats strfmt.Registry) error {

	if err := validate.Required("fileSystemType", "body", m.FileSystemType); err != nil {
		return err
	}

	return nil
}

func (m *StructuredParametersQueryRequest) validateStorageName(formats strfmt.Registry) error {

	if err := validate.Required("storageName", "body", m.StorageName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StructuredParametersQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StructuredParametersQueryRequest) UnmarshalBinary(b []byte) error {
	var res StructuredParametersQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
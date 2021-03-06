// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageResponse image response
// swagger:model ImageResponse
type ImageResponse struct {

	// date
	Date string `json:"date,omitempty"`

	// default image
	DefaultImage *bool `json:"defaultImage,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// images
	Images map[string]map[string]string `json:"images,omitempty"`

	// os
	Os string `json:"os,omitempty"`

	// os type
	OsType string `json:"osType,omitempty"`

	// package versions
	PackageVersions map[string]string `json:"packageVersions,omitempty"`

	// repo
	Repo map[string]string `json:"repo,omitempty"`

	// stack details
	StackDetails *StackDetailsJSON `json:"stackDetails,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this image response
func (m *ImageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStackDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageResponse) validateStackDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.StackDetails) { // not required
		return nil
	}

	if m.StackDetails != nil {
		if err := m.StackDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageResponse) UnmarshalBinary(b []byte) error {
	var res ImageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

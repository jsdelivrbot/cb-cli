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

// RegionResponse region response
// swagger:model RegionResponse
type RegionResponse struct {

	// availability zones
	AvailabilityZones map[string][]string `json:"availabilityZones,omitempty"`

	// default regions
	DefaultRegion string `json:"defaultRegion,omitempty"`

	// regions with displayNames
	DisplayNames map[string]string `json:"displayNames,omitempty"`

	// regions with location data
	// Unique: true
	Locations []string `json:"locations"`

	// regions
	// Unique: true
	Regions []string `json:"regions"`
}

// Validate validates this region response
func (m *RegionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegionResponse) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	if err := validate.UniqueItems("locations", "body", m.Locations); err != nil {
		return err
	}

	return nil
}

func (m *RegionResponse) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegionResponse) UnmarshalBinary(b []byte) error {
	var res RegionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

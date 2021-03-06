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

// ImageCatalogResponse image catalog response
// swagger:model ImageCatalogResponse
type ImageCatalogResponse struct {

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// id of the resource
	// Required: true
	ID *int64 `json:"id"`

	// image response in imagecatalog
	ImagesResponse *ImagesResponse `json:"imagesResponse,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// resource is visible in account
	// Required: true
	PublicInAccount bool `json:"publicInAccount"`

	// custom image catalog's URL
	// Required: true
	// Pattern: ^http[s]?://.*
	URL *string `json:"url"`

	// true if image catalog is the default one
	// Required: true
	UsedAsDefault bool `json:"usedAsDefault"`

	// workspace of the resource
	Workspace *WorkspaceResourceResponse `json:"workspace,omitempty"`
}

// Validate validates this image catalog response
func (m *ImageCatalogResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImagesResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicInAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedAsDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageCatalogResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *ImageCatalogResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ImageCatalogResponse) validateImagesResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.ImagesResponse) { // not required
		return nil
	}

	if m.ImagesResponse != nil {
		if err := m.ImagesResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imagesResponse")
			}
			return err
		}
	}

	return nil
}

func (m *ImageCatalogResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *ImageCatalogResponse) validatePublicInAccount(formats strfmt.Registry) error {

	if err := validate.Required("publicInAccount", "body", bool(m.PublicInAccount)); err != nil {
		return err
	}

	return nil
}

func (m *ImageCatalogResponse) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", string(*m.URL), `^http[s]?://.*`); err != nil {
		return err
	}

	return nil
}

func (m *ImageCatalogResponse) validateUsedAsDefault(formats strfmt.Registry) error {

	if err := validate.Required("usedAsDefault", "body", bool(m.UsedAsDefault)); err != nil {
		return err
	}

	return nil
}

func (m *ImageCatalogResponse) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageCatalogResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageCatalogResponse) UnmarshalBinary(b []byte) error {
	var res ImageCatalogResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ClusterTemplateResponse cluster template response

swagger:model ClusterTemplateResponse
*/
type ClusterTemplateResponse struct {

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* name of the cluster template

	Required: true
	*/
	Name string `json:"name"`

	/* stringified template JSON
	 */
	Template *string `json:"template,omitempty"`

	/* type of the cluster template
	 */
	Type *string `json:"type,omitempty"`
}

// Validate validates this cluster template response
func (m *ClusterTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTemplateResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

var clusterTemplateResponseTypeTypePropEnum []interface{}

func (m *ClusterTemplateResponse) validateTypeEnum(path, location string, value string) error {
	if clusterTemplateResponseTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["QUICK_START"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			clusterTemplateResponseTypeTypePropEnum = append(clusterTemplateResponseTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, clusterTemplateResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}
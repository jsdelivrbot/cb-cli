package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Greeting greeting

swagger:model Greeting
*/
type Greeting struct {

	/* content
	 */
	Content *string `json:"content,omitempty"`

	/* id
	 */
	ID *int64 `json:"id,omitempty"`
}

// Validate validates this greeting
func (m *Greeting) Validate(formats strfmt.Registry) error {
	return nil
}

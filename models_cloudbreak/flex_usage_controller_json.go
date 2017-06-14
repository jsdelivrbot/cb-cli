package models_cloudbreak

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*FlexUsageControllerJSON flex usage controller json

swagger:model FlexUsageControllerJson
*/
type FlexUsageControllerJSON struct {

	/* guid
	 */
	GUID *string `json:"guid,omitempty"`

	/* instance id
	 */
	InstanceID *string `json:"instanceId,omitempty"`

	/* provider
	 */
	Provider *string `json:"provider,omitempty"`

	/* region
	 */
	Region *string `json:"region,omitempty"`

	/* smart sense id
	 */
	SmartSenseID *string `json:"smartSenseId,omitempty"`

	/* user name
	 */
	UserName *string `json:"userName,omitempty"`
}

// Validate validates this flex usage controller json
func (m *FlexUsageControllerJSON) Validate(formats strfmt.Registry) error {
	return nil
}
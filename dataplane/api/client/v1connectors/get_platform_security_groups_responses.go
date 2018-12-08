// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPlatformSecurityGroupsReader is a Reader for the GetPlatformSecurityGroups structure.
type GetPlatformSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlatformSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlatformSecurityGroupsOK creates a GetPlatformSecurityGroupsOK with default headers values
func NewGetPlatformSecurityGroupsOK() *GetPlatformSecurityGroupsOK {
	return &GetPlatformSecurityGroupsOK{}
}

/*GetPlatformSecurityGroupsOK handles this case with default header values.

successful operation
*/
type GetPlatformSecurityGroupsOK struct {
	Payload *model.PlatformSecurityGroupsResponse
}

func (o *GetPlatformSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[POST /v1/connectors/securitygroups][%d] getPlatformSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *GetPlatformSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformSecurityGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
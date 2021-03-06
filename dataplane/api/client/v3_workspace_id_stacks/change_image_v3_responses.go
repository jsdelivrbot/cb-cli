// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeImageV3Reader is a Reader for the ChangeImageV3 structure.
type ChangeImageV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeImageV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewChangeImageV3Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewChangeImageV3Default creates a ChangeImageV3Default with default headers values
func NewChangeImageV3Default(code int) *ChangeImageV3Default {
	return &ChangeImageV3Default{
		_statusCode: code,
	}
}

/*ChangeImageV3Default handles this case with default header values.

successful operation
*/
type ChangeImageV3Default struct {
	_statusCode int
}

// Code gets the status code for the change image v3 default response
func (o *ChangeImageV3Default) Code() int {
	return o._statusCode
}

func (o *ChangeImageV3Default) Error() string {
	return fmt.Sprintf("[PUT /v3/{workspaceId}/stacks/{name}/changeImage][%d] changeImageV3 default ", o._statusCode)
}

func (o *ChangeImageV3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

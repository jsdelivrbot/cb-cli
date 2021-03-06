// Code generated by go-swagger; DO NOT EDIT.

package v1mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePrivateManagementPackReader is a Reader for the DeletePrivateManagementPack structure.
type DeletePrivateManagementPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateManagementPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePrivateManagementPackDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePrivateManagementPackDefault creates a DeletePrivateManagementPackDefault with default headers values
func NewDeletePrivateManagementPackDefault(code int) *DeletePrivateManagementPackDefault {
	return &DeletePrivateManagementPackDefault{
		_statusCode: code,
	}
}

/*DeletePrivateManagementPackDefault handles this case with default header values.

successful operation
*/
type DeletePrivateManagementPackDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private management pack default response
func (o *DeletePrivateManagementPackDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateManagementPackDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/mpacks/user/{name}][%d] deletePrivateManagementPack default ", o._statusCode)
}

func (o *DeletePrivateManagementPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

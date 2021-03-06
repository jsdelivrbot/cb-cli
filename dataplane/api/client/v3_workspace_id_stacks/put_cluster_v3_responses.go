// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutClusterV3Reader is a Reader for the PutClusterV3 structure.
type PutClusterV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutClusterV3Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutClusterV3Default creates a PutClusterV3Default with default headers values
func NewPutClusterV3Default(code int) *PutClusterV3Default {
	return &PutClusterV3Default{
		_statusCode: code,
	}
}

/*PutClusterV3Default handles this case with default header values.

successful operation
*/
type PutClusterV3Default struct {
	_statusCode int
}

// Code gets the status code for the put cluster v3 default response
func (o *PutClusterV3Default) Code() int {
	return o._statusCode
}

func (o *PutClusterV3Default) Error() string {
	return fmt.Sprintf("[PUT /v3/{workspaceId}/stacks/{name}/cluster][%d] putClusterV3 default ", o._statusCode)
}

func (o *PutClusterV3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

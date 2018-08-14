// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutstopStackV3Reader is a Reader for the PutstopStackV3 structure.
type PutstopStackV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutstopStackV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutstopStackV3Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutstopStackV3Default creates a PutstopStackV3Default with default headers values
func NewPutstopStackV3Default(code int) *PutstopStackV3Default {
	return &PutstopStackV3Default{
		_statusCode: code,
	}
}

/*PutstopStackV3Default handles this case with default header values.

successful operation
*/
type PutstopStackV3Default struct {
	_statusCode int
}

// Code gets the status code for the putstop stack v3 default response
func (o *PutstopStackV3Default) Code() int {
	return o._statusCode
}

func (o *PutstopStackV3Default) Error() string {
	return fmt.Sprintf("[PUT /v3/{organizationId}/stack/{name}/stop][%d] putstopStackV3 default ", o._statusCode)
}

func (o *PutstopStackV3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

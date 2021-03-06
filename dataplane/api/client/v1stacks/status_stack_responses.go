// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StatusStackReader is a Reader for the StatusStack structure.
type StatusStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatusStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatusStackOK creates a StatusStackOK with default headers values
func NewStatusStackOK() *StatusStackOK {
	return &StatusStackOK{}
}

/*StatusStackOK handles this case with default header values.

successful operation
*/
type StatusStackOK struct {
	Payload map[string]interface{}
}

func (o *StatusStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/status][%d] statusStackOK  %+v", 200, o.Payload)
}

func (o *StatusStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

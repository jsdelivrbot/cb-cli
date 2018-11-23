// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetRegionAvByTypeReader is a Reader for the GetRegionAvByType structure.
type GetRegionAvByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionAvByTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionAvByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRegionAvByTypeOK creates a GetRegionAvByTypeOK with default headers values
func NewGetRegionAvByTypeOK() *GetRegionAvByTypeOK {
	return &GetRegionAvByTypeOK{}
}

/*GetRegionAvByTypeOK handles this case with default header values.

successful operation
*/
type GetRegionAvByTypeOK struct {
	Payload GetRegionAvByTypeOKBody
}

func (o *GetRegionAvByTypeOK) Error() string {
	return fmt.Sprintf("[GET /v1/connectors/regions/av/{type}][%d] getRegionAvByTypeOK  %+v", 200, o.Payload)
}

func (o *GetRegionAvByTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRegionAvByTypeOKBody get region av by type o k body
swagger:model GetRegionAvByTypeOKBody
*/

type GetRegionAvByTypeOKBody map[string][]string

// Validate validates this get region av by type o k body
func (o GetRegionAvByTypeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
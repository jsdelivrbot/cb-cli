package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetRegionAvByTypeReader is a Reader for the GetRegionAvByType structure.
type GetRegionAvByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetRegionAvByTypeReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionAvByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
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
	Payload GetRegionAvByTypeOKBodyBody
}

func (o *GetRegionAvByTypeOK) Error() string {
	return fmt.Sprintf("[GET /connectors/regions/av/{type}][%d] getRegionAvByTypeOK  %+v", 200, o.Payload)
}

func (o *GetRegionAvByTypeOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRegionAvByTypeOKBodyBody get region av by type o k body body

swagger:model GetRegionAvByTypeOKBodyBody
*/
type GetRegionAvByTypeOKBodyBody map[string][]string

// Validate validates this get region av by type o k body body
func (o GetRegionAvByTypeOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("getRegionAvByTypeOK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
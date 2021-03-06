// Code generated by go-swagger; DO NOT EDIT.

package v1imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPublicsImageCatalogsReader is a Reader for the GetPublicsImageCatalogs structure.
type GetPublicsImageCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsImageCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsImageCatalogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsImageCatalogsOK creates a GetPublicsImageCatalogsOK with default headers values
func NewGetPublicsImageCatalogsOK() *GetPublicsImageCatalogsOK {
	return &GetPublicsImageCatalogsOK{}
}

/*GetPublicsImageCatalogsOK handles this case with default header values.

successful operation
*/
type GetPublicsImageCatalogsOK struct {
	Payload []*model.ImageCatalogResponse
}

func (o *GetPublicsImageCatalogsOK) Error() string {
	return fmt.Sprintf("[GET /v1/imagecatalogs/account][%d] getPublicsImageCatalogsOK  %+v", 200, o.Payload)
}

func (o *GetPublicsImageCatalogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

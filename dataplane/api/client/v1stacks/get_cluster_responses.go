// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetClusterReader is a Reader for the GetCluster structure.
type GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterOK creates a GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

/*GetClusterOK handles this case with default header values.

successful operation
*/
type GetClusterOK struct {
	Payload *model.ClusterResponse
}

func (o *GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/cluster][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

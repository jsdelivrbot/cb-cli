// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PutSetDefaultImageCatalogByNameInWorkspaceReader is a Reader for the PutSetDefaultImageCatalogByNameInWorkspace structure.
type PutSetDefaultImageCatalogByNameInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSetDefaultImageCatalogByNameInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutSetDefaultImageCatalogByNameInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSetDefaultImageCatalogByNameInWorkspaceOK creates a PutSetDefaultImageCatalogByNameInWorkspaceOK with default headers values
func NewPutSetDefaultImageCatalogByNameInWorkspaceOK() *PutSetDefaultImageCatalogByNameInWorkspaceOK {
	return &PutSetDefaultImageCatalogByNameInWorkspaceOK{}
}

/*PutSetDefaultImageCatalogByNameInWorkspaceOK handles this case with default header values.

successful operation
*/
type PutSetDefaultImageCatalogByNameInWorkspaceOK struct {
	Payload *model.ImageCatalogResponse
}

func (o *PutSetDefaultImageCatalogByNameInWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /v3/{workspaceId}/imagecatalogs/setdefault/{name}][%d] putSetDefaultImageCatalogByNameInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutSetDefaultImageCatalogByNameInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ImageCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

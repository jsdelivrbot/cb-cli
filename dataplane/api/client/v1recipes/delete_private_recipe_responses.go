// Code generated by go-swagger; DO NOT EDIT.

package v1recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePrivateRecipeReader is a Reader for the DeletePrivateRecipe structure.
type DeletePrivateRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePrivateRecipeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePrivateRecipeDefault creates a DeletePrivateRecipeDefault with default headers values
func NewDeletePrivateRecipeDefault(code int) *DeletePrivateRecipeDefault {
	return &DeletePrivateRecipeDefault{
		_statusCode: code,
	}
}

/*DeletePrivateRecipeDefault handles this case with default header values.

successful operation
*/
type DeletePrivateRecipeDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private recipe default response
func (o *DeletePrivateRecipeDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateRecipeDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/recipes/user/{name}][%d] deletePrivateRecipe default ", o._statusCode)
}

func (o *DeletePrivateRecipeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

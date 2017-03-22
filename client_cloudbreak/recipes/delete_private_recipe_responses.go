package recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeletePrivateRecipeReader is a Reader for the DeletePrivateRecipe structure.
type DeletePrivateRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeletePrivateRecipeReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeletePrivateRecipeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
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
	return fmt.Sprintf("[DELETE /recipes/user/{name}][%d] deletePrivateRecipe default ", o._statusCode)
}

func (o *DeletePrivateRecipeDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
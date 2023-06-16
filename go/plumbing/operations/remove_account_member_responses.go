// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// RemoveAccountMemberReader is a Reader for the RemoveAccountMember structure.
type RemoveAccountMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAccountMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveAccountMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveAccountMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveAccountMemberNoContent creates a RemoveAccountMemberNoContent with default headers values
func NewRemoveAccountMemberNoContent() *RemoveAccountMemberNoContent {
	return &RemoveAccountMemberNoContent{}
}

/*RemoveAccountMemberNoContent handles this case with default header values.

Not Content
*/
type RemoveAccountMemberNoContent struct {
}

func (o *RemoveAccountMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{account_slug}/members/{member_id}][%d] removeAccountMemberNoContent ", 204)
}

func (o *RemoveAccountMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccountMemberDefault creates a RemoveAccountMemberDefault with default headers values
func NewRemoveAccountMemberDefault(code int) *RemoveAccountMemberDefault {
	return &RemoveAccountMemberDefault{
		_statusCode: code,
	}
}

/*RemoveAccountMemberDefault handles this case with default header values.

error
*/
type RemoveAccountMemberDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the remove account member default response
func (o *RemoveAccountMemberDefault) Code() int {
	return o._statusCode
}

func (o *RemoveAccountMemberDefault) Error() string {
	return fmt.Sprintf("[DELETE /{account_slug}/members/{member_id}][%d] removeAccountMember default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveAccountMemberDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveAccountMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// DeleteDNSRecordReader is a Reader for the DeleteDNSRecord structure.
type DeleteDNSRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDNSRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDNSRecordNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDNSRecordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDNSRecordNoContent creates a DeleteDNSRecordNoContent with default headers values
func NewDeleteDNSRecordNoContent() *DeleteDNSRecordNoContent {
	return &DeleteDNSRecordNoContent{}
}

/*DeleteDNSRecordNoContent handles this case with default header values.

record deleted
*/
type DeleteDNSRecordNoContent struct {
}

func (o *DeleteDNSRecordNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dns_zones/{zone_id}/dns_records/{dns_record_id}][%d] deleteDnsRecordNoContent ", 204)
}

func (o *DeleteDNSRecordNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDNSRecordDefault creates a DeleteDNSRecordDefault with default headers values
func NewDeleteDNSRecordDefault(code int) *DeleteDNSRecordDefault {
	return &DeleteDNSRecordDefault{
		_statusCode: code,
	}
}

/*DeleteDNSRecordDefault handles this case with default header values.

error
*/
type DeleteDNSRecordDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete Dns record default response
func (o *DeleteDNSRecordDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDNSRecordDefault) Error() string {
	return fmt.Sprintf("[DELETE /dns_zones/{zone_id}/dns_records/{dns_record_id}][%d] deleteDnsRecord default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDNSRecordDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDNSRecordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetEnvVarsReader is a Reader for the GetEnvVars structure.
type GetEnvVarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvVarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvVarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEnvVarsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEnvVarsOK creates a GetEnvVarsOK with default headers values
func NewGetEnvVarsOK() *GetEnvVarsOK {
	return &GetEnvVarsOK{}
}

/*GetEnvVarsOK handles this case with default header values.

OK
*/
type GetEnvVarsOK struct {
	Payload []*models.EnvVar
}

func (o *GetEnvVarsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/env][%d] getEnvVarsOK  %+v", 200, o.Payload)
}

func (o *GetEnvVarsOK) GetPayload() []*models.EnvVar {
	return o.Payload
}

func (o *GetEnvVarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvVarsDefault creates a GetEnvVarsDefault with default headers values
func NewGetEnvVarsDefault(code int) *GetEnvVarsDefault {
	return &GetEnvVarsDefault{
		_statusCode: code,
	}
}

/*GetEnvVarsDefault handles this case with default header values.

error
*/
type GetEnvVarsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get env vars default response
func (o *GetEnvVarsDefault) Code() int {
	return o._statusCode
}

func (o *GetEnvVarsDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/env][%d] getEnvVars default  %+v", o._statusCode, o.Payload)
}

func (o *GetEnvVarsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEnvVarsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

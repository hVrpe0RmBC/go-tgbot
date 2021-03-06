// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendVoiceReader is a Reader for the SendVoice structure.
type SendVoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendVoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendVoiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendVoiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendVoiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendVoiceEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendVoiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVoiceOK creates a SendVoiceOK with default headers values
func NewSendVoiceOK() *SendVoiceOK {
	return &SendVoiceOK{}
}

/*SendVoiceOK handles this case with default header values.

SendVoiceOK send voice o k
*/
type SendVoiceOK struct {
	Payload *models.ResponseMessage
}

func (o *SendVoiceOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceOK  %+v", 200, o.Payload)
}

func (o *SendVoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVoiceBadRequest creates a SendVoiceBadRequest with default headers values
func NewSendVoiceBadRequest() *SendVoiceBadRequest {
	return &SendVoiceBadRequest{}
}

/*SendVoiceBadRequest handles this case with default header values.

Bad Request
*/
type SendVoiceBadRequest struct {
	Payload *models.Error
}

func (o *SendVoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceBadRequest  %+v", 400, o.Payload)
}

func (o *SendVoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVoiceUnauthorized creates a SendVoiceUnauthorized with default headers values
func NewSendVoiceUnauthorized() *SendVoiceUnauthorized {
	return &SendVoiceUnauthorized{}
}

/*SendVoiceUnauthorized handles this case with default header values.

Unauthorized
*/
type SendVoiceUnauthorized struct {
	Payload *models.Error
}

func (o *SendVoiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceUnauthorized  %+v", 401, o.Payload)
}

func (o *SendVoiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVoiceForbidden creates a SendVoiceForbidden with default headers values
func NewSendVoiceForbidden() *SendVoiceForbidden {
	return &SendVoiceForbidden{}
}

/*SendVoiceForbidden handles this case with default header values.

Forbidden
*/
type SendVoiceForbidden struct {
	Payload *models.Error
}

func (o *SendVoiceForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceForbidden  %+v", 403, o.Payload)
}

func (o *SendVoiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVoiceNotFound creates a SendVoiceNotFound with default headers values
func NewSendVoiceNotFound() *SendVoiceNotFound {
	return &SendVoiceNotFound{}
}

/*SendVoiceNotFound handles this case with default header values.

Not Found
*/
type SendVoiceNotFound struct {
	Payload *models.Error
}

func (o *SendVoiceNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceNotFound  %+v", 404, o.Payload)
}

func (o *SendVoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVoiceEnhanceYourCalm creates a SendVoiceEnhanceYourCalm with default headers values
func NewSendVoiceEnhanceYourCalm() *SendVoiceEnhanceYourCalm {
	return &SendVoiceEnhanceYourCalm{}
}

/*SendVoiceEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendVoiceEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendVoiceEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendVoiceEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVoiceInternalServerError creates a SendVoiceInternalServerError with default headers values
func NewSendVoiceInternalServerError() *SendVoiceInternalServerError {
	return &SendVoiceInternalServerError{}
}

/*SendVoiceInternalServerError handles this case with default header values.

Internal
*/
type SendVoiceInternalServerError struct {
	Payload *models.Error
}

func (o *SendVoiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceInternalServerError  %+v", 500, o.Payload)
}

func (o *SendVoiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

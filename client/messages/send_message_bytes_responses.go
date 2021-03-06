// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendMessageBytesReader is a Reader for the SendMessageBytes structure.
type SendMessageBytesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendMessageBytesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendMessageBytesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendMessageBytesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendMessageBytesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendMessageBytesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendMessageBytesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendMessageBytesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendMessageBytesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendMessageBytesOK creates a SendMessageBytesOK with default headers values
func NewSendMessageBytesOK() *SendMessageBytesOK {
	return &SendMessageBytesOK{}
}

/*SendMessageBytesOK handles this case with default header values.

SendMessageBytesOK send message bytes o k
*/
type SendMessageBytesOK struct {
	Payload *models.ResponseMessage
}

func (o *SendMessageBytesOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesOK  %+v", 200, o.Payload)
}

func (o *SendMessageBytesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBytesBadRequest creates a SendMessageBytesBadRequest with default headers values
func NewSendMessageBytesBadRequest() *SendMessageBytesBadRequest {
	return &SendMessageBytesBadRequest{}
}

/*SendMessageBytesBadRequest handles this case with default header values.

Bad Request
*/
type SendMessageBytesBadRequest struct {
	Payload *models.Error
}

func (o *SendMessageBytesBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesBadRequest  %+v", 400, o.Payload)
}

func (o *SendMessageBytesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBytesUnauthorized creates a SendMessageBytesUnauthorized with default headers values
func NewSendMessageBytesUnauthorized() *SendMessageBytesUnauthorized {
	return &SendMessageBytesUnauthorized{}
}

/*SendMessageBytesUnauthorized handles this case with default header values.

Unauthorized
*/
type SendMessageBytesUnauthorized struct {
	Payload *models.Error
}

func (o *SendMessageBytesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesUnauthorized  %+v", 401, o.Payload)
}

func (o *SendMessageBytesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBytesForbidden creates a SendMessageBytesForbidden with default headers values
func NewSendMessageBytesForbidden() *SendMessageBytesForbidden {
	return &SendMessageBytesForbidden{}
}

/*SendMessageBytesForbidden handles this case with default header values.

Forbidden
*/
type SendMessageBytesForbidden struct {
	Payload *models.Error
}

func (o *SendMessageBytesForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesForbidden  %+v", 403, o.Payload)
}

func (o *SendMessageBytesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBytesNotFound creates a SendMessageBytesNotFound with default headers values
func NewSendMessageBytesNotFound() *SendMessageBytesNotFound {
	return &SendMessageBytesNotFound{}
}

/*SendMessageBytesNotFound handles this case with default header values.

Not Found
*/
type SendMessageBytesNotFound struct {
	Payload *models.Error
}

func (o *SendMessageBytesNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesNotFound  %+v", 404, o.Payload)
}

func (o *SendMessageBytesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBytesEnhanceYourCalm creates a SendMessageBytesEnhanceYourCalm with default headers values
func NewSendMessageBytesEnhanceYourCalm() *SendMessageBytesEnhanceYourCalm {
	return &SendMessageBytesEnhanceYourCalm{}
}

/*SendMessageBytesEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendMessageBytesEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendMessageBytesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendMessageBytesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBytesInternalServerError creates a SendMessageBytesInternalServerError with default headers values
func NewSendMessageBytesInternalServerError() *SendMessageBytesInternalServerError {
	return &SendMessageBytesInternalServerError{}
}

/*SendMessageBytesInternalServerError handles this case with default header values.

Internal
*/
type SendMessageBytesInternalServerError struct {
	Payload *models.Error
}

func (o *SendMessageBytesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendMessage#bytes][%d] sendMessageBytesInternalServerError  %+v", 500, o.Payload)
}

func (o *SendMessageBytesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

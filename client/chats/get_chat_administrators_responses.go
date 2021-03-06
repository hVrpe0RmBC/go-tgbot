// Code generated by go-swagger; DO NOT EDIT.

package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// GetChatAdministratorsReader is a Reader for the GetChatAdministrators structure.
type GetChatAdministratorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChatAdministratorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChatAdministratorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetChatAdministratorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetChatAdministratorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetChatAdministratorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetChatAdministratorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetChatAdministratorsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetChatAdministratorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChatAdministratorsOK creates a GetChatAdministratorsOK with default headers values
func NewGetChatAdministratorsOK() *GetChatAdministratorsOK {
	return &GetChatAdministratorsOK{}
}

/*GetChatAdministratorsOK handles this case with default header values.

GetChatAdministratorsOK get chat administrators o k
*/
type GetChatAdministratorsOK struct {
	Payload *models.GetChatAdministratorsOKBody
}

func (o *GetChatAdministratorsOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsOK  %+v", 200, o.Payload)
}

func (o *GetChatAdministratorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetChatAdministratorsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatAdministratorsBadRequest creates a GetChatAdministratorsBadRequest with default headers values
func NewGetChatAdministratorsBadRequest() *GetChatAdministratorsBadRequest {
	return &GetChatAdministratorsBadRequest{}
}

/*GetChatAdministratorsBadRequest handles this case with default header values.

Bad Request
*/
type GetChatAdministratorsBadRequest struct {
	Payload *models.Error
}

func (o *GetChatAdministratorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsBadRequest  %+v", 400, o.Payload)
}

func (o *GetChatAdministratorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatAdministratorsUnauthorized creates a GetChatAdministratorsUnauthorized with default headers values
func NewGetChatAdministratorsUnauthorized() *GetChatAdministratorsUnauthorized {
	return &GetChatAdministratorsUnauthorized{}
}

/*GetChatAdministratorsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChatAdministratorsUnauthorized struct {
	Payload *models.Error
}

func (o *GetChatAdministratorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChatAdministratorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatAdministratorsForbidden creates a GetChatAdministratorsForbidden with default headers values
func NewGetChatAdministratorsForbidden() *GetChatAdministratorsForbidden {
	return &GetChatAdministratorsForbidden{}
}

/*GetChatAdministratorsForbidden handles this case with default header values.

Forbidden
*/
type GetChatAdministratorsForbidden struct {
	Payload *models.Error
}

func (o *GetChatAdministratorsForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsForbidden  %+v", 403, o.Payload)
}

func (o *GetChatAdministratorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatAdministratorsNotFound creates a GetChatAdministratorsNotFound with default headers values
func NewGetChatAdministratorsNotFound() *GetChatAdministratorsNotFound {
	return &GetChatAdministratorsNotFound{}
}

/*GetChatAdministratorsNotFound handles this case with default header values.

Not Found
*/
type GetChatAdministratorsNotFound struct {
	Payload *models.Error
}

func (o *GetChatAdministratorsNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsNotFound  %+v", 404, o.Payload)
}

func (o *GetChatAdministratorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatAdministratorsEnhanceYourCalm creates a GetChatAdministratorsEnhanceYourCalm with default headers values
func NewGetChatAdministratorsEnhanceYourCalm() *GetChatAdministratorsEnhanceYourCalm {
	return &GetChatAdministratorsEnhanceYourCalm{}
}

/*GetChatAdministratorsEnhanceYourCalm handles this case with default header values.

Flood
*/
type GetChatAdministratorsEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *GetChatAdministratorsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetChatAdministratorsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatAdministratorsInternalServerError creates a GetChatAdministratorsInternalServerError with default headers values
func NewGetChatAdministratorsInternalServerError() *GetChatAdministratorsInternalServerError {
	return &GetChatAdministratorsInternalServerError{}
}

/*GetChatAdministratorsInternalServerError handles this case with default header values.

Internal
*/
type GetChatAdministratorsInternalServerError struct {
	Payload *models.Error
}

func (o *GetChatAdministratorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatAdministrators][%d] getChatAdministratorsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChatAdministratorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

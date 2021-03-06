// Code generated by go-swagger; DO NOT EDIT.

package inline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// AnswerInlineQueryReader is a Reader for the AnswerInlineQuery structure.
type AnswerInlineQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AnswerInlineQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAnswerInlineQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAnswerInlineQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAnswerInlineQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAnswerInlineQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAnswerInlineQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewAnswerInlineQueryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAnswerInlineQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAnswerInlineQueryOK creates a AnswerInlineQueryOK with default headers values
func NewAnswerInlineQueryOK() *AnswerInlineQueryOK {
	return &AnswerInlineQueryOK{}
}

/*AnswerInlineQueryOK handles this case with default header values.

AnswerInlineQueryOK answer inline query o k
*/
type AnswerInlineQueryOK struct {
	Payload *models.ResponseBool
}

func (o *AnswerInlineQueryOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryOK  %+v", 200, o.Payload)
}

func (o *AnswerInlineQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerInlineQueryBadRequest creates a AnswerInlineQueryBadRequest with default headers values
func NewAnswerInlineQueryBadRequest() *AnswerInlineQueryBadRequest {
	return &AnswerInlineQueryBadRequest{}
}

/*AnswerInlineQueryBadRequest handles this case with default header values.

Bad Request
*/
type AnswerInlineQueryBadRequest struct {
	Payload *models.Error
}

func (o *AnswerInlineQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryBadRequest  %+v", 400, o.Payload)
}

func (o *AnswerInlineQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerInlineQueryUnauthorized creates a AnswerInlineQueryUnauthorized with default headers values
func NewAnswerInlineQueryUnauthorized() *AnswerInlineQueryUnauthorized {
	return &AnswerInlineQueryUnauthorized{}
}

/*AnswerInlineQueryUnauthorized handles this case with default header values.

Unauthorized
*/
type AnswerInlineQueryUnauthorized struct {
	Payload *models.Error
}

func (o *AnswerInlineQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *AnswerInlineQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerInlineQueryForbidden creates a AnswerInlineQueryForbidden with default headers values
func NewAnswerInlineQueryForbidden() *AnswerInlineQueryForbidden {
	return &AnswerInlineQueryForbidden{}
}

/*AnswerInlineQueryForbidden handles this case with default header values.

Forbidden
*/
type AnswerInlineQueryForbidden struct {
	Payload *models.Error
}

func (o *AnswerInlineQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryForbidden  %+v", 403, o.Payload)
}

func (o *AnswerInlineQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerInlineQueryNotFound creates a AnswerInlineQueryNotFound with default headers values
func NewAnswerInlineQueryNotFound() *AnswerInlineQueryNotFound {
	return &AnswerInlineQueryNotFound{}
}

/*AnswerInlineQueryNotFound handles this case with default header values.

Not Found
*/
type AnswerInlineQueryNotFound struct {
	Payload *models.Error
}

func (o *AnswerInlineQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryNotFound  %+v", 404, o.Payload)
}

func (o *AnswerInlineQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerInlineQueryEnhanceYourCalm creates a AnswerInlineQueryEnhanceYourCalm with default headers values
func NewAnswerInlineQueryEnhanceYourCalm() *AnswerInlineQueryEnhanceYourCalm {
	return &AnswerInlineQueryEnhanceYourCalm{}
}

/*AnswerInlineQueryEnhanceYourCalm handles this case with default header values.

Flood
*/
type AnswerInlineQueryEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *AnswerInlineQueryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *AnswerInlineQueryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnswerInlineQueryInternalServerError creates a AnswerInlineQueryInternalServerError with default headers values
func NewAnswerInlineQueryInternalServerError() *AnswerInlineQueryInternalServerError {
	return &AnswerInlineQueryInternalServerError{}
}

/*AnswerInlineQueryInternalServerError handles this case with default header values.

Internal
*/
type AnswerInlineQueryInternalServerError struct {
	Payload *models.Error
}

func (o *AnswerInlineQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/answerInlineQuery][%d] answerInlineQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *AnswerInlineQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

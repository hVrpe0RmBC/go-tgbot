package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetChatMembersCountReader is a Reader for the GetChatMembersCount structure.
type GetChatMembersCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChatMembersCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChatMembersCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChatMembersCountOK creates a GetChatMembersCountOK with default headers values
func NewGetChatMembersCountOK() *GetChatMembersCountOK {
	return &GetChatMembersCountOK{}
}

/*GetChatMembersCountOK handles this case with default header values.

GetChatMembersCountOK get chat members count o k
*/
type GetChatMembersCountOK struct {
	Payload GetChatMembersCountOKBody
}

func (o *GetChatMembersCountOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChatMembersCount][%d] getChatMembersCountOK  %+v", 200, o.Payload)
}

func (o *GetChatMembersCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetChatMembersCountOKBody get chat members count o k body
swagger:model GetChatMembersCountOKBody
*/
type GetChatMembersCountOKBody struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// error code
	// Required: true
	ErrorCode *int64 `json:"error_code"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// result
	// Required: true
	Result *int64 `json:"result"`
}
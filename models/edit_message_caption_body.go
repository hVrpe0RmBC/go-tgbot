package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// EditMessageCaptionBody edit message caption body
// swagger:model EditMessageCaptionBody
type EditMessageCaptionBody struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// chat id
	ChatID interface{} `json:"chat_id,omitempty"`

	// inline message id
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// message id
	MessageID int64 `json:"message_id,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Validate validates this edit message caption body
func (m *EditMessageCaptionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditMessageCaptionBody) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
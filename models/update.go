// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Update update
// swagger:model Update
type Update struct {

	// callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// channel post
	ChannelPost *Message `json:"channel_post,omitempty"`

	// chosen inline result
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// edited channel post
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// edited message
	EditedMessage *Message `json:"edited_message,omitempty"`

	// inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// message
	Message *Message `json:"message,omitempty"`

	// pre checkout query
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// shipping query
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// update id
	UpdateID int64 `json:"update_id,omitempty"`
}

// Validate validates this update
func (m *Update) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbackQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChannelPost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChosenInlineResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEditedChannelPost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEditedMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInlineQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePreCheckoutQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShippingQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Update) validateCallbackQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.CallbackQuery) { // not required
		return nil
	}

	if m.CallbackQuery != nil {

		if err := m.CallbackQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callback_query")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateChannelPost(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelPost) { // not required
		return nil
	}

	if m.ChannelPost != nil {

		if err := m.ChannelPost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel_post")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateChosenInlineResult(formats strfmt.Registry) error {

	if swag.IsZero(m.ChosenInlineResult) { // not required
		return nil
	}

	if m.ChosenInlineResult != nil {

		if err := m.ChosenInlineResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chosen_inline_result")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateEditedChannelPost(formats strfmt.Registry) error {

	if swag.IsZero(m.EditedChannelPost) { // not required
		return nil
	}

	if m.EditedChannelPost != nil {

		if err := m.EditedChannelPost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edited_channel_post")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateEditedMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.EditedMessage) { // not required
		return nil
	}

	if m.EditedMessage != nil {

		if err := m.EditedMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edited_message")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateInlineQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.InlineQuery) { // not required
		return nil
	}

	if m.InlineQuery != nil {

		if err := m.InlineQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inline_query")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {

		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validatePreCheckoutQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.PreCheckoutQuery) { // not required
		return nil
	}

	if m.PreCheckoutQuery != nil {

		if err := m.PreCheckoutQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pre_checkout_query")
			}
			return err
		}
	}

	return nil
}

func (m *Update) validateShippingQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingQuery) { // not required
		return nil
	}

	if m.ShippingQuery != nil {

		if err := m.ShippingQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipping_query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Update) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Update) UnmarshalBinary(b []byte) error {
	var res Update
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EditMessageTextOKBody edit message text o k body
// swagger:model editMessageTextOKBody
type EditMessageTextOKBody struct {

	// description
	Description string `json:"description,omitempty"`

	// error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// ok
	Ok bool `json:"ok,omitempty"`

	// result
	Result interface{} `json:"result,omitempty"`
}

// Validate validates this edit message text o k body
func (m *EditMessageTextOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EditMessageTextOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditMessageTextOKBody) UnmarshalBinary(b []byte) error {
	var res EditMessageTextOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

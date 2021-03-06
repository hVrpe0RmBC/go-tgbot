// Code generated by go-swagger; DO NOT EDIT.

package callbacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new callbacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for callbacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AnswerCallbackQuery answer callback query API
*/
func (a *Client) AnswerCallbackQuery(params *AnswerCallbackQueryParams) (*AnswerCallbackQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnswerCallbackQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "answerCallbackQuery",
		Method:             "POST",
		PathPattern:        "/bot{token}/answerCallbackQuery",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AnswerCallbackQueryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AnswerCallbackQueryOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new chats API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for chats API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetChat get chat API
*/
func (a *Client) GetChat(params *GetChatParams) (*GetChatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChatParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChat",
		Method:             "GET",
		PathPattern:        "/bot{token}/getChat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChatOK), nil

}

/*
GetChatAdministrators get chat administrators API
*/
func (a *Client) GetChatAdministrators(params *GetChatAdministratorsParams) (*GetChatAdministratorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChatAdministratorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChatAdministrators",
		Method:             "GET",
		PathPattern:        "/bot{token}/getChatAdministrators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChatAdministratorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChatAdministratorsOK), nil

}

/*
GetChatMember get chat member API
*/
func (a *Client) GetChatMember(params *GetChatMemberParams) (*GetChatMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChatMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChatMember",
		Method:             "GET",
		PathPattern:        "/bot{token}/getChatMember",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChatMemberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChatMemberOK), nil

}

/*
GetChatMembersCount get chat members count API
*/
func (a *Client) GetChatMembersCount(params *GetChatMembersCountParams) (*GetChatMembersCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChatMembersCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChatMembersCount",
		Method:             "GET",
		PathPattern:        "/bot{token}/getChatMembersCount",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChatMembersCountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChatMembersCountOK), nil

}

/*
KickChatMember kick chat member API
*/
func (a *Client) KickChatMember(params *KickChatMemberParams) (*KickChatMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKickChatMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "kickChatMember",
		Method:             "GET",
		PathPattern:        "/bot{token}/kickChatMember",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KickChatMemberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KickChatMemberOK), nil

}

/*
LeaveChat leave chat API
*/
func (a *Client) LeaveChat(params *LeaveChatParams) (*LeaveChatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLeaveChatParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "leaveChat",
		Method:             "GET",
		PathPattern:        "/bot{token}/leaveChat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LeaveChatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LeaveChatOK), nil

}

/*
SendChatAction send chat action API
*/
func (a *Client) SendChatAction(params *SendChatActionParams) (*SendChatActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendChatActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendChatAction",
		Method:             "GET",
		PathPattern:        "/bot{token}/sendChatAction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendChatActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendChatActionOK), nil

}

/*
UnbanChatMember unban chat member API
*/
func (a *Client) UnbanChatMember(params *UnbanChatMemberParams) (*UnbanChatMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnbanChatMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unbanChatMember",
		Method:             "GET",
		PathPattern:        "/bot{token}/unbanChatMember",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnbanChatMemberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UnbanChatMemberOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
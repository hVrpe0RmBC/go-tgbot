// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// NewEditMessageTextParams creates a new EditMessageTextParams object
// with the default values initialized.
func NewEditMessageTextParams() *EditMessageTextParams {
	var ()
	return &EditMessageTextParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditMessageTextParamsWithTimeout creates a new EditMessageTextParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditMessageTextParamsWithTimeout(timeout time.Duration) *EditMessageTextParams {
	var ()
	return &EditMessageTextParams{

		timeout: timeout,
	}
}

// NewEditMessageTextParamsWithContext creates a new EditMessageTextParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditMessageTextParamsWithContext(ctx context.Context) *EditMessageTextParams {
	var ()
	return &EditMessageTextParams{

		Context: ctx,
	}
}

// NewEditMessageTextParamsWithHTTPClient creates a new EditMessageTextParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditMessageTextParamsWithHTTPClient(client *http.Client) *EditMessageTextParams {
	var ()
	return &EditMessageTextParams{
		HTTPClient: client,
	}
}

/*EditMessageTextParams contains all the parameters to send to the API endpoint
for the edit message text operation typically these are written to a http.Request
*/
type EditMessageTextParams struct {

	/*Body*/
	Body *models.EditMessageTextBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit message text params
func (o *EditMessageTextParams) WithTimeout(timeout time.Duration) *EditMessageTextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit message text params
func (o *EditMessageTextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit message text params
func (o *EditMessageTextParams) WithContext(ctx context.Context) *EditMessageTextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit message text params
func (o *EditMessageTextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit message text params
func (o *EditMessageTextParams) WithHTTPClient(client *http.Client) *EditMessageTextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit message text params
func (o *EditMessageTextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit message text params
func (o *EditMessageTextParams) WithBody(body *models.EditMessageTextBody) *EditMessageTextParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit message text params
func (o *EditMessageTextParams) SetBody(body *models.EditMessageTextBody) {
	o.Body = body
}

// WithToken adds the token to the edit message text params
func (o *EditMessageTextParams) WithToken(token *string) *EditMessageTextParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the edit message text params
func (o *EditMessageTextParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *EditMessageTextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

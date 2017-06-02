package payments

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

// NewSendInvoiceParams creates a new SendInvoiceParams object
// with the default values initialized.
func NewSendInvoiceParams() *SendInvoiceParams {
	var ()
	return &SendInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendInvoiceParamsWithTimeout creates a new SendInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendInvoiceParamsWithTimeout(timeout time.Duration) *SendInvoiceParams {
	var ()
	return &SendInvoiceParams{

		timeout: timeout,
	}
}

// NewSendInvoiceParamsWithContext creates a new SendInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendInvoiceParamsWithContext(ctx context.Context) *SendInvoiceParams {
	var ()
	return &SendInvoiceParams{

		Context: ctx,
	}
}

// NewSendInvoiceParamsWithHTTPClient creates a new SendInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendInvoiceParamsWithHTTPClient(client *http.Client) *SendInvoiceParams {
	var ()
	return &SendInvoiceParams{
		HTTPClient: client,
	}
}

/*SendInvoiceParams contains all the parameters to send to the API endpoint
for the send invoice operation typically these are written to a http.Request
*/
type SendInvoiceParams struct {

	/*Body*/
	Body *models.SendInvoiceBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send invoice params
func (o *SendInvoiceParams) WithTimeout(timeout time.Duration) *SendInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send invoice params
func (o *SendInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send invoice params
func (o *SendInvoiceParams) WithContext(ctx context.Context) *SendInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send invoice params
func (o *SendInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send invoice params
func (o *SendInvoiceParams) WithHTTPClient(client *http.Client) *SendInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send invoice params
func (o *SendInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send invoice params
func (o *SendInvoiceParams) WithBody(body *models.SendInvoiceBody) *SendInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send invoice params
func (o *SendInvoiceParams) SetBody(body *models.SendInvoiceBody) {
	o.Body = body
}

// WithToken adds the token to the send invoice params
func (o *SendInvoiceParams) WithToken(token *string) *SendInvoiceParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send invoice params
func (o *SendInvoiceParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.SendInvoiceBody)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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

// Code generated by go-swagger; DO NOT EDIT.

package updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// NewGetUpdatesParams creates a new GetUpdatesParams object
// with the default values initialized.
func NewGetUpdatesParams() *GetUpdatesParams {
	var ()
	return &GetUpdatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpdatesParamsWithTimeout creates a new GetUpdatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUpdatesParamsWithTimeout(timeout time.Duration) *GetUpdatesParams {
	var ()
	return &GetUpdatesParams{

		timeout: timeout,
	}
}

// NewGetUpdatesParamsWithContext creates a new GetUpdatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUpdatesParamsWithContext(ctx context.Context) *GetUpdatesParams {
	var ()
	return &GetUpdatesParams{

		Context: ctx,
	}
}

// NewGetUpdatesParamsWithHTTPClient creates a new GetUpdatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUpdatesParamsWithHTTPClient(client *http.Client) *GetUpdatesParams {
	var ()
	return &GetUpdatesParams{
		HTTPClient: client,
	}
}

/*GetUpdatesParams contains all the parameters to send to the API endpoint
for the get updates operation typically these are written to a http.Request
*/
type GetUpdatesParams struct {

	/*Body*/
	Body *models.GetUpdatesBody
	/*Offset*/
	Offset *int64
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get updates params
func (o *GetUpdatesParams) WithTimeout(timeout time.Duration) *GetUpdatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get updates params
func (o *GetUpdatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get updates params
func (o *GetUpdatesParams) WithContext(ctx context.Context) *GetUpdatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get updates params
func (o *GetUpdatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get updates params
func (o *GetUpdatesParams) WithHTTPClient(client *http.Client) *GetUpdatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get updates params
func (o *GetUpdatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get updates params
func (o *GetUpdatesParams) WithBody(body *models.GetUpdatesBody) *GetUpdatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get updates params
func (o *GetUpdatesParams) SetBody(body *models.GetUpdatesBody) {
	o.Body = body
}

// WithOffset adds the offset to the get updates params
func (o *GetUpdatesParams) WithOffset(offset *int64) *GetUpdatesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get updates params
func (o *GetUpdatesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithToken adds the token to the get updates params
func (o *GetUpdatesParams) WithToken(token *string) *GetUpdatesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get updates params
func (o *GetUpdatesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpdatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.GetUpdatesBody)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
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

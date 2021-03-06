// Code generated by go-swagger; DO NOT EDIT.

package stickers

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
)

// NewUploadStickerFileLinkParams creates a new UploadStickerFileLinkParams object
// with the default values initialized.
func NewUploadStickerFileLinkParams() *UploadStickerFileLinkParams {
	var ()
	return &UploadStickerFileLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadStickerFileLinkParamsWithTimeout creates a new UploadStickerFileLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadStickerFileLinkParamsWithTimeout(timeout time.Duration) *UploadStickerFileLinkParams {
	var ()
	return &UploadStickerFileLinkParams{

		timeout: timeout,
	}
}

// NewUploadStickerFileLinkParamsWithContext creates a new UploadStickerFileLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadStickerFileLinkParamsWithContext(ctx context.Context) *UploadStickerFileLinkParams {
	var ()
	return &UploadStickerFileLinkParams{

		Context: ctx,
	}
}

// NewUploadStickerFileLinkParamsWithHTTPClient creates a new UploadStickerFileLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadStickerFileLinkParamsWithHTTPClient(client *http.Client) *UploadStickerFileLinkParams {
	var ()
	return &UploadStickerFileLinkParams{
		HTTPClient: client,
	}
}

/*UploadStickerFileLinkParams contains all the parameters to send to the API endpoint
for the upload sticker file link operation typically these are written to a http.Request
*/
type UploadStickerFileLinkParams struct {

	/*PngSticker*/
	PngSticker string
	/*Token
	  bot's token to authorize the request

	*/
	Token *string
	/*UserID*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload sticker file link params
func (o *UploadStickerFileLinkParams) WithTimeout(timeout time.Duration) *UploadStickerFileLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload sticker file link params
func (o *UploadStickerFileLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload sticker file link params
func (o *UploadStickerFileLinkParams) WithContext(ctx context.Context) *UploadStickerFileLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload sticker file link params
func (o *UploadStickerFileLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload sticker file link params
func (o *UploadStickerFileLinkParams) WithHTTPClient(client *http.Client) *UploadStickerFileLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload sticker file link params
func (o *UploadStickerFileLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPngSticker adds the pngSticker to the upload sticker file link params
func (o *UploadStickerFileLinkParams) WithPngSticker(pngSticker string) *UploadStickerFileLinkParams {
	o.SetPngSticker(pngSticker)
	return o
}

// SetPngSticker adds the pngSticker to the upload sticker file link params
func (o *UploadStickerFileLinkParams) SetPngSticker(pngSticker string) {
	o.PngSticker = pngSticker
}

// WithToken adds the token to the upload sticker file link params
func (o *UploadStickerFileLinkParams) WithToken(token *string) *UploadStickerFileLinkParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the upload sticker file link params
func (o *UploadStickerFileLinkParams) SetToken(token *string) {
	o.Token = token
}

// WithUserID adds the userID to the upload sticker file link params
func (o *UploadStickerFileLinkParams) WithUserID(userID int64) *UploadStickerFileLinkParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the upload sticker file link params
func (o *UploadStickerFileLinkParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadStickerFileLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param png_sticker
	frPngSticker := o.PngSticker
	fPngSticker := frPngSticker
	if fPngSticker != "" {
		if err := r.SetFormParam("png_sticker", fPngSticker); err != nil {
			return err
		}
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	// form param user_id
	frUserID := o.UserID
	fUserID := swag.FormatInt64(frUserID)
	if fUserID != "" {
		if err := r.SetFormParam("user_id", fUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

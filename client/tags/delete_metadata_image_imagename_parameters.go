// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteMetadataImageImagenameParams creates a new DeleteMetadataImageImagenameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMetadataImageImagenameParams() *DeleteMetadataImageImagenameParams {
	return &DeleteMetadataImageImagenameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMetadataImageImagenameParamsWithTimeout creates a new DeleteMetadataImageImagenameParams object
// with the ability to set a timeout on a request.
func NewDeleteMetadataImageImagenameParamsWithTimeout(timeout time.Duration) *DeleteMetadataImageImagenameParams {
	return &DeleteMetadataImageImagenameParams{
		timeout: timeout,
	}
}

// NewDeleteMetadataImageImagenameParamsWithContext creates a new DeleteMetadataImageImagenameParams object
// with the ability to set a context for a request.
func NewDeleteMetadataImageImagenameParamsWithContext(ctx context.Context) *DeleteMetadataImageImagenameParams {
	return &DeleteMetadataImageImagenameParams{
		Context: ctx,
	}
}

// NewDeleteMetadataImageImagenameParamsWithHTTPClient creates a new DeleteMetadataImageImagenameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMetadataImageImagenameParamsWithHTTPClient(client *http.Client) *DeleteMetadataImageImagenameParams {
	return &DeleteMetadataImageImagenameParams{
		HTTPClient: client,
	}
}

/* DeleteMetadataImageImagenameParams contains all the parameters to send to the API endpoint
   for the delete metadata image imagename operation.

   Typically these are written to a http.Request.
*/
type DeleteMetadataImageImagenameParams struct {

	// TagName.
	TagName DeleteMetadataImageImagenameBody

	// Imagename.
	Imagename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete metadata image imagename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMetadataImageImagenameParams) WithDefaults() *DeleteMetadataImageImagenameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete metadata image imagename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMetadataImageImagenameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) WithTimeout(timeout time.Duration) *DeleteMetadataImageImagenameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) WithContext(ctx context.Context) *DeleteMetadataImageImagenameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) WithHTTPClient(client *http.Client) *DeleteMetadataImageImagenameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagName adds the tagName to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) WithTagName(tagName DeleteMetadataImageImagenameBody) *DeleteMetadataImageImagenameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) SetTagName(tagName DeleteMetadataImageImagenameBody) {
	o.TagName = tagName
}

// WithImagename adds the imagename to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) WithImagename(imagename string) *DeleteMetadataImageImagenameParams {
	o.SetImagename(imagename)
	return o
}

// SetImagename adds the imagename to the delete metadata image imagename params
func (o *DeleteMetadataImageImagenameParams) SetImagename(imagename string) {
	o.Imagename = imagename
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMetadataImageImagenameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.TagName); err != nil {
		return err
	}

	// path param imagename
	if err := r.SetPathParam("imagename", o.Imagename); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

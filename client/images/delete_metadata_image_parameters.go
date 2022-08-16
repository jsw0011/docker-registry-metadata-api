// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewDeleteMetadataImageParams creates a new DeleteMetadataImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMetadataImageParams() *DeleteMetadataImageParams {
	return &DeleteMetadataImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMetadataImageParamsWithTimeout creates a new DeleteMetadataImageParams object
// with the ability to set a timeout on a request.
func NewDeleteMetadataImageParamsWithTimeout(timeout time.Duration) *DeleteMetadataImageParams {
	return &DeleteMetadataImageParams{
		timeout: timeout,
	}
}

// NewDeleteMetadataImageParamsWithContext creates a new DeleteMetadataImageParams object
// with the ability to set a context for a request.
func NewDeleteMetadataImageParamsWithContext(ctx context.Context) *DeleteMetadataImageParams {
	return &DeleteMetadataImageParams{
		Context: ctx,
	}
}

// NewDeleteMetadataImageParamsWithHTTPClient creates a new DeleteMetadataImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMetadataImageParamsWithHTTPClient(client *http.Client) *DeleteMetadataImageParams {
	return &DeleteMetadataImageParams{
		HTTPClient: client,
	}
}

/* DeleteMetadataImageParams contains all the parameters to send to the API endpoint
   for the delete metadata image operation.

   Typically these are written to a http.Request.
*/
type DeleteMetadataImageParams struct {

	// ImageName.
	ImageName DeleteMetadataImageBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete metadata image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMetadataImageParams) WithDefaults() *DeleteMetadataImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete metadata image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMetadataImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete metadata image params
func (o *DeleteMetadataImageParams) WithTimeout(timeout time.Duration) *DeleteMetadataImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete metadata image params
func (o *DeleteMetadataImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete metadata image params
func (o *DeleteMetadataImageParams) WithContext(ctx context.Context) *DeleteMetadataImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete metadata image params
func (o *DeleteMetadataImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete metadata image params
func (o *DeleteMetadataImageParams) WithHTTPClient(client *http.Client) *DeleteMetadataImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete metadata image params
func (o *DeleteMetadataImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageName adds the imageName to the delete metadata image params
func (o *DeleteMetadataImageParams) WithImageName(imageName DeleteMetadataImageBody) *DeleteMetadataImageParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the delete metadata image params
func (o *DeleteMetadataImageParams) SetImageName(imageName DeleteMetadataImageBody) {
	o.ImageName = imageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMetadataImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.ImageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
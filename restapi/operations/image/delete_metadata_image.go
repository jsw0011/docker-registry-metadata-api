// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteMetadataImageHandlerFunc turns a function with the right signature into a delete metadata image handler
type DeleteMetadataImageHandlerFunc func(DeleteMetadataImageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMetadataImageHandlerFunc) Handle(params DeleteMetadataImageParams) middleware.Responder {
	return fn(params)
}

// DeleteMetadataImageHandler interface for that can handle valid delete metadata image params
type DeleteMetadataImageHandler interface {
	Handle(DeleteMetadataImageParams) middleware.Responder
}

// NewDeleteMetadataImage creates a new http.Handler for the delete metadata image operation
func NewDeleteMetadataImage(ctx *middleware.Context, handler DeleteMetadataImageHandler) *DeleteMetadataImage {
	return &DeleteMetadataImage{Context: ctx, Handler: handler}
}

/* DeleteMetadataImage swagger:route DELETE /metadata/image Image deleteMetadataImage

Unregister image from metadata database

All tags and metadata will be removed

*/
type DeleteMetadataImage struct {
	Context *middleware.Context
	Handler DeleteMetadataImageHandler
}

func (o *DeleteMetadataImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMetadataImageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteMetadataImageBody delete metadata image body
//
// swagger:model DeleteMetadataImageBody
type DeleteMetadataImageBody struct {

	// image name
	ImageName string `json:"ImageName,omitempty"`
}

// Validate validates this delete metadata image body
func (o *DeleteMetadataImageBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete metadata image body based on context it is used
func (o *DeleteMetadataImageBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteMetadataImageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteMetadataImageBody) UnmarshalBinary(b []byte) error {
	var res DeleteMetadataImageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

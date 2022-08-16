// Code generated by go-swagger; DO NOT EDIT.

package image_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteMetadataImageTagHandlerFunc turns a function with the right signature into a delete metadata image tag handler
type DeleteMetadataImageTagHandlerFunc func(DeleteMetadataImageTagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMetadataImageTagHandlerFunc) Handle(params DeleteMetadataImageTagParams) middleware.Responder {
	return fn(params)
}

// DeleteMetadataImageTagHandler interface for that can handle valid delete metadata image tag params
type DeleteMetadataImageTagHandler interface {
	Handle(DeleteMetadataImageTagParams) middleware.Responder
}

// NewDeleteMetadataImageTag creates a new http.Handler for the delete metadata image tag operation
func NewDeleteMetadataImageTag(ctx *middleware.Context, handler DeleteMetadataImageTagHandler) *DeleteMetadataImageTag {
	return &DeleteMetadataImageTag{Context: ctx, Handler: handler}
}

/* DeleteMetadataImageTag swagger:route DELETE /metadata/image/tag ImageTag deleteMetadataImageTag

Remove tag

Remove tag and all its metadata

*/
type DeleteMetadataImageTag struct {
	Context *middleware.Context
	Handler DeleteMetadataImageTagHandler
}

func (o *DeleteMetadataImageTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMetadataImageTagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteMetadataImageTagBody delete metadata image tag body
//
// swagger:model DeleteMetadataImageTagBody
type DeleteMetadataImageTagBody struct {

	// image name
	ImageName string `json:"ImageName,omitempty"`

	// tag name
	TagName string `json:"TagName,omitempty"`
}

// Validate validates this delete metadata image tag body
func (o *DeleteMetadataImageTagBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete metadata image tag body based on context it is used
func (o *DeleteMetadataImageTagBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteMetadataImageTagBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteMetadataImageTagBody) UnmarshalBinary(b []byte) error {
	var res DeleteMetadataImageTagBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
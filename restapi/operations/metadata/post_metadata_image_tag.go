// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostMetadataImageTagHandlerFunc turns a function with the right signature into a post metadata image tag handler
type PostMetadataImageTagHandlerFunc func(PostMetadataImageTagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMetadataImageTagHandlerFunc) Handle(params PostMetadataImageTagParams) middleware.Responder {
	return fn(params)
}

// PostMetadataImageTagHandler interface for that can handle valid post metadata image tag params
type PostMetadataImageTagHandler interface {
	Handle(PostMetadataImageTagParams) middleware.Responder
}

// NewPostMetadataImageTag creates a new http.Handler for the post metadata image tag operation
func NewPostMetadataImageTag(ctx *middleware.Context, handler PostMetadataImageTagHandler) *PostMetadataImageTag {
	return &PostMetadataImageTag{Context: ctx, Handler: handler}
}

/* PostMetadataImageTag swagger:route POST /metadata/image/tag Metadata postMetadataImageTag

Get metadata

*/
type PostMetadataImageTag struct {
	Context *middleware.Context
	Handler PostMetadataImageTagHandler
}

func (o *PostMetadataImageTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostMetadataImageTagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostMetadataImageTagBody post metadata image tag body
//
// swagger:model PostMetadataImageTagBody
type PostMetadataImageTagBody struct {

	// image name
	ImageName string `json:"ImageName,omitempty"`

	// tag name
	TagName string `json:"TagName,omitempty"`
}

// Validate validates this post metadata image tag body
func (o *PostMetadataImageTagBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post metadata image tag body based on context it is used
func (o *PostMetadataImageTagBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMetadataImageTagBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMetadataImageTagBody) UnmarshalBinary(b []byte) error {
	var res PostMetadataImageTagBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

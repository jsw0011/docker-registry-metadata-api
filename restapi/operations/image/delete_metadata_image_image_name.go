// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMetadataImageImageNameHandlerFunc turns a function with the right signature into a delete metadata image image name handler
type DeleteMetadataImageImageNameHandlerFunc func(DeleteMetadataImageImageNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMetadataImageImageNameHandlerFunc) Handle(params DeleteMetadataImageImageNameParams) middleware.Responder {
	return fn(params)
}

// DeleteMetadataImageImageNameHandler interface for that can handle valid delete metadata image image name params
type DeleteMetadataImageImageNameHandler interface {
	Handle(DeleteMetadataImageImageNameParams) middleware.Responder
}

// NewDeleteMetadataImageImageName creates a new http.Handler for the delete metadata image image name operation
func NewDeleteMetadataImageImageName(ctx *middleware.Context, handler DeleteMetadataImageImageNameHandler) *DeleteMetadataImageImageName {
	return &DeleteMetadataImageImageName{Context: ctx, Handler: handler}
}

/* DeleteMetadataImageImageName swagger:route DELETE /metadata/image/{ImageName} Image deleteMetadataImageImageName

Unregister image from metadata database

All tags and metadata will be removed

*/
type DeleteMetadataImageImageName struct {
	Context *middleware.Context
	Handler DeleteMetadataImageImageNameHandler
}

func (o *DeleteMetadataImageImageName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMetadataImageImageNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

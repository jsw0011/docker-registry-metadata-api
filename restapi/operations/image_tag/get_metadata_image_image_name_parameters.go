// Code generated by go-swagger; DO NOT EDIT.

package image_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetMetadataImageImageNameParams creates a new GetMetadataImageImageNameParams object
//
// There are no default values defined in the spec.
func NewGetMetadataImageImageNameParams() GetMetadataImageImageNameParams {

	return GetMetadataImageImageNameParams{}
}

// GetMetadataImageImageNameParams contains all the bound params for the get metadata image image name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMetadataImageImageName
type GetMetadataImageImageNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ImageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMetadataImageImageNameParams() beforehand.
func (o *GetMetadataImageImageNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rImageName, rhkImageName, _ := route.Params.GetOK("ImageName")
	if err := o.bindImageName(rImageName, rhkImageName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindImageName binds and validates parameter ImageName from path.
func (o *GetMetadataImageImageNameParams) bindImageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ImageName = raw

	return nil
}
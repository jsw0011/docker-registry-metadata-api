// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"docker_image_metadata_api/models"
)

// GetMetadataImagesOKCode is the HTTP code returned for type GetMetadataImagesOK
const GetMetadataImagesOKCode int = 200

/*GetMetadataImagesOK List of registered images in the metadata service

swagger:response getMetadataImagesOK
*/
type GetMetadataImagesOK struct {

	/*
	  In: Body
	*/
	Payload models.ListOfImages `json:"body,omitempty"`
}

// NewGetMetadataImagesOK creates GetMetadataImagesOK with default headers values
func NewGetMetadataImagesOK() *GetMetadataImagesOK {

	return &GetMetadataImagesOK{}
}

// WithPayload adds the payload to the get metadata images o k response
func (o *GetMetadataImagesOK) WithPayload(payload models.ListOfImages) *GetMetadataImagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metadata images o k response
func (o *GetMetadataImagesOK) SetPayload(payload models.ListOfImages) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetadataImagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ListOfImages{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMetadataImagesInternalServerErrorCode is the HTTP code returned for type GetMetadataImagesInternalServerError
const GetMetadataImagesInternalServerErrorCode int = 500

/*GetMetadataImagesInternalServerError Internal Error Message

swagger:response getMetadataImagesInternalServerError
*/
type GetMetadataImagesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetMetadataImagesInternalServerError creates GetMetadataImagesInternalServerError with default headers values
func NewGetMetadataImagesInternalServerError() *GetMetadataImagesInternalServerError {

	return &GetMetadataImagesInternalServerError{}
}

// WithPayload adds the payload to the get metadata images internal server error response
func (o *GetMetadataImagesInternalServerError) WithPayload(payload *models.ErrorResponse) *GetMetadataImagesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metadata images internal server error response
func (o *GetMetadataImagesInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetadataImagesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
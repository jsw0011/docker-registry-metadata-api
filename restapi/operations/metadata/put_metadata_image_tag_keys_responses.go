// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"docker_image_metadata_api/models"
)

// PutMetadataImageTagKeysOKCode is the HTTP code returned for type PutMetadataImageTagKeysOK
const PutMetadataImageTagKeysOKCode int = 200

/*PutMetadataImageTagKeysOK Successfully updated

swagger:response putMetadataImageTagKeysOK
*/
type PutMetadataImageTagKeysOK struct {
}

// NewPutMetadataImageTagKeysOK creates PutMetadataImageTagKeysOK with default headers values
func NewPutMetadataImageTagKeysOK() *PutMetadataImageTagKeysOK {

	return &PutMetadataImageTagKeysOK{}
}

// WriteResponse to the client
func (o *PutMetadataImageTagKeysOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutMetadataImageTagKeysNotFoundCode is the HTTP code returned for type PutMetadataImageTagKeysNotFound
const PutMetadataImageTagKeysNotFoundCode int = 404

/*PutMetadataImageTagKeysNotFound Not Found

swagger:response putMetadataImageTagKeysNotFound
*/
type PutMetadataImageTagKeysNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFoundResponse `json:"body,omitempty"`
}

// NewPutMetadataImageTagKeysNotFound creates PutMetadataImageTagKeysNotFound with default headers values
func NewPutMetadataImageTagKeysNotFound() *PutMetadataImageTagKeysNotFound {

	return &PutMetadataImageTagKeysNotFound{}
}

// WithPayload adds the payload to the put metadata image tag keys not found response
func (o *PutMetadataImageTagKeysNotFound) WithPayload(payload *models.NotFoundResponse) *PutMetadataImageTagKeysNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put metadata image tag keys not found response
func (o *PutMetadataImageTagKeysNotFound) SetPayload(payload *models.NotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMetadataImageTagKeysNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutMetadataImageTagKeysInternalServerErrorCode is the HTTP code returned for type PutMetadataImageTagKeysInternalServerError
const PutMetadataImageTagKeysInternalServerErrorCode int = 500

/*PutMetadataImageTagKeysInternalServerError Internal Error Message

swagger:response putMetadataImageTagKeysInternalServerError
*/
type PutMetadataImageTagKeysInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPutMetadataImageTagKeysInternalServerError creates PutMetadataImageTagKeysInternalServerError with default headers values
func NewPutMetadataImageTagKeysInternalServerError() *PutMetadataImageTagKeysInternalServerError {

	return &PutMetadataImageTagKeysInternalServerError{}
}

// WithPayload adds the payload to the put metadata image tag keys internal server error response
func (o *PutMetadataImageTagKeysInternalServerError) WithPayload(payload *models.ErrorResponse) *PutMetadataImageTagKeysInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put metadata image tag keys internal server error response
func (o *PutMetadataImageTagKeysInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMetadataImageTagKeysInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

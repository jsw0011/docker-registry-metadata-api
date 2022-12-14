// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"docker_image_metadata_api/models"
)

// PostMetadataImageTagKeysCreatedCode is the HTTP code returned for type PostMetadataImageTagKeysCreated
const PostMetadataImageTagKeysCreatedCode int = 201

/*PostMetadataImageTagKeysCreated Created

swagger:response postMetadataImageTagKeysCreated
*/
type PostMetadataImageTagKeysCreated struct {
}

// NewPostMetadataImageTagKeysCreated creates PostMetadataImageTagKeysCreated with default headers values
func NewPostMetadataImageTagKeysCreated() *PostMetadataImageTagKeysCreated {

	return &PostMetadataImageTagKeysCreated{}
}

// WriteResponse to the client
func (o *PostMetadataImageTagKeysCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostMetadataImageTagKeysNotFoundCode is the HTTP code returned for type PostMetadataImageTagKeysNotFound
const PostMetadataImageTagKeysNotFoundCode int = 404

/*PostMetadataImageTagKeysNotFound Image with tag not found

swagger:response postMetadataImageTagKeysNotFound
*/
type PostMetadataImageTagKeysNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFoundResponse `json:"body,omitempty"`
}

// NewPostMetadataImageTagKeysNotFound creates PostMetadataImageTagKeysNotFound with default headers values
func NewPostMetadataImageTagKeysNotFound() *PostMetadataImageTagKeysNotFound {

	return &PostMetadataImageTagKeysNotFound{}
}

// WithPayload adds the payload to the post metadata image tag keys not found response
func (o *PostMetadataImageTagKeysNotFound) WithPayload(payload *models.NotFoundResponse) *PostMetadataImageTagKeysNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post metadata image tag keys not found response
func (o *PostMetadataImageTagKeysNotFound) SetPayload(payload *models.NotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMetadataImageTagKeysNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMetadataImageTagKeysConflictCode is the HTTP code returned for type PostMetadataImageTagKeysConflict
const PostMetadataImageTagKeysConflictCode int = 409

/*PostMetadataImageTagKeysConflict Key already exists

swagger:response postMetadataImageTagKeysConflict
*/
type PostMetadataImageTagKeysConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ConflictResponse `json:"body,omitempty"`
}

// NewPostMetadataImageTagKeysConflict creates PostMetadataImageTagKeysConflict with default headers values
func NewPostMetadataImageTagKeysConflict() *PostMetadataImageTagKeysConflict {

	return &PostMetadataImageTagKeysConflict{}
}

// WithPayload adds the payload to the post metadata image tag keys conflict response
func (o *PostMetadataImageTagKeysConflict) WithPayload(payload *models.ConflictResponse) *PostMetadataImageTagKeysConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post metadata image tag keys conflict response
func (o *PostMetadataImageTagKeysConflict) SetPayload(payload *models.ConflictResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMetadataImageTagKeysConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMetadataImageTagKeysInternalServerErrorCode is the HTTP code returned for type PostMetadataImageTagKeysInternalServerError
const PostMetadataImageTagKeysInternalServerErrorCode int = 500

/*PostMetadataImageTagKeysInternalServerError Internal Error Message

swagger:response postMetadataImageTagKeysInternalServerError
*/
type PostMetadataImageTagKeysInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostMetadataImageTagKeysInternalServerError creates PostMetadataImageTagKeysInternalServerError with default headers values
func NewPostMetadataImageTagKeysInternalServerError() *PostMetadataImageTagKeysInternalServerError {

	return &PostMetadataImageTagKeysInternalServerError{}
}

// WithPayload adds the payload to the post metadata image tag keys internal server error response
func (o *PostMetadataImageTagKeysInternalServerError) WithPayload(payload *models.ErrorResponse) *PostMetadataImageTagKeysInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post metadata image tag keys internal server error response
func (o *PostMetadataImageTagKeysInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMetadataImageTagKeysInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

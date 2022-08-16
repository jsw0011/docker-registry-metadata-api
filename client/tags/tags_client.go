// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the tags client
type API interface {
	/*
	   DeleteMetadataImageImagename removes tag

	   Remove tag and all his*/
	DeleteMetadataImageImagename(ctx context.Context, params *DeleteMetadataImageImagenameParams) (*DeleteMetadataImageImagenameCreated, error)
	/*
	   GetMetadataImageImagename lists of registered image tags in the metadata service*/
	GetMetadataImageImagename(ctx context.Context, params *GetMetadataImageImagenameParams) (*GetMetadataImageImagenameOK, error)
	/*
	   PostMetadataImageImagename registers new tag

	   Register and verify if the image with given tag exists in docker registry*/
	PostMetadataImageImagename(ctx context.Context, params *PostMetadataImageImagenameParams) (*PostMetadataImageImagenameCreated, error)
}

// New creates a new tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteMetadataImageImagename removes tag

Remove tag and all his
*/
func (a *Client) DeleteMetadataImageImagename(ctx context.Context, params *DeleteMetadataImageImagenameParams) (*DeleteMetadataImageImagenameCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMetadataImageImagename",
		Method:             "DELETE",
		PathPattern:        "/metadata/image/{imagename}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteMetadataImageImagenameReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMetadataImageImagenameCreated), nil

}

/*
GetMetadataImageImagename lists of registered image tags in the metadata service
*/
func (a *Client) GetMetadataImageImagename(ctx context.Context, params *GetMetadataImageImagenameParams) (*GetMetadataImageImagenameOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetadataImageImagename",
		Method:             "GET",
		PathPattern:        "/metadata/image/{imagename}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMetadataImageImagenameReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetadataImageImagenameOK), nil

}

/*
PostMetadataImageImagename registers new tag

Register and verify if the image with given tag exists in docker registry
*/
func (a *Client) PostMetadataImageImagename(ctx context.Context, params *PostMetadataImageImagenameParams) (*PostMetadataImageImagenameCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostMetadataImageImagename",
		Method:             "POST",
		PathPattern:        "/metadata/image/{imagename}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostMetadataImageImagenameReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostMetadataImageImagenameCreated), nil

}

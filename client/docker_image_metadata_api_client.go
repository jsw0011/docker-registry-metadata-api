// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"docker_image_metadata_api/client/images"
	"docker_image_metadata_api/client/metadata"
	"docker_image_metadata_api/client/tags"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// New creates a new docker image metadata API HTTP client.
func New(c Config) *DockerImageMetadataAPI {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(DockerImageMetadataAPI)
	cli.Transport = transport
	cli.Images = images.New(transport, strfmt.Default, c.AuthInfo)
	cli.Metadata = metadata.New(transport, strfmt.Default, c.AuthInfo)
	cli.Tags = tags.New(transport, strfmt.Default, c.AuthInfo)
	return cli
}

// DockerImageMetadataAPI is a client for docker image metadata API
type DockerImageMetadataAPI struct {
	Images    *images.Client
	Metadata  *metadata.Client
	Tags      *tags.Client
	Transport runtime.ClientTransport
}
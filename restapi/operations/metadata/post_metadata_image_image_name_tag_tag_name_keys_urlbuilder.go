// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PostMetadataImageImageNameTagTagNameKeysURL generates an URL for the post metadata image image name tag tag name keys operation
type PostMetadataImageImageNameTagTagNameKeysURL struct {
	ImageName string
	TagName   string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostMetadataImageImageNameTagTagNameKeysURL) WithBasePath(bp string) *PostMetadataImageImageNameTagTagNameKeysURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostMetadataImageImageNameTagTagNameKeysURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostMetadataImageImageNameTagTagNameKeysURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/metadata/image/{ImageName}/tag/{TagName}/keys"

	imageName := o.ImageName
	if imageName != "" {
		_path = strings.Replace(_path, "{ImageName}", imageName, -1)
	} else {
		return nil, errors.New("imageName is required on PostMetadataImageImageNameTagTagNameKeysURL")
	}

	tagName := o.TagName
	if tagName != "" {
		_path = strings.Replace(_path, "{TagName}", tagName, -1)
	} else {
		return nil, errors.New("tagName is required on PostMetadataImageImageNameTagTagNameKeysURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostMetadataImageImageNameTagTagNameKeysURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostMetadataImageImageNameTagTagNameKeysURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostMetadataImageImageNameTagTagNameKeysURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostMetadataImageImageNameTagTagNameKeysURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostMetadataImageImageNameTagTagNameKeysURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PostMetadataImageImageNameTagTagNameKeysURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

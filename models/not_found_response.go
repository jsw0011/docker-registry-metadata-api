// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotFoundResponse not found response
//
// swagger:model NotFoundResponse
type NotFoundResponse struct {

	// message
	Message string `json:"Message,omitempty"`
}

// Validate validates this not found response
func (m *NotFoundResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this not found response based on context it is used
func (m *NotFoundResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotFoundResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotFoundResponse) UnmarshalBinary(b []byte) error {
	var res NotFoundResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NetworkIds network ids
// swagger:model NetworkIDs
type NetworkIds struct {

	// an array of network IDs
	NetworkIds []string `json:"networkIDs"`
}

// Validate validates this network ids
func (m *NetworkIds) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkIds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkIds) UnmarshalBinary(b []byte) error {
	var res NetworkIds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
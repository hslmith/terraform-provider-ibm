// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ServiceInstanceAsyncOperation service instance async operation
// swagger:model ServiceInstanceAsyncOperation
type ServiceInstanceAsyncOperation struct {

	// dashboard url
	DashboardURL string `json:"dashboard_url,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`
}

// Validate validates this service instance async operation
func (m *ServiceInstanceAsyncOperation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstanceAsyncOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstanceAsyncOperation) UnmarshalBinary(b []byte) error {
	var res ServiceInstanceAsyncOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Operator operator
//
// swagger:model operator
type Operator struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// operator type
	OperatorType OperatorType `json:"operator_type,omitempty"`

	// JSON-formatted string containing the properties for each operator
	Properties string `json:"properties,omitempty"`
}

// Validate validates this operator
func (m *Operator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Operator) validateOperatorType(formats strfmt.Registry) error {

	if swag.IsZero(m.OperatorType) { // not required
		return nil
	}

	if err := m.OperatorType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Operator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Operator) UnmarshalBinary(b []byte) error {
	var res Operator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

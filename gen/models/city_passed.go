// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CityPassed city passed
//
// swagger:model CityPassed
type CityPassed struct {

	// city
	// Required: true
	City struct {
		City
	} `json:"city" gorm:"foreigKey:id;type:string;not null"`

	// route id
	// Required: true
	RouteID struct {
		Route
	} `json:"route_id" gorm:"foreigKey:id;type:integer;not null"`
}

// Validate validates this city passed
func (m *CityPassed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CityPassed) validateCity(formats strfmt.Registry) error {

	return nil
}

func (m *CityPassed) validateRouteID(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this city passed based on the context it is used
func (m *CityPassed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CityPassed) contextValidateCity(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CityPassed) contextValidateRouteID(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CityPassed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CityPassed) UnmarshalBinary(b []byte) error {
	var res CityPassed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

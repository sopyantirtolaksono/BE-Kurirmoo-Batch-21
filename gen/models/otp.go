// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Otp otp
//
// swagger:model Otp
type Otp struct {
	ModelTrackTime

	ModelIdentifier

	// is used
	IsUsed bool `json:"is_used,omitempty"`

	// otp
	Otp string `json:"otp,omitempty" gorm:"type:varchar(5);not null"`

	// otp type
	// Enum: [whatsapp]
	OtpType string `json:"otp_type,omitempty" gorm:"type:varchar(255);not null;default:'whatsapp'"`

	// phone number
	// Required: true
	PhoneNumber *string `json:"phone_number" gorm:"type:varchar(40);not null"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Otp) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModelTrackTime
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModelTrackTime = aO0

	// AO1
	var aO1 ModelIdentifier
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ModelIdentifier = aO1

	// AO2
	var dataAO2 struct {
		IsUsed bool `json:"is_used,omitempty"`

		Otp string `json:"otp,omitempty"`

		OtpType string `json:"otp_type,omitempty"`

		PhoneNumber *string `json:"phone_number"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.IsUsed = dataAO2.IsUsed

	m.Otp = dataAO2.Otp

	m.OtpType = dataAO2.OtpType

	m.PhoneNumber = dataAO2.PhoneNumber

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Otp) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.ModelTrackTime)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ModelIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		IsUsed bool `json:"is_used,omitempty"`

		Otp string `json:"otp,omitempty"`

		OtpType string `json:"otp_type,omitempty"`

		PhoneNumber *string `json:"phone_number"`
	}

	dataAO2.IsUsed = m.IsUsed

	dataAO2.Otp = m.Otp

	dataAO2.OtpType = m.OtpType

	dataAO2.PhoneNumber = m.PhoneNumber

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this otp
func (m *Otp) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtpType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var otpTypeOtpTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["whatsapp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		otpTypeOtpTypePropEnum = append(otpTypeOtpTypePropEnum, v)
	}
}

// property enum
func (m *Otp) validateOtpTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, otpTypeOtpTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Otp) validateOtpType(formats strfmt.Registry) error {

	if swag.IsZero(m.OtpType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOtpTypeEnum("otp_type", "body", m.OtpType); err != nil {
		return err
	}

	return nil
}

func (m *Otp) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phone_number", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this otp based on the context it is used
func (m *Otp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Otp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Otp) UnmarshalBinary(b []byte) error {
	var res Otp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

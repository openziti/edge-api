// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CurrentAPISessionCertificateDetail current Api session certificate detail
//
// swagger:model currentApiSessionCertificateDetail
type CurrentAPISessionCertificateDetail struct {
	BaseEntity

	// certificate
	// Required: true
	Certificate *string `json:"certificate"`

	// fingerprint
	// Required: true
	Fingerprint *string `json:"fingerprint"`

	// subject
	// Required: true
	Subject *string `json:"subject"`

	// valid from
	// Required: true
	// Format: date-time
	ValidFrom *strfmt.DateTime `json:"validFrom"`

	// valid to
	// Required: true
	// Format: date-time
	ValidTo *strfmt.DateTime `json:"validTo"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CurrentAPISessionCertificateDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		Certificate *string `json:"certificate"`

		Fingerprint *string `json:"fingerprint"`

		Subject *string `json:"subject"`

		ValidFrom *strfmt.DateTime `json:"validFrom"`

		ValidTo *strfmt.DateTime `json:"validTo"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Certificate = dataAO1.Certificate

	m.Fingerprint = dataAO1.Fingerprint

	m.Subject = dataAO1.Subject

	m.ValidFrom = dataAO1.ValidFrom

	m.ValidTo = dataAO1.ValidTo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CurrentAPISessionCertificateDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Certificate *string `json:"certificate"`

		Fingerprint *string `json:"fingerprint"`

		Subject *string `json:"subject"`

		ValidFrom *strfmt.DateTime `json:"validFrom"`

		ValidTo *strfmt.DateTime `json:"validTo"`
	}

	dataAO1.Certificate = m.Certificate

	dataAO1.Fingerprint = m.Fingerprint

	dataAO1.Subject = m.Subject

	dataAO1.ValidFrom = m.ValidFrom

	dataAO1.ValidTo = m.ValidTo

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this current Api session certificate detail
func (m *CurrentAPISessionCertificateDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentAPISessionCertificateDetail) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *CurrentAPISessionCertificateDetail) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *CurrentAPISessionCertificateDetail) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *CurrentAPISessionCertificateDetail) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CurrentAPISessionCertificateDetail) validateValidTo(formats strfmt.Registry) error {

	if err := validate.Required("validTo", "body", m.ValidTo); err != nil {
		return err
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this current Api session certificate detail based on the context it is used
func (m *CurrentAPISessionCertificateDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CurrentAPISessionCertificateDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentAPISessionCertificateDetail) UnmarshalBinary(b []byte) error {
	var res CurrentAPISessionCertificateDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

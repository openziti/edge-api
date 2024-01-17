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

// ControllerDetail A controller resource
//
// swagger:model controllerDetail
type ControllerDetail struct {
	BaseEntity

	// address
	// Required: true
	Address *string `json:"address"`

	// cert pem
	// Required: true
	CertPem *string `json:"certPem"`

	// fingerprint
	// Required: true
	Fingerprint *string `json:"fingerprint"`

	// is online
	// Required: true
	IsOnline *bool `json:"isOnline"`

	// last joined at
	// Required: true
	// Format: date-time
	LastJoinedAt *strfmt.DateTime `json:"lastJoinedAt"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ControllerDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		Address *string `json:"address"`

		CertPem *string `json:"certPem"`

		Fingerprint *string `json:"fingerprint"`

		IsOnline *bool `json:"isOnline"`

		LastJoinedAt *strfmt.DateTime `json:"lastJoinedAt"`

		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Address = dataAO1.Address

	m.CertPem = dataAO1.CertPem

	m.Fingerprint = dataAO1.Fingerprint

	m.IsOnline = dataAO1.IsOnline

	m.LastJoinedAt = dataAO1.LastJoinedAt

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ControllerDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Address *string `json:"address"`

		CertPem *string `json:"certPem"`

		Fingerprint *string `json:"fingerprint"`

		IsOnline *bool `json:"isOnline"`

		LastJoinedAt *strfmt.DateTime `json:"lastJoinedAt"`

		Name *string `json:"name"`
	}

	dataAO1.Address = m.Address

	dataAO1.CertPem = m.CertPem

	dataAO1.Fingerprint = m.Fingerprint

	dataAO1.IsOnline = m.IsOnline

	dataAO1.LastJoinedAt = m.LastJoinedAt

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this controller detail
func (m *ControllerDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertPem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastJoinedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerDetail) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *ControllerDetail) validateCertPem(formats strfmt.Registry) error {

	if err := validate.Required("certPem", "body", m.CertPem); err != nil {
		return err
	}

	return nil
}

func (m *ControllerDetail) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *ControllerDetail) validateIsOnline(formats strfmt.Registry) error {

	if err := validate.Required("isOnline", "body", m.IsOnline); err != nil {
		return err
	}

	return nil
}

func (m *ControllerDetail) validateLastJoinedAt(formats strfmt.Registry) error {

	if err := validate.Required("lastJoinedAt", "body", m.LastJoinedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("lastJoinedAt", "body", "date-time", m.LastJoinedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ControllerDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this controller detail based on the context it is used
func (m *ControllerDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *ControllerDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerDetail) UnmarshalBinary(b []byte) error {
	var res ControllerDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
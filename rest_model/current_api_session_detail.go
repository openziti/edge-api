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

// CurrentAPISessionDetail An API Session object for the current API session
//
// swagger:model currentApiSessionDetail
type CurrentAPISessionDetail struct {
	APISessionDetail

	// expiration seconds
	// Required: true
	ExpirationSeconds *int64 `json:"expirationSeconds"`

	// expires at
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expiresAt"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CurrentAPISessionDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 APISessionDetail
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.APISessionDetail = aO0

	// AO1
	var dataAO1 struct {
		ExpirationSeconds *int64 `json:"expirationSeconds"`

		ExpiresAt *strfmt.DateTime `json:"expiresAt"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ExpirationSeconds = dataAO1.ExpirationSeconds

	m.ExpiresAt = dataAO1.ExpiresAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CurrentAPISessionDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.APISessionDetail)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ExpirationSeconds *int64 `json:"expirationSeconds"`

		ExpiresAt *strfmt.DateTime `json:"expiresAt"`
	}

	dataAO1.ExpirationSeconds = m.ExpirationSeconds

	dataAO1.ExpiresAt = m.ExpiresAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this current Api session detail
func (m *CurrentAPISessionDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with APISessionDetail
	if err := m.APISessionDetail.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentAPISessionDetail) validateExpirationSeconds(formats strfmt.Registry) error {

	if err := validate.Required("expirationSeconds", "body", m.ExpirationSeconds); err != nil {
		return err
	}

	return nil
}

func (m *CurrentAPISessionDetail) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expiresAt", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this current Api session detail based on the context it is used
func (m *CurrentAPISessionDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with APISessionDetail
	if err := m.APISessionDetail.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CurrentAPISessionDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentAPISessionDetail) UnmarshalBinary(b []byte) error {
	var res CurrentAPISessionDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

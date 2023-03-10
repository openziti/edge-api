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

// CaUpdate ca update
//
// swagger:model caUpdate
type CaUpdate struct {

	// external Id claim
	ExternalIDClaim *ExternalIDClaim `json:"externalIdClaim,omitempty"`

	// identity name format
	// Required: true
	IdentityNameFormat *string `json:"identityNameFormat"`

	// identity roles
	// Required: true
	IdentityRoles Roles `json:"identityRoles"`

	// is auth enabled
	// Example: true
	// Required: true
	IsAuthEnabled *bool `json:"isAuthEnabled"`

	// is auto ca enrollment enabled
	// Example: true
	// Required: true
	IsAutoCaEnrollmentEnabled *bool `json:"isAutoCaEnrollmentEnabled"`

	// is ott ca enrollment enabled
	// Example: true
	// Required: true
	IsOttCaEnrollmentEnabled *bool `json:"isOttCaEnrollmentEnabled"`

	// name
	// Example: My CA
	// Required: true
	Name *string `json:"name"`

	// tags
	Tags *Tags `json:"tags,omitempty"`
}

// Validate validates this ca update
func (m *CaUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalIDClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityNameFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAuthEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAutoCaEnrollmentEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsOttCaEnrollmentEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaUpdate) validateExternalIDClaim(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalIDClaim) { // not required
		return nil
	}

	if m.ExternalIDClaim != nil {
		if err := m.ExternalIDClaim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalIdClaim")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalIdClaim")
			}
			return err
		}
	}

	return nil
}

func (m *CaUpdate) validateIdentityNameFormat(formats strfmt.Registry) error {

	if err := validate.Required("identityNameFormat", "body", m.IdentityNameFormat); err != nil {
		return err
	}

	return nil
}

func (m *CaUpdate) validateIdentityRoles(formats strfmt.Registry) error {

	if err := validate.Required("identityRoles", "body", m.IdentityRoles); err != nil {
		return err
	}

	if err := m.IdentityRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identityRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identityRoles")
		}
		return err
	}

	return nil
}

func (m *CaUpdate) validateIsAuthEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isAuthEnabled", "body", m.IsAuthEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CaUpdate) validateIsAutoCaEnrollmentEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isAutoCaEnrollmentEnabled", "body", m.IsAutoCaEnrollmentEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CaUpdate) validateIsOttCaEnrollmentEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isOttCaEnrollmentEnabled", "body", m.IsOttCaEnrollmentEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CaUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CaUpdate) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ca update based on the context it is used
func (m *CaUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalIDClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaUpdate) contextValidateExternalIDClaim(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalIDClaim != nil {
		if err := m.ExternalIDClaim.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalIdClaim")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalIdClaim")
			}
			return err
		}
	}

	return nil
}

func (m *CaUpdate) contextValidateIdentityRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IdentityRoles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identityRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identityRoles")
		}
		return err
	}

	return nil
}

func (m *CaUpdate) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if m.Tags != nil {
		if err := m.Tags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CaUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CaUpdate) UnmarshalBinary(b []byte) error {
	var res CaUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

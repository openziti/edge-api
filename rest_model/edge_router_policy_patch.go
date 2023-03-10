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
)

// EdgeRouterPolicyPatch edge router policy patch
//
// swagger:model edgeRouterPolicyPatch
type EdgeRouterPolicyPatch struct {

	// edge router roles
	EdgeRouterRoles Roles `json:"edgeRouterRoles"`

	// identity roles
	IdentityRoles Roles `json:"identityRoles"`

	// name
	Name string `json:"name,omitempty"`

	// semantic
	Semantic Semantic `json:"semantic,omitempty"`

	// tags
	Tags *Tags `json:"tags,omitempty"`
}

// Validate validates this edge router policy patch
func (m *EdgeRouterPolicyPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeRouterRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemantic(formats); err != nil {
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

func (m *EdgeRouterPolicyPatch) validateEdgeRouterRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeRouterRoles) { // not required
		return nil
	}

	if err := m.EdgeRouterRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRoles")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyPatch) validateIdentityRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentityRoles) { // not required
		return nil
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

func (m *EdgeRouterPolicyPatch) validateSemantic(formats strfmt.Registry) error {
	if swag.IsZero(m.Semantic) { // not required
		return nil
	}

	if err := m.Semantic.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("semantic")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("semantic")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyPatch) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this edge router policy patch based on the context it is used
func (m *EdgeRouterPolicyPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeRouterRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSemantic(ctx, formats); err != nil {
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

func (m *EdgeRouterPolicyPatch) contextValidateEdgeRouterRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EdgeRouterRoles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRoles")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyPatch) contextValidateIdentityRoles(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EdgeRouterPolicyPatch) contextValidateSemantic(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Semantic.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("semantic")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("semantic")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyPatch) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EdgeRouterPolicyPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeRouterPolicyPatch) UnmarshalBinary(b []byte) error {
	var res EdgeRouterPolicyPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

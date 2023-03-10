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

// ServiceEdgeRouterPolicyDetail service edge router policy detail
//
// swagger:model serviceEdgeRouterPolicyDetail
type ServiceEdgeRouterPolicyDetail struct {
	BaseEntity

	// edge router roles
	// Required: true
	EdgeRouterRoles Roles `json:"edgeRouterRoles"`

	// edge router roles display
	// Required: true
	EdgeRouterRolesDisplay NamedRoles `json:"edgeRouterRolesDisplay"`

	// name
	// Required: true
	Name *string `json:"name"`

	// semantic
	// Required: true
	Semantic *Semantic `json:"semantic"`

	// service roles
	// Required: true
	ServiceRoles Roles `json:"serviceRoles"`

	// service roles display
	// Required: true
	ServiceRolesDisplay NamedRoles `json:"serviceRolesDisplay"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServiceEdgeRouterPolicyDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		EdgeRouterRoles Roles `json:"edgeRouterRoles"`

		EdgeRouterRolesDisplay NamedRoles `json:"edgeRouterRolesDisplay"`

		Name *string `json:"name"`

		Semantic *Semantic `json:"semantic"`

		ServiceRoles Roles `json:"serviceRoles"`

		ServiceRolesDisplay NamedRoles `json:"serviceRolesDisplay"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EdgeRouterRoles = dataAO1.EdgeRouterRoles

	m.EdgeRouterRolesDisplay = dataAO1.EdgeRouterRolesDisplay

	m.Name = dataAO1.Name

	m.Semantic = dataAO1.Semantic

	m.ServiceRoles = dataAO1.ServiceRoles

	m.ServiceRolesDisplay = dataAO1.ServiceRolesDisplay

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServiceEdgeRouterPolicyDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		EdgeRouterRoles Roles `json:"edgeRouterRoles"`

		EdgeRouterRolesDisplay NamedRoles `json:"edgeRouterRolesDisplay"`

		Name *string `json:"name"`

		Semantic *Semantic `json:"semantic"`

		ServiceRoles Roles `json:"serviceRoles"`

		ServiceRolesDisplay NamedRoles `json:"serviceRolesDisplay"`
	}

	dataAO1.EdgeRouterRoles = m.EdgeRouterRoles

	dataAO1.EdgeRouterRolesDisplay = m.EdgeRouterRolesDisplay

	dataAO1.Name = m.Name

	dataAO1.Semantic = m.Semantic

	dataAO1.ServiceRoles = m.ServiceRoles

	dataAO1.ServiceRolesDisplay = m.ServiceRolesDisplay

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this service edge router policy detail
func (m *ServiceEdgeRouterPolicyDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeRouterRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeRouterRolesDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemantic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRolesDisplay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateEdgeRouterRoles(formats strfmt.Registry) error {

	if err := validate.Required("edgeRouterRoles", "body", m.EdgeRouterRoles); err != nil {
		return err
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

func (m *ServiceEdgeRouterPolicyDetail) validateEdgeRouterRolesDisplay(formats strfmt.Registry) error {

	if err := validate.Required("edgeRouterRolesDisplay", "body", m.EdgeRouterRolesDisplay); err != nil {
		return err
	}

	if err := m.EdgeRouterRolesDisplay.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRolesDisplay")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateSemantic(formats strfmt.Registry) error {

	if err := validate.Required("semantic", "body", m.Semantic); err != nil {
		return err
	}

	if err := validate.Required("semantic", "body", m.Semantic); err != nil {
		return err
	}

	if m.Semantic != nil {
		if err := m.Semantic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semantic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semantic")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateServiceRoles(formats strfmt.Registry) error {

	if err := validate.Required("serviceRoles", "body", m.ServiceRoles); err != nil {
		return err
	}

	if err := m.ServiceRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceRoles")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateServiceRolesDisplay(formats strfmt.Registry) error {

	if err := validate.Required("serviceRolesDisplay", "body", m.ServiceRolesDisplay); err != nil {
		return err
	}

	if err := m.ServiceRolesDisplay.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceRolesDisplay")
		}
		return err
	}

	return nil
}

// ContextValidate validate this service edge router policy detail based on the context it is used
func (m *ServiceEdgeRouterPolicyDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeRouterRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeRouterRolesDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSemantic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceRolesDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) contextValidateEdgeRouterRoles(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceEdgeRouterPolicyDetail) contextValidateEdgeRouterRolesDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EdgeRouterRolesDisplay.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRolesDisplay")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) contextValidateSemantic(ctx context.Context, formats strfmt.Registry) error {

	if m.Semantic != nil {
		if err := m.Semantic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semantic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semantic")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) contextValidateServiceRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServiceRoles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceRoles")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) contextValidateServiceRolesDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServiceRolesDisplay.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceRolesDisplay")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEdgeRouterPolicyDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEdgeRouterPolicyDetail) UnmarshalBinary(b []byte) error {
	var res ServiceEdgeRouterPolicyDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

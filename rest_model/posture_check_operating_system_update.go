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
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureCheckOperatingSystemUpdate posture check operating system update
//
// swagger:model postureCheckOperatingSystemUpdate
type PostureCheckOperatingSystemUpdate struct {
	nameField *string

	roleAttributesField *Attributes

	tagsField *Tags

	// operating systems
	// Required: true
	// Min Items: 1
	OperatingSystems []*OperatingSystem `json:"operatingSystems"`
}

// Name gets the name of this subtype
func (m *PostureCheckOperatingSystemUpdate) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PostureCheckOperatingSystemUpdate) SetName(val *string) {
	m.nameField = val
}

// RoleAttributes gets the role attributes of this subtype
func (m *PostureCheckOperatingSystemUpdate) RoleAttributes() *Attributes {
	return m.roleAttributesField
}

// SetRoleAttributes sets the role attributes of this subtype
func (m *PostureCheckOperatingSystemUpdate) SetRoleAttributes(val *Attributes) {
	m.roleAttributesField = val
}

// Tags gets the tags of this subtype
func (m *PostureCheckOperatingSystemUpdate) Tags() *Tags {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *PostureCheckOperatingSystemUpdate) SetTags(val *Tags) {
	m.tagsField = val
}

// TypeID gets the type Id of this subtype
func (m *PostureCheckOperatingSystemUpdate) TypeID() PostureCheckType {
	return "OS"
}

// SetTypeID sets the type Id of this subtype
func (m *PostureCheckOperatingSystemUpdate) SetTypeID(val PostureCheckType) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckOperatingSystemUpdate) UnmarshalJSON(raw []byte) error {
	var data struct {

		// operating systems
		// Required: true
		// Min Items: 1
		OperatingSystems []*OperatingSystem `json:"operatingSystems"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name *string `json:"name"`

		RoleAttributes *Attributes `json:"roleAttributes,omitempty"`

		Tags *Tags `json:"tags,omitempty"`

		TypeID PostureCheckType `json:"typeId,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckOperatingSystemUpdate

	result.nameField = base.Name

	result.roleAttributesField = base.RoleAttributes

	result.tagsField = base.Tags

	if base.TypeID != result.TypeID() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid typeId value: %q", base.TypeID)
	}

	result.OperatingSystems = data.OperatingSystems

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckOperatingSystemUpdate) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// operating systems
		// Required: true
		// Min Items: 1
		OperatingSystems []*OperatingSystem `json:"operatingSystems"`
	}{

		OperatingSystems: m.OperatingSystems,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name *string `json:"name"`

		RoleAttributes *Attributes `json:"roleAttributes,omitempty"`

		Tags *Tags `json:"tags,omitempty"`

		TypeID PostureCheckType `json:"typeId,omitempty"`
	}{

		Name: m.Name(),

		RoleAttributes: m.RoleAttributes(),

		Tags: m.Tags(),

		TypeID: m.TypeID(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this posture check operating system update
func (m *PostureCheckOperatingSystemUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckOperatingSystemUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckOperatingSystemUpdate) validateRoleAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleAttributes()) { // not required
		return nil
	}

	if m.RoleAttributes() != nil {
		if err := m.RoleAttributes().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckOperatingSystemUpdate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags()) { // not required
		return nil
	}

	if m.Tags() != nil {
		if err := m.Tags().Validate(formats); err != nil {
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

func (m *PostureCheckOperatingSystemUpdate) validateOperatingSystems(formats strfmt.Registry) error {

	if err := validate.Required("operatingSystems", "body", m.OperatingSystems); err != nil {
		return err
	}

	iOperatingSystemsSize := int64(len(m.OperatingSystems))

	if err := validate.MinItems("operatingSystems", "body", iOperatingSystemsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.OperatingSystems); i++ {
		if swag.IsZero(m.OperatingSystems[i]) { // not required
			continue
		}

		if m.OperatingSystems[i] != nil {
			if err := m.OperatingSystems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operatingSystems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("operatingSystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this posture check operating system update based on the context it is used
func (m *PostureCheckOperatingSystemUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingSystems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckOperatingSystemUpdate) contextValidateRoleAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleAttributes() != nil {
		if err := m.RoleAttributes().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckOperatingSystemUpdate) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if m.Tags() != nil {
		if err := m.Tags().ContextValidate(ctx, formats); err != nil {
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

func (m *PostureCheckOperatingSystemUpdate) contextValidateTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TypeID().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("typeId")
		}
		return err
	}

	return nil
}

func (m *PostureCheckOperatingSystemUpdate) contextValidateOperatingSystems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OperatingSystems); i++ {

		if m.OperatingSystems[i] != nil {
			if err := m.OperatingSystems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operatingSystems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("operatingSystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckOperatingSystemUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckOperatingSystemUpdate) UnmarshalBinary(b []byte) error {
	var res PostureCheckOperatingSystemUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

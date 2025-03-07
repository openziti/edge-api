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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TargetToken Defines the target token type
//
// swagger:model targetToken
type TargetToken string

func NewTargetToken(value TargetToken) *TargetToken {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TargetToken.
func (m TargetToken) Pointer() *TargetToken {
	return &m
}

const (

	// TargetTokenACCESS captures enum value "ACCESS"
	TargetTokenACCESS TargetToken = "ACCESS"

	// TargetTokenID captures enum value "ID"
	TargetTokenID TargetToken = "ID"
)

// for schema
var targetTokenEnum []interface{}

func init() {
	var res []TargetToken
	if err := json.Unmarshal([]byte(`["ACCESS","ID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		targetTokenEnum = append(targetTokenEnum, v)
	}
}

func (m TargetToken) validateTargetTokenEnum(path, location string, value TargetToken) error {
	if err := validate.EnumCase(path, location, value, targetTokenEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this target token
func (m TargetToken) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTargetTokenEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this target token based on context it is used
func (m TargetToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

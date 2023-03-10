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
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PostureCheckFailure posture check failure
//
// swagger:discriminator postureCheckFailure postureCheckType
type PostureCheckFailure interface {
	runtime.Validatable
	runtime.ContextValidatable

	// posture check Id
	// Required: true
	PostureCheckID() *string
	SetPostureCheckID(*string)

	// posture check name
	// Required: true
	PostureCheckName() *string
	SetPostureCheckName(*string)

	// posture check type
	// Required: true
	PostureCheckType() string
	SetPostureCheckType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type postureCheckFailure struct {
	postureCheckIdField *string

	postureCheckNameField *string

	postureCheckTypeField string
}

// PostureCheckID gets the posture check Id of this polymorphic type
func (m *postureCheckFailure) PostureCheckID() *string {
	return m.postureCheckIdField
}

// SetPostureCheckID sets the posture check Id of this polymorphic type
func (m *postureCheckFailure) SetPostureCheckID(val *string) {
	m.postureCheckIdField = val
}

// PostureCheckName gets the posture check name of this polymorphic type
func (m *postureCheckFailure) PostureCheckName() *string {
	return m.postureCheckNameField
}

// SetPostureCheckName sets the posture check name of this polymorphic type
func (m *postureCheckFailure) SetPostureCheckName(val *string) {
	m.postureCheckNameField = val
}

// PostureCheckType gets the posture check type of this polymorphic type
func (m *postureCheckFailure) PostureCheckType() string {
	return "postureCheckFailure"
}

// SetPostureCheckType sets the posture check type of this polymorphic type
func (m *postureCheckFailure) SetPostureCheckType(val string) {
}

// UnmarshalPostureCheckFailureSlice unmarshals polymorphic slices of PostureCheckFailure
func UnmarshalPostureCheckFailureSlice(reader io.Reader, consumer runtime.Consumer) ([]PostureCheckFailure, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PostureCheckFailure
	for _, element := range elements {
		obj, err := unmarshalPostureCheckFailure(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPostureCheckFailure unmarshals polymorphic PostureCheckFailure
func UnmarshalPostureCheckFailure(reader io.Reader, consumer runtime.Consumer) (PostureCheckFailure, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPostureCheckFailure(data, consumer)
}

func unmarshalPostureCheckFailure(data []byte, consumer runtime.Consumer) (PostureCheckFailure, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the postureCheckType property.
	var getType struct {
		PostureCheckType string `json:"postureCheckType"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("postureCheckType", "body", getType.PostureCheckType); err != nil {
		return nil, err
	}

	// The value of postureCheckType is used to determine which type to create and unmarshal the data into
	switch getType.PostureCheckType {
	case "DOMAIN":
		var result PostureCheckFailureDomain
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MAC":
		var result PostureCheckFailureMacAddress
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MFA":
		var result PostureCheckFailureMfa
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "OS":
		var result PostureCheckFailureOperatingSystem
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "PROCESS":
		var result PostureCheckFailureProcess
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "PROCESS_MULTI":
		var result PostureCheckFailureProcessMulti
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "postureCheckFailure":
		var result postureCheckFailure
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid postureCheckType value: %q", getType.PostureCheckType)
}

// Validate validates this posture check failure
func (m *postureCheckFailure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePostureCheckID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostureCheckName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *postureCheckFailure) validatePostureCheckID(formats strfmt.Registry) error {

	if err := validate.Required("postureCheckId", "body", m.PostureCheckID()); err != nil {
		return err
	}

	return nil
}

func (m *postureCheckFailure) validatePostureCheckName(formats strfmt.Registry) error {

	if err := validate.Required("postureCheckName", "body", m.PostureCheckName()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this posture check failure based on context it is used
func (m *postureCheckFailure) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

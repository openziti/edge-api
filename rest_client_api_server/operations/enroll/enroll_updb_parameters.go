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

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewEnrollUpdbParams creates a new EnrollUpdbParams object
//
// There are no default values defined in the spec.
func NewEnrollUpdbParams() EnrollUpdbParams {

	return EnrollUpdbParams{}
}

// EnrollUpdbParams contains all the bound params for the enroll updb operation
// typically these are obtained from a http.Request
//
// swagger:parameters enrollUpdb
type EnrollUpdbParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Token strfmt.UUID
	/*
	  Required: true
	  In: body
	*/
	UpdbCredentials EnrollUpdbBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEnrollUpdbParams() beforehand.
func (o *EnrollUpdbParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qToken, qhkToken, _ := qs.GetOK("token")
	if err := o.bindToken(qToken, qhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body EnrollUpdbBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("updbCredentials", "body", ""))
			} else {
				res = append(res, errors.NewParseError("updbCredentials", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UpdbCredentials = body
			}
		}
	} else {
		res = append(res, errors.Required("updbCredentials", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindToken binds and validates parameter Token from query.
func (o *EnrollUpdbParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("token", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("token", "query", raw); err != nil {
		return err
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("token", "query", "strfmt.UUID", raw)
	}
	o.Token = *(value.(*strfmt.UUID))

	if err := o.validateToken(formats); err != nil {
		return err
	}

	return nil
}

// validateToken carries on validations for parameter Token
func (o *EnrollUpdbParams) validateToken(formats strfmt.Registry) error {

	if err := validate.FormatOf("token", "query", "uuid", o.Token.String(), formats); err != nil {
		return err
	}
	return nil
}

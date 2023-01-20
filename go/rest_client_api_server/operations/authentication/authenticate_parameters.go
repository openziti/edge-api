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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/openziti/edge-api/go/rest_model"
)

// NewAuthenticateParams creates a new AuthenticateParams object
//
// There are no default values defined in the spec.
func NewAuthenticateParams() AuthenticateParams {

	return AuthenticateParams{}
}

// AuthenticateParams contains all the bound params for the authenticate operation
// typically these are obtained from a http.Request
//
// swagger:parameters authenticate
type AuthenticateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Auth *rest_model.Authenticate
	/*
	  Required: true
	  In: query
	*/
	Method string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAuthenticateParams() beforehand.
func (o *AuthenticateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body rest_model.Authenticate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("auth", "body", "", err))
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
				o.Auth = &body
			}
		}
	}

	qMethod, qhkMethod, _ := qs.GetOK("method")
	if err := o.bindMethod(qMethod, qhkMethod, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMethod binds and validates parameter Method from query.
func (o *AuthenticateParams) bindMethod(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("method", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("method", "query", raw); err != nil {
		return err
	}
	o.Method = raw

	if err := o.validateMethod(formats); err != nil {
		return err
	}

	return nil
}

// validateMethod carries on validations for parameter Method
func (o *AuthenticateParams) validateMethod(formats strfmt.Registry) error {

	if err := validate.EnumCase("method", "query", o.Method, []interface{}{"password", "cert", "ext-jwt"}, true); err != nil {
		return err
	}

	return nil
}

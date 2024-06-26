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
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// NewEnrollmentChallengeParams creates a new EnrollmentChallengeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnrollmentChallengeParams() *EnrollmentChallengeParams {
	return &EnrollmentChallengeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnrollmentChallengeParamsWithTimeout creates a new EnrollmentChallengeParams object
// with the ability to set a timeout on a request.
func NewEnrollmentChallengeParamsWithTimeout(timeout time.Duration) *EnrollmentChallengeParams {
	return &EnrollmentChallengeParams{
		timeout: timeout,
	}
}

// NewEnrollmentChallengeParamsWithContext creates a new EnrollmentChallengeParams object
// with the ability to set a context for a request.
func NewEnrollmentChallengeParamsWithContext(ctx context.Context) *EnrollmentChallengeParams {
	return &EnrollmentChallengeParams{
		Context: ctx,
	}
}

// NewEnrollmentChallengeParamsWithHTTPClient creates a new EnrollmentChallengeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnrollmentChallengeParamsWithHTTPClient(client *http.Client) *EnrollmentChallengeParams {
	return &EnrollmentChallengeParams{
		HTTPClient: client,
	}
}

/* EnrollmentChallengeParams contains all the parameters to send to the API endpoint
   for the enrollment challenge operation.

   Typically these are written to a http.Request.
*/
type EnrollmentChallengeParams struct {

	// Nonce.
	Nonce *rest_model.NonceChallenge

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enrollment challenge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnrollmentChallengeParams) WithDefaults() *EnrollmentChallengeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enrollment challenge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnrollmentChallengeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enrollment challenge params
func (o *EnrollmentChallengeParams) WithTimeout(timeout time.Duration) *EnrollmentChallengeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enrollment challenge params
func (o *EnrollmentChallengeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enrollment challenge params
func (o *EnrollmentChallengeParams) WithContext(ctx context.Context) *EnrollmentChallengeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enrollment challenge params
func (o *EnrollmentChallengeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enrollment challenge params
func (o *EnrollmentChallengeParams) WithHTTPClient(client *http.Client) *EnrollmentChallengeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enrollment challenge params
func (o *EnrollmentChallengeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNonce adds the nonce to the enrollment challenge params
func (o *EnrollmentChallengeParams) WithNonce(nonce *rest_model.NonceChallenge) *EnrollmentChallengeParams {
	o.SetNonce(nonce)
	return o
}

// SetNonce adds the nonce to the enrollment challenge params
func (o *EnrollmentChallengeParams) SetNonce(nonce *rest_model.NonceChallenge) {
	o.Nonce = nonce
}

// WriteToRequest writes these params to a swagger request
func (o *EnrollmentChallengeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Nonce != nil {
		if err := r.SetBodyParam(o.Nonce); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

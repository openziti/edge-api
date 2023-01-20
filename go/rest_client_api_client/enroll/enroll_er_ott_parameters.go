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
)

// NewEnrollErOttParams creates a new EnrollErOttParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnrollErOttParams() *EnrollErOttParams {
	return &EnrollErOttParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnrollErOttParamsWithTimeout creates a new EnrollErOttParams object
// with the ability to set a timeout on a request.
func NewEnrollErOttParamsWithTimeout(timeout time.Duration) *EnrollErOttParams {
	return &EnrollErOttParams{
		timeout: timeout,
	}
}

// NewEnrollErOttParamsWithContext creates a new EnrollErOttParams object
// with the ability to set a context for a request.
func NewEnrollErOttParamsWithContext(ctx context.Context) *EnrollErOttParams {
	return &EnrollErOttParams{
		Context: ctx,
	}
}

// NewEnrollErOttParamsWithHTTPClient creates a new EnrollErOttParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnrollErOttParamsWithHTTPClient(client *http.Client) *EnrollErOttParams {
	return &EnrollErOttParams{
		HTTPClient: client,
	}
}

/* EnrollErOttParams contains all the parameters to send to the API endpoint
   for the enroll er ott operation.

   Typically these are written to a http.Request.
*/
type EnrollErOttParams struct {

	// Token.
	//
	// Format: uuid
	Token strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enroll er ott params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnrollErOttParams) WithDefaults() *EnrollErOttParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enroll er ott params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnrollErOttParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enroll er ott params
func (o *EnrollErOttParams) WithTimeout(timeout time.Duration) *EnrollErOttParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enroll er ott params
func (o *EnrollErOttParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enroll er ott params
func (o *EnrollErOttParams) WithContext(ctx context.Context) *EnrollErOttParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enroll er ott params
func (o *EnrollErOttParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enroll er ott params
func (o *EnrollErOttParams) WithHTTPClient(client *http.Client) *EnrollErOttParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enroll er ott params
func (o *EnrollErOttParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the enroll er ott params
func (o *EnrollErOttParams) WithToken(token strfmt.UUID) *EnrollErOttParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the enroll er ott params
func (o *EnrollErOttParams) SetToken(token strfmt.UUID) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *EnrollErOttParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param token
	qrToken := o.Token
	qToken := qrToken.String()
	if qToken != "" {

		if err := r.SetQueryParam("token", qToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

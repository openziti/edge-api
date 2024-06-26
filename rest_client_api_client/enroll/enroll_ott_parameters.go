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

// NewEnrollOttParams creates a new EnrollOttParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnrollOttParams() *EnrollOttParams {
	return &EnrollOttParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnrollOttParamsWithTimeout creates a new EnrollOttParams object
// with the ability to set a timeout on a request.
func NewEnrollOttParamsWithTimeout(timeout time.Duration) *EnrollOttParams {
	return &EnrollOttParams{
		timeout: timeout,
	}
}

// NewEnrollOttParamsWithContext creates a new EnrollOttParams object
// with the ability to set a context for a request.
func NewEnrollOttParamsWithContext(ctx context.Context) *EnrollOttParams {
	return &EnrollOttParams{
		Context: ctx,
	}
}

// NewEnrollOttParamsWithHTTPClient creates a new EnrollOttParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnrollOttParamsWithHTTPClient(client *http.Client) *EnrollOttParams {
	return &EnrollOttParams{
		HTTPClient: client,
	}
}

/* EnrollOttParams contains all the parameters to send to the API endpoint
   for the enroll ott operation.

   Typically these are written to a http.Request.
*/
type EnrollOttParams struct {

	/* OttEnrollmentRequest.

	   An OTT enrollment request
	*/
	OttEnrollmentRequest *rest_model.OttEnrollmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enroll ott params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnrollOttParams) WithDefaults() *EnrollOttParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enroll ott params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnrollOttParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enroll ott params
func (o *EnrollOttParams) WithTimeout(timeout time.Duration) *EnrollOttParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enroll ott params
func (o *EnrollOttParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enroll ott params
func (o *EnrollOttParams) WithContext(ctx context.Context) *EnrollOttParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enroll ott params
func (o *EnrollOttParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enroll ott params
func (o *EnrollOttParams) WithHTTPClient(client *http.Client) *EnrollOttParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enroll ott params
func (o *EnrollOttParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOttEnrollmentRequest adds the ottEnrollmentRequest to the enroll ott params
func (o *EnrollOttParams) WithOttEnrollmentRequest(ottEnrollmentRequest *rest_model.OttEnrollmentRequest) *EnrollOttParams {
	o.SetOttEnrollmentRequest(ottEnrollmentRequest)
	return o
}

// SetOttEnrollmentRequest adds the ottEnrollmentRequest to the enroll ott params
func (o *EnrollOttParams) SetOttEnrollmentRequest(ottEnrollmentRequest *rest_model.OttEnrollmentRequest) {
	o.OttEnrollmentRequest = ottEnrollmentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *EnrollOttParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.OttEnrollmentRequest != nil {
		if err := r.SetBodyParam(o.OttEnrollmentRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

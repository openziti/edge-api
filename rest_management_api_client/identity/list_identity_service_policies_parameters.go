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

package identity

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

// NewListIdentityServicePoliciesParams creates a new ListIdentityServicePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListIdentityServicePoliciesParams() *ListIdentityServicePoliciesParams {
	return &ListIdentityServicePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListIdentityServicePoliciesParamsWithTimeout creates a new ListIdentityServicePoliciesParams object
// with the ability to set a timeout on a request.
func NewListIdentityServicePoliciesParamsWithTimeout(timeout time.Duration) *ListIdentityServicePoliciesParams {
	return &ListIdentityServicePoliciesParams{
		timeout: timeout,
	}
}

// NewListIdentityServicePoliciesParamsWithContext creates a new ListIdentityServicePoliciesParams object
// with the ability to set a context for a request.
func NewListIdentityServicePoliciesParamsWithContext(ctx context.Context) *ListIdentityServicePoliciesParams {
	return &ListIdentityServicePoliciesParams{
		Context: ctx,
	}
}

// NewListIdentityServicePoliciesParamsWithHTTPClient creates a new ListIdentityServicePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListIdentityServicePoliciesParamsWithHTTPClient(client *http.Client) *ListIdentityServicePoliciesParams {
	return &ListIdentityServicePoliciesParams{
		HTTPClient: client,
	}
}

/* ListIdentityServicePoliciesParams contains all the parameters to send to the API endpoint
   for the list identity service policies operation.

   Typically these are written to a http.Request.
*/
type ListIdentityServicePoliciesParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list identity service policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIdentityServicePoliciesParams) WithDefaults() *ListIdentityServicePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list identity service policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIdentityServicePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) WithTimeout(timeout time.Duration) *ListIdentityServicePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) WithContext(ctx context.Context) *ListIdentityServicePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) WithHTTPClient(client *http.Client) *ListIdentityServicePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) WithID(id string) *ListIdentityServicePoliciesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list identity service policies params
func (o *ListIdentityServicePoliciesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListIdentityServicePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

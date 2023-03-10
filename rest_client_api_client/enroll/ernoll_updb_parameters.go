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

// NewErnollUpdbParams creates a new ErnollUpdbParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewErnollUpdbParams() *ErnollUpdbParams {
	return &ErnollUpdbParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewErnollUpdbParamsWithTimeout creates a new ErnollUpdbParams object
// with the ability to set a timeout on a request.
func NewErnollUpdbParamsWithTimeout(timeout time.Duration) *ErnollUpdbParams {
	return &ErnollUpdbParams{
		timeout: timeout,
	}
}

// NewErnollUpdbParamsWithContext creates a new ErnollUpdbParams object
// with the ability to set a context for a request.
func NewErnollUpdbParamsWithContext(ctx context.Context) *ErnollUpdbParams {
	return &ErnollUpdbParams{
		Context: ctx,
	}
}

// NewErnollUpdbParamsWithHTTPClient creates a new ErnollUpdbParams object
// with the ability to set a custom HTTPClient for a request.
func NewErnollUpdbParamsWithHTTPClient(client *http.Client) *ErnollUpdbParams {
	return &ErnollUpdbParams{
		HTTPClient: client,
	}
}

/* ErnollUpdbParams contains all the parameters to send to the API endpoint
   for the ernoll updb operation.

   Typically these are written to a http.Request.
*/
type ErnollUpdbParams struct {

	// Token.
	//
	// Format: uuid
	Token strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ernoll updb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErnollUpdbParams) WithDefaults() *ErnollUpdbParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ernoll updb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErnollUpdbParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ernoll updb params
func (o *ErnollUpdbParams) WithTimeout(timeout time.Duration) *ErnollUpdbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ernoll updb params
func (o *ErnollUpdbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ernoll updb params
func (o *ErnollUpdbParams) WithContext(ctx context.Context) *ErnollUpdbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ernoll updb params
func (o *ErnollUpdbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ernoll updb params
func (o *ErnollUpdbParams) WithHTTPClient(client *http.Client) *ErnollUpdbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ernoll updb params
func (o *ErnollUpdbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the ernoll updb params
func (o *ErnollUpdbParams) WithToken(token strfmt.UUID) *ErnollUpdbParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the ernoll updb params
func (o *ErnollUpdbParams) SetToken(token strfmt.UUID) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *ErnollUpdbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

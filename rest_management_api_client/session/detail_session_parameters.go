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

package session

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

// NewDetailSessionParams creates a new DetailSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetailSessionParams() *DetailSessionParams {
	return &DetailSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetailSessionParamsWithTimeout creates a new DetailSessionParams object
// with the ability to set a timeout on a request.
func NewDetailSessionParamsWithTimeout(timeout time.Duration) *DetailSessionParams {
	return &DetailSessionParams{
		timeout: timeout,
	}
}

// NewDetailSessionParamsWithContext creates a new DetailSessionParams object
// with the ability to set a context for a request.
func NewDetailSessionParamsWithContext(ctx context.Context) *DetailSessionParams {
	return &DetailSessionParams{
		Context: ctx,
	}
}

// NewDetailSessionParamsWithHTTPClient creates a new DetailSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetailSessionParamsWithHTTPClient(client *http.Client) *DetailSessionParams {
	return &DetailSessionParams{
		HTTPClient: client,
	}
}

/* DetailSessionParams contains all the parameters to send to the API endpoint
   for the detail session operation.

   Typically these are written to a http.Request.
*/
type DetailSessionParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detail session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailSessionParams) WithDefaults() *DetailSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detail session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detail session params
func (o *DetailSessionParams) WithTimeout(timeout time.Duration) *DetailSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail session params
func (o *DetailSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail session params
func (o *DetailSessionParams) WithContext(ctx context.Context) *DetailSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail session params
func (o *DetailSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail session params
func (o *DetailSessionParams) WithHTTPClient(client *http.Client) *DetailSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail session params
func (o *DetailSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the detail session params
func (o *DetailSessionParams) WithID(id string) *DetailSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the detail session params
func (o *DetailSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DetailSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

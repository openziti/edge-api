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

package router

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

// NewDetailRouterParams creates a new DetailRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetailRouterParams() *DetailRouterParams {
	return &DetailRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetailRouterParamsWithTimeout creates a new DetailRouterParams object
// with the ability to set a timeout on a request.
func NewDetailRouterParamsWithTimeout(timeout time.Duration) *DetailRouterParams {
	return &DetailRouterParams{
		timeout: timeout,
	}
}

// NewDetailRouterParamsWithContext creates a new DetailRouterParams object
// with the ability to set a context for a request.
func NewDetailRouterParamsWithContext(ctx context.Context) *DetailRouterParams {
	return &DetailRouterParams{
		Context: ctx,
	}
}

// NewDetailRouterParamsWithHTTPClient creates a new DetailRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetailRouterParamsWithHTTPClient(client *http.Client) *DetailRouterParams {
	return &DetailRouterParams{
		HTTPClient: client,
	}
}

/* DetailRouterParams contains all the parameters to send to the API endpoint
   for the detail router operation.

   Typically these are written to a http.Request.
*/
type DetailRouterParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detail router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailRouterParams) WithDefaults() *DetailRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detail router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detail router params
func (o *DetailRouterParams) WithTimeout(timeout time.Duration) *DetailRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail router params
func (o *DetailRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail router params
func (o *DetailRouterParams) WithContext(ctx context.Context) *DetailRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail router params
func (o *DetailRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail router params
func (o *DetailRouterParams) WithHTTPClient(client *http.Client) *DetailRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail router params
func (o *DetailRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the detail router params
func (o *DetailRouterParams) WithID(id string) *DetailRouterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the detail router params
func (o *DetailRouterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DetailRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

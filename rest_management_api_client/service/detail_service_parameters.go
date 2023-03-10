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

package service

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

// NewDetailServiceParams creates a new DetailServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetailServiceParams() *DetailServiceParams {
	return &DetailServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetailServiceParamsWithTimeout creates a new DetailServiceParams object
// with the ability to set a timeout on a request.
func NewDetailServiceParamsWithTimeout(timeout time.Duration) *DetailServiceParams {
	return &DetailServiceParams{
		timeout: timeout,
	}
}

// NewDetailServiceParamsWithContext creates a new DetailServiceParams object
// with the ability to set a context for a request.
func NewDetailServiceParamsWithContext(ctx context.Context) *DetailServiceParams {
	return &DetailServiceParams{
		Context: ctx,
	}
}

// NewDetailServiceParamsWithHTTPClient creates a new DetailServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetailServiceParamsWithHTTPClient(client *http.Client) *DetailServiceParams {
	return &DetailServiceParams{
		HTTPClient: client,
	}
}

/* DetailServiceParams contains all the parameters to send to the API endpoint
   for the detail service operation.

   Typically these are written to a http.Request.
*/
type DetailServiceParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detail service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailServiceParams) WithDefaults() *DetailServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detail service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detail service params
func (o *DetailServiceParams) WithTimeout(timeout time.Duration) *DetailServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail service params
func (o *DetailServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail service params
func (o *DetailServiceParams) WithContext(ctx context.Context) *DetailServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail service params
func (o *DetailServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail service params
func (o *DetailServiceParams) WithHTTPClient(client *http.Client) *DetailServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail service params
func (o *DetailServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the detail service params
func (o *DetailServiceParams) WithID(id string) *DetailServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the detail service params
func (o *DetailServiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DetailServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

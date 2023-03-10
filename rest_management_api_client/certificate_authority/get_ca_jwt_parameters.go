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

package certificate_authority

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

// NewGetCaJWTParams creates a new GetCaJWTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCaJWTParams() *GetCaJWTParams {
	return &GetCaJWTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCaJWTParamsWithTimeout creates a new GetCaJWTParams object
// with the ability to set a timeout on a request.
func NewGetCaJWTParamsWithTimeout(timeout time.Duration) *GetCaJWTParams {
	return &GetCaJWTParams{
		timeout: timeout,
	}
}

// NewGetCaJWTParamsWithContext creates a new GetCaJWTParams object
// with the ability to set a context for a request.
func NewGetCaJWTParamsWithContext(ctx context.Context) *GetCaJWTParams {
	return &GetCaJWTParams{
		Context: ctx,
	}
}

// NewGetCaJWTParamsWithHTTPClient creates a new GetCaJWTParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCaJWTParamsWithHTTPClient(client *http.Client) *GetCaJWTParams {
	return &GetCaJWTParams{
		HTTPClient: client,
	}
}

/* GetCaJWTParams contains all the parameters to send to the API endpoint
   for the get ca Jwt operation.

   Typically these are written to a http.Request.
*/
type GetCaJWTParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ca Jwt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCaJWTParams) WithDefaults() *GetCaJWTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ca Jwt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCaJWTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ca Jwt params
func (o *GetCaJWTParams) WithTimeout(timeout time.Duration) *GetCaJWTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ca Jwt params
func (o *GetCaJWTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ca Jwt params
func (o *GetCaJWTParams) WithContext(ctx context.Context) *GetCaJWTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ca Jwt params
func (o *GetCaJWTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ca Jwt params
func (o *GetCaJWTParams) WithHTTPClient(client *http.Client) *GetCaJWTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ca Jwt params
func (o *GetCaJWTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get ca Jwt params
func (o *GetCaJWTParams) WithID(id string) *GetCaJWTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ca Jwt params
func (o *GetCaJWTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCaJWTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

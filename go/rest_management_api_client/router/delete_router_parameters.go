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

// NewDeleteRouterParams creates a new DeleteRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRouterParams() *DeleteRouterParams {
	return &DeleteRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRouterParamsWithTimeout creates a new DeleteRouterParams object
// with the ability to set a timeout on a request.
func NewDeleteRouterParamsWithTimeout(timeout time.Duration) *DeleteRouterParams {
	return &DeleteRouterParams{
		timeout: timeout,
	}
}

// NewDeleteRouterParamsWithContext creates a new DeleteRouterParams object
// with the ability to set a context for a request.
func NewDeleteRouterParamsWithContext(ctx context.Context) *DeleteRouterParams {
	return &DeleteRouterParams{
		Context: ctx,
	}
}

// NewDeleteRouterParamsWithHTTPClient creates a new DeleteRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRouterParamsWithHTTPClient(client *http.Client) *DeleteRouterParams {
	return &DeleteRouterParams{
		HTTPClient: client,
	}
}

/* DeleteRouterParams contains all the parameters to send to the API endpoint
   for the delete router operation.

   Typically these are written to a http.Request.
*/
type DeleteRouterParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRouterParams) WithDefaults() *DeleteRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete router params
func (o *DeleteRouterParams) WithTimeout(timeout time.Duration) *DeleteRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete router params
func (o *DeleteRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete router params
func (o *DeleteRouterParams) WithContext(ctx context.Context) *DeleteRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete router params
func (o *DeleteRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete router params
func (o *DeleteRouterParams) WithHTTPClient(client *http.Client) *DeleteRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete router params
func (o *DeleteRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete router params
func (o *DeleteRouterParams) WithID(id string) *DeleteRouterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete router params
func (o *DeleteRouterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

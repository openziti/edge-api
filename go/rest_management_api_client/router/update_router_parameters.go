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

	"github.com/openziti/edge-api/go/rest_model"
)

// NewUpdateRouterParams creates a new UpdateRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRouterParams() *UpdateRouterParams {
	return &UpdateRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRouterParamsWithTimeout creates a new UpdateRouterParams object
// with the ability to set a timeout on a request.
func NewUpdateRouterParamsWithTimeout(timeout time.Duration) *UpdateRouterParams {
	return &UpdateRouterParams{
		timeout: timeout,
	}
}

// NewUpdateRouterParamsWithContext creates a new UpdateRouterParams object
// with the ability to set a context for a request.
func NewUpdateRouterParamsWithContext(ctx context.Context) *UpdateRouterParams {
	return &UpdateRouterParams{
		Context: ctx,
	}
}

// NewUpdateRouterParamsWithHTTPClient creates a new UpdateRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRouterParamsWithHTTPClient(client *http.Client) *UpdateRouterParams {
	return &UpdateRouterParams{
		HTTPClient: client,
	}
}

/* UpdateRouterParams contains all the parameters to send to the API endpoint
   for the update router operation.

   Typically these are written to a http.Request.
*/
type UpdateRouterParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* Router.

	   A router update object
	*/
	Router *rest_model.RouterUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRouterParams) WithDefaults() *UpdateRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update router params
func (o *UpdateRouterParams) WithTimeout(timeout time.Duration) *UpdateRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update router params
func (o *UpdateRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update router params
func (o *UpdateRouterParams) WithContext(ctx context.Context) *UpdateRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update router params
func (o *UpdateRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update router params
func (o *UpdateRouterParams) WithHTTPClient(client *http.Client) *UpdateRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update router params
func (o *UpdateRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update router params
func (o *UpdateRouterParams) WithID(id string) *UpdateRouterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update router params
func (o *UpdateRouterParams) SetID(id string) {
	o.ID = id
}

// WithRouter adds the router to the update router params
func (o *UpdateRouterParams) WithRouter(router *rest_model.RouterUpdate) *UpdateRouterParams {
	o.SetRouter(router)
	return o
}

// SetRouter adds the router to the update router params
func (o *UpdateRouterParams) SetRouter(router *rest_model.RouterUpdate) {
	o.Router = router
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Router != nil {
		if err := r.SetBodyParam(o.Router); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

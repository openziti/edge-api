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

// NewCreateTransitRouterParams creates a new CreateTransitRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTransitRouterParams() *CreateTransitRouterParams {
	return &CreateTransitRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTransitRouterParamsWithTimeout creates a new CreateTransitRouterParams object
// with the ability to set a timeout on a request.
func NewCreateTransitRouterParamsWithTimeout(timeout time.Duration) *CreateTransitRouterParams {
	return &CreateTransitRouterParams{
		timeout: timeout,
	}
}

// NewCreateTransitRouterParamsWithContext creates a new CreateTransitRouterParams object
// with the ability to set a context for a request.
func NewCreateTransitRouterParamsWithContext(ctx context.Context) *CreateTransitRouterParams {
	return &CreateTransitRouterParams{
		Context: ctx,
	}
}

// NewCreateTransitRouterParamsWithHTTPClient creates a new CreateTransitRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTransitRouterParamsWithHTTPClient(client *http.Client) *CreateTransitRouterParams {
	return &CreateTransitRouterParams{
		HTTPClient: client,
	}
}

/* CreateTransitRouterParams contains all the parameters to send to the API endpoint
   for the create transit router operation.

   Typically these are written to a http.Request.
*/
type CreateTransitRouterParams struct {

	/* Router.

	   A router to create
	*/
	Router *rest_model.RouterCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create transit router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTransitRouterParams) WithDefaults() *CreateTransitRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create transit router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTransitRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create transit router params
func (o *CreateTransitRouterParams) WithTimeout(timeout time.Duration) *CreateTransitRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create transit router params
func (o *CreateTransitRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create transit router params
func (o *CreateTransitRouterParams) WithContext(ctx context.Context) *CreateTransitRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create transit router params
func (o *CreateTransitRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create transit router params
func (o *CreateTransitRouterParams) WithHTTPClient(client *http.Client) *CreateTransitRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create transit router params
func (o *CreateTransitRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRouter adds the router to the create transit router params
func (o *CreateTransitRouterParams) WithRouter(router *rest_model.RouterCreate) *CreateTransitRouterParams {
	o.SetRouter(router)
	return o
}

// SetRouter adds the router to the create transit router params
func (o *CreateTransitRouterParams) SetRouter(router *rest_model.RouterCreate) {
	o.Router = router
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTransitRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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

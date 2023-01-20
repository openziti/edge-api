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

package service_edge_router_policy

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

// NewUpdateServiceEdgeRouterPolicyParams creates a new UpdateServiceEdgeRouterPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServiceEdgeRouterPolicyParams() *UpdateServiceEdgeRouterPolicyParams {
	return &UpdateServiceEdgeRouterPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceEdgeRouterPolicyParamsWithTimeout creates a new UpdateServiceEdgeRouterPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateServiceEdgeRouterPolicyParamsWithTimeout(timeout time.Duration) *UpdateServiceEdgeRouterPolicyParams {
	return &UpdateServiceEdgeRouterPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateServiceEdgeRouterPolicyParamsWithContext creates a new UpdateServiceEdgeRouterPolicyParams object
// with the ability to set a context for a request.
func NewUpdateServiceEdgeRouterPolicyParamsWithContext(ctx context.Context) *UpdateServiceEdgeRouterPolicyParams {
	return &UpdateServiceEdgeRouterPolicyParams{
		Context: ctx,
	}
}

// NewUpdateServiceEdgeRouterPolicyParamsWithHTTPClient creates a new UpdateServiceEdgeRouterPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServiceEdgeRouterPolicyParamsWithHTTPClient(client *http.Client) *UpdateServiceEdgeRouterPolicyParams {
	return &UpdateServiceEdgeRouterPolicyParams{
		HTTPClient: client,
	}
}

/* UpdateServiceEdgeRouterPolicyParams contains all the parameters to send to the API endpoint
   for the update service edge router policy operation.

   Typically these are written to a http.Request.
*/
type UpdateServiceEdgeRouterPolicyParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* Policy.

	   A service edge router policy update object
	*/
	Policy *rest_model.ServiceEdgeRouterPolicyUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update service edge router policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceEdgeRouterPolicyParams) WithDefaults() *UpdateServiceEdgeRouterPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update service edge router policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceEdgeRouterPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) WithTimeout(timeout time.Duration) *UpdateServiceEdgeRouterPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) WithContext(ctx context.Context) *UpdateServiceEdgeRouterPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) WithHTTPClient(client *http.Client) *UpdateServiceEdgeRouterPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) WithID(id string) *UpdateServiceEdgeRouterPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) SetID(id string) {
	o.ID = id
}

// WithPolicy adds the policy to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) WithPolicy(policy *rest_model.ServiceEdgeRouterPolicyUpdate) *UpdateServiceEdgeRouterPolicyParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the update service edge router policy params
func (o *UpdateServiceEdgeRouterPolicyParams) SetPolicy(policy *rest_model.ServiceEdgeRouterPolicyUpdate) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceEdgeRouterPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

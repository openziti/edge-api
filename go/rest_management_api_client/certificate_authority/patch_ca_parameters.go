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

	"github.com/openziti/edge-api/go/rest_model"
)

// NewPatchCaParams creates a new PatchCaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchCaParams() *PatchCaParams {
	return &PatchCaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCaParamsWithTimeout creates a new PatchCaParams object
// with the ability to set a timeout on a request.
func NewPatchCaParamsWithTimeout(timeout time.Duration) *PatchCaParams {
	return &PatchCaParams{
		timeout: timeout,
	}
}

// NewPatchCaParamsWithContext creates a new PatchCaParams object
// with the ability to set a context for a request.
func NewPatchCaParamsWithContext(ctx context.Context) *PatchCaParams {
	return &PatchCaParams{
		Context: ctx,
	}
}

// NewPatchCaParamsWithHTTPClient creates a new PatchCaParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchCaParamsWithHTTPClient(client *http.Client) *PatchCaParams {
	return &PatchCaParams{
		HTTPClient: client,
	}
}

/* PatchCaParams contains all the parameters to send to the API endpoint
   for the patch ca operation.

   Typically these are written to a http.Request.
*/
type PatchCaParams struct {

	/* Ca.

	   A CA patch object
	*/
	Ca *rest_model.CaPatch

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch ca params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCaParams) WithDefaults() *PatchCaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch ca params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch ca params
func (o *PatchCaParams) WithTimeout(timeout time.Duration) *PatchCaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch ca params
func (o *PatchCaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch ca params
func (o *PatchCaParams) WithContext(ctx context.Context) *PatchCaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch ca params
func (o *PatchCaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch ca params
func (o *PatchCaParams) WithHTTPClient(client *http.Client) *PatchCaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch ca params
func (o *PatchCaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCa adds the ca to the patch ca params
func (o *PatchCaParams) WithCa(ca *rest_model.CaPatch) *PatchCaParams {
	o.SetCa(ca)
	return o
}

// SetCa adds the ca to the patch ca params
func (o *PatchCaParams) SetCa(ca *rest_model.CaPatch) {
	o.Ca = ca
}

// WithID adds the id to the patch ca params
func (o *PatchCaParams) WithID(id string) *PatchCaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch ca params
func (o *PatchCaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Ca != nil {
		if err := r.SetBodyParam(o.Ca); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

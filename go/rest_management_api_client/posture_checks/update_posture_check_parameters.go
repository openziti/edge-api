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

package posture_checks

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

// NewUpdatePostureCheckParams creates a new UpdatePostureCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePostureCheckParams() *UpdatePostureCheckParams {
	return &UpdatePostureCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePostureCheckParamsWithTimeout creates a new UpdatePostureCheckParams object
// with the ability to set a timeout on a request.
func NewUpdatePostureCheckParamsWithTimeout(timeout time.Duration) *UpdatePostureCheckParams {
	return &UpdatePostureCheckParams{
		timeout: timeout,
	}
}

// NewUpdatePostureCheckParamsWithContext creates a new UpdatePostureCheckParams object
// with the ability to set a context for a request.
func NewUpdatePostureCheckParamsWithContext(ctx context.Context) *UpdatePostureCheckParams {
	return &UpdatePostureCheckParams{
		Context: ctx,
	}
}

// NewUpdatePostureCheckParamsWithHTTPClient creates a new UpdatePostureCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePostureCheckParamsWithHTTPClient(client *http.Client) *UpdatePostureCheckParams {
	return &UpdatePostureCheckParams{
		HTTPClient: client,
	}
}

/* UpdatePostureCheckParams contains all the parameters to send to the API endpoint
   for the update posture check operation.

   Typically these are written to a http.Request.
*/
type UpdatePostureCheckParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* PostureCheck.

	   A Posture Check update object
	*/
	PostureCheck rest_model.PostureCheckUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update posture check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePostureCheckParams) WithDefaults() *UpdatePostureCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update posture check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePostureCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update posture check params
func (o *UpdatePostureCheckParams) WithTimeout(timeout time.Duration) *UpdatePostureCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update posture check params
func (o *UpdatePostureCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update posture check params
func (o *UpdatePostureCheckParams) WithContext(ctx context.Context) *UpdatePostureCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update posture check params
func (o *UpdatePostureCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update posture check params
func (o *UpdatePostureCheckParams) WithHTTPClient(client *http.Client) *UpdatePostureCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update posture check params
func (o *UpdatePostureCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update posture check params
func (o *UpdatePostureCheckParams) WithID(id string) *UpdatePostureCheckParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update posture check params
func (o *UpdatePostureCheckParams) SetID(id string) {
	o.ID = id
}

// WithPostureCheck adds the postureCheck to the update posture check params
func (o *UpdatePostureCheckParams) WithPostureCheck(postureCheck rest_model.PostureCheckUpdate) *UpdatePostureCheckParams {
	o.SetPostureCheck(postureCheck)
	return o
}

// SetPostureCheck adds the postureCheck to the update posture check params
func (o *UpdatePostureCheckParams) SetPostureCheck(postureCheck rest_model.PostureCheckUpdate) {
	o.PostureCheck = postureCheck
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePostureCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.PostureCheck); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

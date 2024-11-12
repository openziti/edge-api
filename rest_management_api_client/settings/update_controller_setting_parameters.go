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

package settings

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

	"github.com/openziti/edge-api/rest_model"
)

// NewUpdateControllerSettingParams creates a new UpdateControllerSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateControllerSettingParams() *UpdateControllerSettingParams {
	return &UpdateControllerSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateControllerSettingParamsWithTimeout creates a new UpdateControllerSettingParams object
// with the ability to set a timeout on a request.
func NewUpdateControllerSettingParamsWithTimeout(timeout time.Duration) *UpdateControllerSettingParams {
	return &UpdateControllerSettingParams{
		timeout: timeout,
	}
}

// NewUpdateControllerSettingParamsWithContext creates a new UpdateControllerSettingParams object
// with the ability to set a context for a request.
func NewUpdateControllerSettingParamsWithContext(ctx context.Context) *UpdateControllerSettingParams {
	return &UpdateControllerSettingParams{
		Context: ctx,
	}
}

// NewUpdateControllerSettingParamsWithHTTPClient creates a new UpdateControllerSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateControllerSettingParamsWithHTTPClient(client *http.Client) *UpdateControllerSettingParams {
	return &UpdateControllerSettingParams{
		HTTPClient: client,
	}
}

/* UpdateControllerSettingParams contains all the parameters to send to the API endpoint
   for the update controller setting operation.

   Typically these are written to a http.Request.
*/
type UpdateControllerSettingParams struct {

	/* ControllerSetting.

	   A controller setting update object
	*/
	ControllerSetting *rest_model.ControllerSettingUpdate

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update controller setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateControllerSettingParams) WithDefaults() *UpdateControllerSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update controller setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateControllerSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update controller setting params
func (o *UpdateControllerSettingParams) WithTimeout(timeout time.Duration) *UpdateControllerSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update controller setting params
func (o *UpdateControllerSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update controller setting params
func (o *UpdateControllerSettingParams) WithContext(ctx context.Context) *UpdateControllerSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update controller setting params
func (o *UpdateControllerSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update controller setting params
func (o *UpdateControllerSettingParams) WithHTTPClient(client *http.Client) *UpdateControllerSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update controller setting params
func (o *UpdateControllerSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithControllerSetting adds the controllerSetting to the update controller setting params
func (o *UpdateControllerSettingParams) WithControllerSetting(controllerSetting *rest_model.ControllerSettingUpdate) *UpdateControllerSettingParams {
	o.SetControllerSetting(controllerSetting)
	return o
}

// SetControllerSetting adds the controllerSetting to the update controller setting params
func (o *UpdateControllerSettingParams) SetControllerSetting(controllerSetting *rest_model.ControllerSettingUpdate) {
	o.ControllerSetting = controllerSetting
}

// WithID adds the id to the update controller setting params
func (o *UpdateControllerSettingParams) WithID(id string) *UpdateControllerSettingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update controller setting params
func (o *UpdateControllerSettingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateControllerSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ControllerSetting != nil {
		if err := r.SetBodyParam(o.ControllerSetting); err != nil {
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

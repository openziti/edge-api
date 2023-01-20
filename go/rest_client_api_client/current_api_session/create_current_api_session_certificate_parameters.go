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

package current_api_session

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

// NewCreateCurrentAPISessionCertificateParams creates a new CreateCurrentAPISessionCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCurrentAPISessionCertificateParams() *CreateCurrentAPISessionCertificateParams {
	return &CreateCurrentAPISessionCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCurrentAPISessionCertificateParamsWithTimeout creates a new CreateCurrentAPISessionCertificateParams object
// with the ability to set a timeout on a request.
func NewCreateCurrentAPISessionCertificateParamsWithTimeout(timeout time.Duration) *CreateCurrentAPISessionCertificateParams {
	return &CreateCurrentAPISessionCertificateParams{
		timeout: timeout,
	}
}

// NewCreateCurrentAPISessionCertificateParamsWithContext creates a new CreateCurrentAPISessionCertificateParams object
// with the ability to set a context for a request.
func NewCreateCurrentAPISessionCertificateParamsWithContext(ctx context.Context) *CreateCurrentAPISessionCertificateParams {
	return &CreateCurrentAPISessionCertificateParams{
		Context: ctx,
	}
}

// NewCreateCurrentAPISessionCertificateParamsWithHTTPClient creates a new CreateCurrentAPISessionCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCurrentAPISessionCertificateParamsWithHTTPClient(client *http.Client) *CreateCurrentAPISessionCertificateParams {
	return &CreateCurrentAPISessionCertificateParams{
		HTTPClient: client,
	}
}

/* CreateCurrentAPISessionCertificateParams contains all the parameters to send to the API endpoint
   for the create current Api session certificate operation.

   Typically these are written to a http.Request.
*/
type CreateCurrentAPISessionCertificateParams struct {

	/* SessionCertificate.

	   The payload describing the CSR used to create a session certificate
	*/
	SessionCertificate *rest_model.CurrentAPISessionCertificateCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create current Api session certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCurrentAPISessionCertificateParams) WithDefaults() *CreateCurrentAPISessionCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create current Api session certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCurrentAPISessionCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) WithTimeout(timeout time.Duration) *CreateCurrentAPISessionCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) WithContext(ctx context.Context) *CreateCurrentAPISessionCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) WithHTTPClient(client *http.Client) *CreateCurrentAPISessionCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionCertificate adds the sessionCertificate to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) WithSessionCertificate(sessionCertificate *rest_model.CurrentAPISessionCertificateCreate) *CreateCurrentAPISessionCertificateParams {
	o.SetSessionCertificate(sessionCertificate)
	return o
}

// SetSessionCertificate adds the sessionCertificate to the create current Api session certificate params
func (o *CreateCurrentAPISessionCertificateParams) SetSessionCertificate(sessionCertificate *rest_model.CurrentAPISessionCertificateCreate) {
	o.SessionCertificate = sessionCertificate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCurrentAPISessionCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.SessionCertificate != nil {
		if err := r.SetBodyParam(o.SessionCertificate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

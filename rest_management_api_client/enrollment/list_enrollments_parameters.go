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

package enrollment

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
	"github.com/go-openapi/swag"
)

// NewListEnrollmentsParams creates a new ListEnrollmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEnrollmentsParams() *ListEnrollmentsParams {
	return &ListEnrollmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEnrollmentsParamsWithTimeout creates a new ListEnrollmentsParams object
// with the ability to set a timeout on a request.
func NewListEnrollmentsParamsWithTimeout(timeout time.Duration) *ListEnrollmentsParams {
	return &ListEnrollmentsParams{
		timeout: timeout,
	}
}

// NewListEnrollmentsParamsWithContext creates a new ListEnrollmentsParams object
// with the ability to set a context for a request.
func NewListEnrollmentsParamsWithContext(ctx context.Context) *ListEnrollmentsParams {
	return &ListEnrollmentsParams{
		Context: ctx,
	}
}

// NewListEnrollmentsParamsWithHTTPClient creates a new ListEnrollmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEnrollmentsParamsWithHTTPClient(client *http.Client) *ListEnrollmentsParams {
	return &ListEnrollmentsParams{
		HTTPClient: client,
	}
}

/* ListEnrollmentsParams contains all the parameters to send to the API endpoint
   for the list enrollments operation.

   Typically these are written to a http.Request.
*/
type ListEnrollmentsParams struct {

	// Filter.
	Filter *string

	// Limit.
	Limit *int64

	// Offset.
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list enrollments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEnrollmentsParams) WithDefaults() *ListEnrollmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list enrollments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEnrollmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list enrollments params
func (o *ListEnrollmentsParams) WithTimeout(timeout time.Duration) *ListEnrollmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list enrollments params
func (o *ListEnrollmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list enrollments params
func (o *ListEnrollmentsParams) WithContext(ctx context.Context) *ListEnrollmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list enrollments params
func (o *ListEnrollmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list enrollments params
func (o *ListEnrollmentsParams) WithHTTPClient(client *http.Client) *ListEnrollmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list enrollments params
func (o *ListEnrollmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the list enrollments params
func (o *ListEnrollmentsParams) WithFilter(filter *string) *ListEnrollmentsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list enrollments params
func (o *ListEnrollmentsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the list enrollments params
func (o *ListEnrollmentsParams) WithLimit(limit *int64) *ListEnrollmentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list enrollments params
func (o *ListEnrollmentsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list enrollments params
func (o *ListEnrollmentsParams) WithOffset(offset *int64) *ListEnrollmentsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list enrollments params
func (o *ListEnrollmentsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListEnrollmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

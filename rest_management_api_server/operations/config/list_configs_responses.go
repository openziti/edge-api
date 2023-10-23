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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListConfigsOKCode is the HTTP code returned for type ListConfigsOK
const ListConfigsOKCode int = 200

/*ListConfigsOK A list of configs

swagger:response listConfigsOK
*/
type ListConfigsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListConfigsEnvelope `json:"body,omitempty"`
}

// NewListConfigsOK creates ListConfigsOK with default headers values
func NewListConfigsOK() *ListConfigsOK {

	return &ListConfigsOK{}
}

// WithPayload adds the payload to the list configs o k response
func (o *ListConfigsOK) WithPayload(payload *rest_model.ListConfigsEnvelope) *ListConfigsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list configs o k response
func (o *ListConfigsOK) SetPayload(payload *rest_model.ListConfigsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigsBadRequestCode is the HTTP code returned for type ListConfigsBadRequest
const ListConfigsBadRequestCode int = 400

/*ListConfigsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listConfigsBadRequest
*/
type ListConfigsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigsBadRequest creates ListConfigsBadRequest with default headers values
func NewListConfigsBadRequest() *ListConfigsBadRequest {

	return &ListConfigsBadRequest{}
}

// WithPayload adds the payload to the list configs bad request response
func (o *ListConfigsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list configs bad request response
func (o *ListConfigsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigsUnauthorizedCode is the HTTP code returned for type ListConfigsUnauthorized
const ListConfigsUnauthorizedCode int = 401

/*ListConfigsUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listConfigsUnauthorized
*/
type ListConfigsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigsUnauthorized creates ListConfigsUnauthorized with default headers values
func NewListConfigsUnauthorized() *ListConfigsUnauthorized {

	return &ListConfigsUnauthorized{}
}

// WithPayload adds the payload to the list configs unauthorized response
func (o *ListConfigsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list configs unauthorized response
func (o *ListConfigsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigsTooManyRequestsCode is the HTTP code returned for type ListConfigsTooManyRequests
const ListConfigsTooManyRequestsCode int = 429

/*ListConfigsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listConfigsTooManyRequests
*/
type ListConfigsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigsTooManyRequests creates ListConfigsTooManyRequests with default headers values
func NewListConfigsTooManyRequests() *ListConfigsTooManyRequests {

	return &ListConfigsTooManyRequests{}
}

// WithPayload adds the payload to the list configs too many requests response
func (o *ListConfigsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list configs too many requests response
func (o *ListConfigsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

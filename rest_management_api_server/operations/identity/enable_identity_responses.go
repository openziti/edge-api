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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// EnableIdentityOKCode is the HTTP code returned for type EnableIdentityOK
const EnableIdentityOKCode int = 200

/*EnableIdentityOK Base empty response

swagger:response enableIdentityOK
*/
type EnableIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewEnableIdentityOK creates EnableIdentityOK with default headers values
func NewEnableIdentityOK() *EnableIdentityOK {

	return &EnableIdentityOK{}
}

// WithPayload adds the payload to the enable identity o k response
func (o *EnableIdentityOK) WithPayload(payload *rest_model.Empty) *EnableIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable identity o k response
func (o *EnableIdentityOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableIdentityUnauthorizedCode is the HTTP code returned for type EnableIdentityUnauthorized
const EnableIdentityUnauthorizedCode int = 401

/*EnableIdentityUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response enableIdentityUnauthorized
*/
type EnableIdentityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnableIdentityUnauthorized creates EnableIdentityUnauthorized with default headers values
func NewEnableIdentityUnauthorized() *EnableIdentityUnauthorized {

	return &EnableIdentityUnauthorized{}
}

// WithPayload adds the payload to the enable identity unauthorized response
func (o *EnableIdentityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *EnableIdentityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable identity unauthorized response
func (o *EnableIdentityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableIdentityNotFoundCode is the HTTP code returned for type EnableIdentityNotFound
const EnableIdentityNotFoundCode int = 404

/*EnableIdentityNotFound The requested resource does not exist

swagger:response enableIdentityNotFound
*/
type EnableIdentityNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnableIdentityNotFound creates EnableIdentityNotFound with default headers values
func NewEnableIdentityNotFound() *EnableIdentityNotFound {

	return &EnableIdentityNotFound{}
}

// WithPayload adds the payload to the enable identity not found response
func (o *EnableIdentityNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *EnableIdentityNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable identity not found response
func (o *EnableIdentityNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableIdentityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableIdentityTooManyRequestsCode is the HTTP code returned for type EnableIdentityTooManyRequests
const EnableIdentityTooManyRequestsCode int = 429

/*EnableIdentityTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response enableIdentityTooManyRequests
*/
type EnableIdentityTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnableIdentityTooManyRequests creates EnableIdentityTooManyRequests with default headers values
func NewEnableIdentityTooManyRequests() *EnableIdentityTooManyRequests {

	return &EnableIdentityTooManyRequests{}
}

// WithPayload adds the payload to the enable identity too many requests response
func (o *EnableIdentityTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *EnableIdentityTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable identity too many requests response
func (o *EnableIdentityTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableIdentityTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableIdentityServiceUnavailableCode is the HTTP code returned for type EnableIdentityServiceUnavailable
const EnableIdentityServiceUnavailableCode int = 503

/*EnableIdentityServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response enableIdentityServiceUnavailable
*/
type EnableIdentityServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnableIdentityServiceUnavailable creates EnableIdentityServiceUnavailable with default headers values
func NewEnableIdentityServiceUnavailable() *EnableIdentityServiceUnavailable {

	return &EnableIdentityServiceUnavailable{}
}

// WithPayload adds the payload to the enable identity service unavailable response
func (o *EnableIdentityServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *EnableIdentityServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable identity service unavailable response
func (o *EnableIdentityServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableIdentityServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

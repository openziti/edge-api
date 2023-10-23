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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchRouterOKCode is the HTTP code returned for type PatchRouterOK
const PatchRouterOKCode int = 200

/*PatchRouterOK The patch request was successful and the resource has been altered

swagger:response patchRouterOK
*/
type PatchRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchRouterOK creates PatchRouterOK with default headers values
func NewPatchRouterOK() *PatchRouterOK {

	return &PatchRouterOK{}
}

// WithPayload adds the payload to the patch router o k response
func (o *PatchRouterOK) WithPayload(payload *rest_model.Empty) *PatchRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch router o k response
func (o *PatchRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRouterBadRequestCode is the HTTP code returned for type PatchRouterBadRequest
const PatchRouterBadRequestCode int = 400

/*PatchRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchRouterBadRequest
*/
type PatchRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchRouterBadRequest creates PatchRouterBadRequest with default headers values
func NewPatchRouterBadRequest() *PatchRouterBadRequest {

	return &PatchRouterBadRequest{}
}

// WithPayload adds the payload to the patch router bad request response
func (o *PatchRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch router bad request response
func (o *PatchRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRouterUnauthorizedCode is the HTTP code returned for type PatchRouterUnauthorized
const PatchRouterUnauthorizedCode int = 401

/*PatchRouterUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchRouterUnauthorized
*/
type PatchRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchRouterUnauthorized creates PatchRouterUnauthorized with default headers values
func NewPatchRouterUnauthorized() *PatchRouterUnauthorized {

	return &PatchRouterUnauthorized{}
}

// WithPayload adds the payload to the patch router unauthorized response
func (o *PatchRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch router unauthorized response
func (o *PatchRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRouterNotFoundCode is the HTTP code returned for type PatchRouterNotFound
const PatchRouterNotFoundCode int = 404

/*PatchRouterNotFound The requested resource does not exist

swagger:response patchRouterNotFound
*/
type PatchRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchRouterNotFound creates PatchRouterNotFound with default headers values
func NewPatchRouterNotFound() *PatchRouterNotFound {

	return &PatchRouterNotFound{}
}

// WithPayload adds the payload to the patch router not found response
func (o *PatchRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch router not found response
func (o *PatchRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRouterTooManyRequestsCode is the HTTP code returned for type PatchRouterTooManyRequests
const PatchRouterTooManyRequestsCode int = 429

/*PatchRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchRouterTooManyRequests
*/
type PatchRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchRouterTooManyRequests creates PatchRouterTooManyRequests with default headers values
func NewPatchRouterTooManyRequests() *PatchRouterTooManyRequests {

	return &PatchRouterTooManyRequests{}
}

// WithPayload adds the payload to the patch router too many requests response
func (o *PatchRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch router too many requests response
func (o *PatchRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

// PatchIdentityOKCode is the HTTP code returned for type PatchIdentityOK
const PatchIdentityOKCode int = 200

/*PatchIdentityOK The patch request was successful and the resource has been altered

swagger:response patchIdentityOK
*/
type PatchIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchIdentityOK creates PatchIdentityOK with default headers values
func NewPatchIdentityOK() *PatchIdentityOK {

	return &PatchIdentityOK{}
}

// WithPayload adds the payload to the patch identity o k response
func (o *PatchIdentityOK) WithPayload(payload *rest_model.Empty) *PatchIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch identity o k response
func (o *PatchIdentityOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchIdentityBadRequestCode is the HTTP code returned for type PatchIdentityBadRequest
const PatchIdentityBadRequestCode int = 400

/*PatchIdentityBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchIdentityBadRequest
*/
type PatchIdentityBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchIdentityBadRequest creates PatchIdentityBadRequest with default headers values
func NewPatchIdentityBadRequest() *PatchIdentityBadRequest {

	return &PatchIdentityBadRequest{}
}

// WithPayload adds the payload to the patch identity bad request response
func (o *PatchIdentityBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchIdentityBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch identity bad request response
func (o *PatchIdentityBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchIdentityBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchIdentityUnauthorizedCode is the HTTP code returned for type PatchIdentityUnauthorized
const PatchIdentityUnauthorizedCode int = 401

/*PatchIdentityUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchIdentityUnauthorized
*/
type PatchIdentityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchIdentityUnauthorized creates PatchIdentityUnauthorized with default headers values
func NewPatchIdentityUnauthorized() *PatchIdentityUnauthorized {

	return &PatchIdentityUnauthorized{}
}

// WithPayload adds the payload to the patch identity unauthorized response
func (o *PatchIdentityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchIdentityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch identity unauthorized response
func (o *PatchIdentityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchIdentityNotFoundCode is the HTTP code returned for type PatchIdentityNotFound
const PatchIdentityNotFoundCode int = 404

/*PatchIdentityNotFound The requested resource does not exist

swagger:response patchIdentityNotFound
*/
type PatchIdentityNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchIdentityNotFound creates PatchIdentityNotFound with default headers values
func NewPatchIdentityNotFound() *PatchIdentityNotFound {

	return &PatchIdentityNotFound{}
}

// WithPayload adds the payload to the patch identity not found response
func (o *PatchIdentityNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchIdentityNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch identity not found response
func (o *PatchIdentityNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchIdentityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchIdentityTooManyRequestsCode is the HTTP code returned for type PatchIdentityTooManyRequests
const PatchIdentityTooManyRequestsCode int = 429

/*PatchIdentityTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchIdentityTooManyRequests
*/
type PatchIdentityTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchIdentityTooManyRequests creates PatchIdentityTooManyRequests with default headers values
func NewPatchIdentityTooManyRequests() *PatchIdentityTooManyRequests {

	return &PatchIdentityTooManyRequests{}
}

// WithPayload adds the payload to the patch identity too many requests response
func (o *PatchIdentityTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchIdentityTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch identity too many requests response
func (o *PatchIdentityTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchIdentityTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

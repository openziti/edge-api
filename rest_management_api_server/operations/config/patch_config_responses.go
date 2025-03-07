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

// PatchConfigOKCode is the HTTP code returned for type PatchConfigOK
const PatchConfigOKCode int = 200

/*PatchConfigOK The patch request was successful and the resource has been altered

swagger:response patchConfigOK
*/
type PatchConfigOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchConfigOK creates PatchConfigOK with default headers values
func NewPatchConfigOK() *PatchConfigOK {

	return &PatchConfigOK{}
}

// WithPayload adds the payload to the patch config o k response
func (o *PatchConfigOK) WithPayload(payload *rest_model.Empty) *PatchConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config o k response
func (o *PatchConfigOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigBadRequestCode is the HTTP code returned for type PatchConfigBadRequest
const PatchConfigBadRequestCode int = 400

/*PatchConfigBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchConfigBadRequest
*/
type PatchConfigBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigBadRequest creates PatchConfigBadRequest with default headers values
func NewPatchConfigBadRequest() *PatchConfigBadRequest {

	return &PatchConfigBadRequest{}
}

// WithPayload adds the payload to the patch config bad request response
func (o *PatchConfigBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config bad request response
func (o *PatchConfigBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigUnauthorizedCode is the HTTP code returned for type PatchConfigUnauthorized
const PatchConfigUnauthorizedCode int = 401

/*PatchConfigUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchConfigUnauthorized
*/
type PatchConfigUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigUnauthorized creates PatchConfigUnauthorized with default headers values
func NewPatchConfigUnauthorized() *PatchConfigUnauthorized {

	return &PatchConfigUnauthorized{}
}

// WithPayload adds the payload to the patch config unauthorized response
func (o *PatchConfigUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config unauthorized response
func (o *PatchConfigUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigNotFoundCode is the HTTP code returned for type PatchConfigNotFound
const PatchConfigNotFoundCode int = 404

/*PatchConfigNotFound The requested resource does not exist

swagger:response patchConfigNotFound
*/
type PatchConfigNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigNotFound creates PatchConfigNotFound with default headers values
func NewPatchConfigNotFound() *PatchConfigNotFound {

	return &PatchConfigNotFound{}
}

// WithPayload adds the payload to the patch config not found response
func (o *PatchConfigNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config not found response
func (o *PatchConfigNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigTooManyRequestsCode is the HTTP code returned for type PatchConfigTooManyRequests
const PatchConfigTooManyRequestsCode int = 429

/*PatchConfigTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchConfigTooManyRequests
*/
type PatchConfigTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigTooManyRequests creates PatchConfigTooManyRequests with default headers values
func NewPatchConfigTooManyRequests() *PatchConfigTooManyRequests {

	return &PatchConfigTooManyRequests{}
}

// WithPayload adds the payload to the patch config too many requests response
func (o *PatchConfigTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config too many requests response
func (o *PatchConfigTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigServiceUnavailableCode is the HTTP code returned for type PatchConfigServiceUnavailable
const PatchConfigServiceUnavailableCode int = 503

/*PatchConfigServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response patchConfigServiceUnavailable
*/
type PatchConfigServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigServiceUnavailable creates PatchConfigServiceUnavailable with default headers values
func NewPatchConfigServiceUnavailable() *PatchConfigServiceUnavailable {

	return &PatchConfigServiceUnavailable{}
}

// WithPayload adds the payload to the patch config service unavailable response
func (o *PatchConfigServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config service unavailable response
func (o *PatchConfigServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

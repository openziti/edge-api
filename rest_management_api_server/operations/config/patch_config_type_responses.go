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

// PatchConfigTypeOKCode is the HTTP code returned for type PatchConfigTypeOK
const PatchConfigTypeOKCode int = 200

/*PatchConfigTypeOK The patch request was successful and the resource has been altered

swagger:response patchConfigTypeOK
*/
type PatchConfigTypeOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchConfigTypeOK creates PatchConfigTypeOK with default headers values
func NewPatchConfigTypeOK() *PatchConfigTypeOK {

	return &PatchConfigTypeOK{}
}

// WithPayload adds the payload to the patch config type o k response
func (o *PatchConfigTypeOK) WithPayload(payload *rest_model.Empty) *PatchConfigTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config type o k response
func (o *PatchConfigTypeOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigTypeBadRequestCode is the HTTP code returned for type PatchConfigTypeBadRequest
const PatchConfigTypeBadRequestCode int = 400

/*PatchConfigTypeBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchConfigTypeBadRequest
*/
type PatchConfigTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigTypeBadRequest creates PatchConfigTypeBadRequest with default headers values
func NewPatchConfigTypeBadRequest() *PatchConfigTypeBadRequest {

	return &PatchConfigTypeBadRequest{}
}

// WithPayload adds the payload to the patch config type bad request response
func (o *PatchConfigTypeBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config type bad request response
func (o *PatchConfigTypeBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigTypeUnauthorizedCode is the HTTP code returned for type PatchConfigTypeUnauthorized
const PatchConfigTypeUnauthorizedCode int = 401

/*PatchConfigTypeUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchConfigTypeUnauthorized
*/
type PatchConfigTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigTypeUnauthorized creates PatchConfigTypeUnauthorized with default headers values
func NewPatchConfigTypeUnauthorized() *PatchConfigTypeUnauthorized {

	return &PatchConfigTypeUnauthorized{}
}

// WithPayload adds the payload to the patch config type unauthorized response
func (o *PatchConfigTypeUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config type unauthorized response
func (o *PatchConfigTypeUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigTypeNotFoundCode is the HTTP code returned for type PatchConfigTypeNotFound
const PatchConfigTypeNotFoundCode int = 404

/*PatchConfigTypeNotFound The requested resource does not exist

swagger:response patchConfigTypeNotFound
*/
type PatchConfigTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigTypeNotFound creates PatchConfigTypeNotFound with default headers values
func NewPatchConfigTypeNotFound() *PatchConfigTypeNotFound {

	return &PatchConfigTypeNotFound{}
}

// WithPayload adds the payload to the patch config type not found response
func (o *PatchConfigTypeNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config type not found response
func (o *PatchConfigTypeNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigTypeTooManyRequestsCode is the HTTP code returned for type PatchConfigTypeTooManyRequests
const PatchConfigTypeTooManyRequestsCode int = 429

/*PatchConfigTypeTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchConfigTypeTooManyRequests
*/
type PatchConfigTypeTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigTypeTooManyRequests creates PatchConfigTypeTooManyRequests with default headers values
func NewPatchConfigTypeTooManyRequests() *PatchConfigTypeTooManyRequests {

	return &PatchConfigTypeTooManyRequests{}
}

// WithPayload adds the payload to the patch config type too many requests response
func (o *PatchConfigTypeTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigTypeTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config type too many requests response
func (o *PatchConfigTypeTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTypeTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchConfigTypeServiceUnavailableCode is the HTTP code returned for type PatchConfigTypeServiceUnavailable
const PatchConfigTypeServiceUnavailableCode int = 503

/*PatchConfigTypeServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response patchConfigTypeServiceUnavailable
*/
type PatchConfigTypeServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchConfigTypeServiceUnavailable creates PatchConfigTypeServiceUnavailable with default headers values
func NewPatchConfigTypeServiceUnavailable() *PatchConfigTypeServiceUnavailable {

	return &PatchConfigTypeServiceUnavailable{}
}

// WithPayload adds the payload to the patch config type service unavailable response
func (o *PatchConfigTypeServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchConfigTypeServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch config type service unavailable response
func (o *PatchConfigTypeServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchConfigTypeServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

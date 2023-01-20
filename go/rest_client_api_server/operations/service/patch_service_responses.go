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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/go/rest_model"
)

// PatchServiceOKCode is the HTTP code returned for type PatchServiceOK
const PatchServiceOKCode int = 200

/*PatchServiceOK The patch request was successful and the resource has been altered

swagger:response patchServiceOK
*/
type PatchServiceOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchServiceOK creates PatchServiceOK with default headers values
func NewPatchServiceOK() *PatchServiceOK {

	return &PatchServiceOK{}
}

// WithPayload adds the payload to the patch service o k response
func (o *PatchServiceOK) WithPayload(payload *rest_model.Empty) *PatchServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service o k response
func (o *PatchServiceOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServiceBadRequestCode is the HTTP code returned for type PatchServiceBadRequest
const PatchServiceBadRequestCode int = 400

/*PatchServiceBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchServiceBadRequest
*/
type PatchServiceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServiceBadRequest creates PatchServiceBadRequest with default headers values
func NewPatchServiceBadRequest() *PatchServiceBadRequest {

	return &PatchServiceBadRequest{}
}

// WithPayload adds the payload to the patch service bad request response
func (o *PatchServiceBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServiceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service bad request response
func (o *PatchServiceBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServiceUnauthorizedCode is the HTTP code returned for type PatchServiceUnauthorized
const PatchServiceUnauthorizedCode int = 401

/*PatchServiceUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchServiceUnauthorized
*/
type PatchServiceUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServiceUnauthorized creates PatchServiceUnauthorized with default headers values
func NewPatchServiceUnauthorized() *PatchServiceUnauthorized {

	return &PatchServiceUnauthorized{}
}

// WithPayload adds the payload to the patch service unauthorized response
func (o *PatchServiceUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServiceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service unauthorized response
func (o *PatchServiceUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServiceNotFoundCode is the HTTP code returned for type PatchServiceNotFound
const PatchServiceNotFoundCode int = 404

/*PatchServiceNotFound The requested resource does not exist

swagger:response patchServiceNotFound
*/
type PatchServiceNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServiceNotFound creates PatchServiceNotFound with default headers values
func NewPatchServiceNotFound() *PatchServiceNotFound {

	return &PatchServiceNotFound{}
}

// WithPayload adds the payload to the patch service not found response
func (o *PatchServiceNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServiceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service not found response
func (o *PatchServiceNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

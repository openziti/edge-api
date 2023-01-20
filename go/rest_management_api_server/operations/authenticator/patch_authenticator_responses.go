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

package authenticator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/go/rest_model"
)

// PatchAuthenticatorOKCode is the HTTP code returned for type PatchAuthenticatorOK
const PatchAuthenticatorOKCode int = 200

/*PatchAuthenticatorOK The patch request was successful and the resource has been altered

swagger:response patchAuthenticatorOK
*/
type PatchAuthenticatorOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchAuthenticatorOK creates PatchAuthenticatorOK with default headers values
func NewPatchAuthenticatorOK() *PatchAuthenticatorOK {

	return &PatchAuthenticatorOK{}
}

// WithPayload adds the payload to the patch authenticator o k response
func (o *PatchAuthenticatorOK) WithPayload(payload *rest_model.Empty) *PatchAuthenticatorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch authenticator o k response
func (o *PatchAuthenticatorOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAuthenticatorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAuthenticatorBadRequestCode is the HTTP code returned for type PatchAuthenticatorBadRequest
const PatchAuthenticatorBadRequestCode int = 400

/*PatchAuthenticatorBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchAuthenticatorBadRequest
*/
type PatchAuthenticatorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchAuthenticatorBadRequest creates PatchAuthenticatorBadRequest with default headers values
func NewPatchAuthenticatorBadRequest() *PatchAuthenticatorBadRequest {

	return &PatchAuthenticatorBadRequest{}
}

// WithPayload adds the payload to the patch authenticator bad request response
func (o *PatchAuthenticatorBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchAuthenticatorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch authenticator bad request response
func (o *PatchAuthenticatorBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAuthenticatorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAuthenticatorUnauthorizedCode is the HTTP code returned for type PatchAuthenticatorUnauthorized
const PatchAuthenticatorUnauthorizedCode int = 401

/*PatchAuthenticatorUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchAuthenticatorUnauthorized
*/
type PatchAuthenticatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchAuthenticatorUnauthorized creates PatchAuthenticatorUnauthorized with default headers values
func NewPatchAuthenticatorUnauthorized() *PatchAuthenticatorUnauthorized {

	return &PatchAuthenticatorUnauthorized{}
}

// WithPayload adds the payload to the patch authenticator unauthorized response
func (o *PatchAuthenticatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchAuthenticatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch authenticator unauthorized response
func (o *PatchAuthenticatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAuthenticatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAuthenticatorNotFoundCode is the HTTP code returned for type PatchAuthenticatorNotFound
const PatchAuthenticatorNotFoundCode int = 404

/*PatchAuthenticatorNotFound The requested resource does not exist

swagger:response patchAuthenticatorNotFound
*/
type PatchAuthenticatorNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchAuthenticatorNotFound creates PatchAuthenticatorNotFound with default headers values
func NewPatchAuthenticatorNotFound() *PatchAuthenticatorNotFound {

	return &PatchAuthenticatorNotFound{}
}

// WithPayload adds the payload to the patch authenticator not found response
func (o *PatchAuthenticatorNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchAuthenticatorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch authenticator not found response
func (o *PatchAuthenticatorNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAuthenticatorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

package external_jwt_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/go/rest_model"
)

// PatchExternalJWTSignerOKCode is the HTTP code returned for type PatchExternalJWTSignerOK
const PatchExternalJWTSignerOKCode int = 200

/*PatchExternalJWTSignerOK The patch request was successful and the resource has been altered

swagger:response patchExternalJwtSignerOK
*/
type PatchExternalJWTSignerOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchExternalJWTSignerOK creates PatchExternalJWTSignerOK with default headers values
func NewPatchExternalJWTSignerOK() *PatchExternalJWTSignerOK {

	return &PatchExternalJWTSignerOK{}
}

// WithPayload adds the payload to the patch external Jwt signer o k response
func (o *PatchExternalJWTSignerOK) WithPayload(payload *rest_model.Empty) *PatchExternalJWTSignerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch external Jwt signer o k response
func (o *PatchExternalJWTSignerOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchExternalJWTSignerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchExternalJWTSignerBadRequestCode is the HTTP code returned for type PatchExternalJWTSignerBadRequest
const PatchExternalJWTSignerBadRequestCode int = 400

/*PatchExternalJWTSignerBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchExternalJwtSignerBadRequest
*/
type PatchExternalJWTSignerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchExternalJWTSignerBadRequest creates PatchExternalJWTSignerBadRequest with default headers values
func NewPatchExternalJWTSignerBadRequest() *PatchExternalJWTSignerBadRequest {

	return &PatchExternalJWTSignerBadRequest{}
}

// WithPayload adds the payload to the patch external Jwt signer bad request response
func (o *PatchExternalJWTSignerBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchExternalJWTSignerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch external Jwt signer bad request response
func (o *PatchExternalJWTSignerBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchExternalJWTSignerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchExternalJWTSignerUnauthorizedCode is the HTTP code returned for type PatchExternalJWTSignerUnauthorized
const PatchExternalJWTSignerUnauthorizedCode int = 401

/*PatchExternalJWTSignerUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchExternalJwtSignerUnauthorized
*/
type PatchExternalJWTSignerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchExternalJWTSignerUnauthorized creates PatchExternalJWTSignerUnauthorized with default headers values
func NewPatchExternalJWTSignerUnauthorized() *PatchExternalJWTSignerUnauthorized {

	return &PatchExternalJWTSignerUnauthorized{}
}

// WithPayload adds the payload to the patch external Jwt signer unauthorized response
func (o *PatchExternalJWTSignerUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchExternalJWTSignerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch external Jwt signer unauthorized response
func (o *PatchExternalJWTSignerUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchExternalJWTSignerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchExternalJWTSignerNotFoundCode is the HTTP code returned for type PatchExternalJWTSignerNotFound
const PatchExternalJWTSignerNotFoundCode int = 404

/*PatchExternalJWTSignerNotFound The requested resource does not exist

swagger:response patchExternalJwtSignerNotFound
*/
type PatchExternalJWTSignerNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchExternalJWTSignerNotFound creates PatchExternalJWTSignerNotFound with default headers values
func NewPatchExternalJWTSignerNotFound() *PatchExternalJWTSignerNotFound {

	return &PatchExternalJWTSignerNotFound{}
}

// WithPayload adds the payload to the patch external Jwt signer not found response
func (o *PatchExternalJWTSignerNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchExternalJWTSignerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch external Jwt signer not found response
func (o *PatchExternalJWTSignerNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchExternalJWTSignerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

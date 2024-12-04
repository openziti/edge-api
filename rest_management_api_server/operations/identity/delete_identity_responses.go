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

// DeleteIdentityOKCode is the HTTP code returned for type DeleteIdentityOK
const DeleteIdentityOKCode int = 200

/*DeleteIdentityOK The delete request was successful and the resource has been removed

swagger:response deleteIdentityOK
*/
type DeleteIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteIdentityOK creates DeleteIdentityOK with default headers values
func NewDeleteIdentityOK() *DeleteIdentityOK {

	return &DeleteIdentityOK{}
}

// WithPayload adds the payload to the delete identity o k response
func (o *DeleteIdentityOK) WithPayload(payload *rest_model.Empty) *DeleteIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete identity o k response
func (o *DeleteIdentityOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIdentityBadRequestCode is the HTTP code returned for type DeleteIdentityBadRequest
const DeleteIdentityBadRequestCode int = 400

/*DeleteIdentityBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteIdentityBadRequest
*/
type DeleteIdentityBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteIdentityBadRequest creates DeleteIdentityBadRequest with default headers values
func NewDeleteIdentityBadRequest() *DeleteIdentityBadRequest {

	return &DeleteIdentityBadRequest{}
}

// WithPayload adds the payload to the delete identity bad request response
func (o *DeleteIdentityBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteIdentityBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete identity bad request response
func (o *DeleteIdentityBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIdentityBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIdentityUnauthorizedCode is the HTTP code returned for type DeleteIdentityUnauthorized
const DeleteIdentityUnauthorizedCode int = 401

/*DeleteIdentityUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteIdentityUnauthorized
*/
type DeleteIdentityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteIdentityUnauthorized creates DeleteIdentityUnauthorized with default headers values
func NewDeleteIdentityUnauthorized() *DeleteIdentityUnauthorized {

	return &DeleteIdentityUnauthorized{}
}

// WithPayload adds the payload to the delete identity unauthorized response
func (o *DeleteIdentityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteIdentityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete identity unauthorized response
func (o *DeleteIdentityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIdentityNotFoundCode is the HTTP code returned for type DeleteIdentityNotFound
const DeleteIdentityNotFoundCode int = 404

/*DeleteIdentityNotFound The requested resource does not exist

swagger:response deleteIdentityNotFound
*/
type DeleteIdentityNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteIdentityNotFound creates DeleteIdentityNotFound with default headers values
func NewDeleteIdentityNotFound() *DeleteIdentityNotFound {

	return &DeleteIdentityNotFound{}
}

// WithPayload adds the payload to the delete identity not found response
func (o *DeleteIdentityNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteIdentityNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete identity not found response
func (o *DeleteIdentityNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIdentityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIdentityConflictCode is the HTTP code returned for type DeleteIdentityConflict
const DeleteIdentityConflictCode int = 409

/*DeleteIdentityConflict The resource requested to be removed/altered cannot be as it is referenced by another object.

swagger:response deleteIdentityConflict
*/
type DeleteIdentityConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteIdentityConflict creates DeleteIdentityConflict with default headers values
func NewDeleteIdentityConflict() *DeleteIdentityConflict {

	return &DeleteIdentityConflict{}
}

// WithPayload adds the payload to the delete identity conflict response
func (o *DeleteIdentityConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteIdentityConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete identity conflict response
func (o *DeleteIdentityConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIdentityConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIdentityTooManyRequestsCode is the HTTP code returned for type DeleteIdentityTooManyRequests
const DeleteIdentityTooManyRequestsCode int = 429

/*DeleteIdentityTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteIdentityTooManyRequests
*/
type DeleteIdentityTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteIdentityTooManyRequests creates DeleteIdentityTooManyRequests with default headers values
func NewDeleteIdentityTooManyRequests() *DeleteIdentityTooManyRequests {

	return &DeleteIdentityTooManyRequests{}
}

// WithPayload adds the payload to the delete identity too many requests response
func (o *DeleteIdentityTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteIdentityTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete identity too many requests response
func (o *DeleteIdentityTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIdentityTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

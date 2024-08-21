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

// DeleteTransitRouterOKCode is the HTTP code returned for type DeleteTransitRouterOK
const DeleteTransitRouterOKCode int = 200

/*DeleteTransitRouterOK The delete request was successful and the resource has been removed

swagger:response deleteTransitRouterOK
*/
type DeleteTransitRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteTransitRouterOK creates DeleteTransitRouterOK with default headers values
func NewDeleteTransitRouterOK() *DeleteTransitRouterOK {

	return &DeleteTransitRouterOK{}
}

// WithPayload adds the payload to the delete transit router o k response
func (o *DeleteTransitRouterOK) WithPayload(payload *rest_model.Empty) *DeleteTransitRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transit router o k response
func (o *DeleteTransitRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransitRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTransitRouterBadRequestCode is the HTTP code returned for type DeleteTransitRouterBadRequest
const DeleteTransitRouterBadRequestCode int = 400

/*DeleteTransitRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteTransitRouterBadRequest
*/
type DeleteTransitRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteTransitRouterBadRequest creates DeleteTransitRouterBadRequest with default headers values
func NewDeleteTransitRouterBadRequest() *DeleteTransitRouterBadRequest {

	return &DeleteTransitRouterBadRequest{}
}

// WithPayload adds the payload to the delete transit router bad request response
func (o *DeleteTransitRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteTransitRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transit router bad request response
func (o *DeleteTransitRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransitRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTransitRouterUnauthorizedCode is the HTTP code returned for type DeleteTransitRouterUnauthorized
const DeleteTransitRouterUnauthorizedCode int = 401

/*DeleteTransitRouterUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteTransitRouterUnauthorized
*/
type DeleteTransitRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteTransitRouterUnauthorized creates DeleteTransitRouterUnauthorized with default headers values
func NewDeleteTransitRouterUnauthorized() *DeleteTransitRouterUnauthorized {

	return &DeleteTransitRouterUnauthorized{}
}

// WithPayload adds the payload to the delete transit router unauthorized response
func (o *DeleteTransitRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteTransitRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transit router unauthorized response
func (o *DeleteTransitRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransitRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTransitRouterConflictCode is the HTTP code returned for type DeleteTransitRouterConflict
const DeleteTransitRouterConflictCode int = 409

/*DeleteTransitRouterConflict The resource requested to be removed/altered cannot be as it is referenced by another object.

swagger:response deleteTransitRouterConflict
*/
type DeleteTransitRouterConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteTransitRouterConflict creates DeleteTransitRouterConflict with default headers values
func NewDeleteTransitRouterConflict() *DeleteTransitRouterConflict {

	return &DeleteTransitRouterConflict{}
}

// WithPayload adds the payload to the delete transit router conflict response
func (o *DeleteTransitRouterConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteTransitRouterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transit router conflict response
func (o *DeleteTransitRouterConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransitRouterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTransitRouterTooManyRequestsCode is the HTTP code returned for type DeleteTransitRouterTooManyRequests
const DeleteTransitRouterTooManyRequestsCode int = 429

/*DeleteTransitRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteTransitRouterTooManyRequests
*/
type DeleteTransitRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteTransitRouterTooManyRequests creates DeleteTransitRouterTooManyRequests with default headers values
func NewDeleteTransitRouterTooManyRequests() *DeleteTransitRouterTooManyRequests {

	return &DeleteTransitRouterTooManyRequests{}
}

// WithPayload adds the payload to the delete transit router too many requests response
func (o *DeleteTransitRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteTransitRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete transit router too many requests response
func (o *DeleteTransitRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTransitRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

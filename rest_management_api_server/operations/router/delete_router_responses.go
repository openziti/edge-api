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

// DeleteRouterOKCode is the HTTP code returned for type DeleteRouterOK
const DeleteRouterOKCode int = 200

/*DeleteRouterOK The delete request was successful and the resource has been removed

swagger:response deleteRouterOK
*/
type DeleteRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteRouterOK creates DeleteRouterOK with default headers values
func NewDeleteRouterOK() *DeleteRouterOK {

	return &DeleteRouterOK{}
}

// WithPayload adds the payload to the delete router o k response
func (o *DeleteRouterOK) WithPayload(payload *rest_model.Empty) *DeleteRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete router o k response
func (o *DeleteRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRouterBadRequestCode is the HTTP code returned for type DeleteRouterBadRequest
const DeleteRouterBadRequestCode int = 400

/*DeleteRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteRouterBadRequest
*/
type DeleteRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteRouterBadRequest creates DeleteRouterBadRequest with default headers values
func NewDeleteRouterBadRequest() *DeleteRouterBadRequest {

	return &DeleteRouterBadRequest{}
}

// WithPayload adds the payload to the delete router bad request response
func (o *DeleteRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete router bad request response
func (o *DeleteRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRouterUnauthorizedCode is the HTTP code returned for type DeleteRouterUnauthorized
const DeleteRouterUnauthorizedCode int = 401

/*DeleteRouterUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteRouterUnauthorized
*/
type DeleteRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteRouterUnauthorized creates DeleteRouterUnauthorized with default headers values
func NewDeleteRouterUnauthorized() *DeleteRouterUnauthorized {

	return &DeleteRouterUnauthorized{}
}

// WithPayload adds the payload to the delete router unauthorized response
func (o *DeleteRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete router unauthorized response
func (o *DeleteRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRouterConflictCode is the HTTP code returned for type DeleteRouterConflict
const DeleteRouterConflictCode int = 409

/*DeleteRouterConflict The resource requested to be removed/altered cannot be as it is referenced by another object.

swagger:response deleteRouterConflict
*/
type DeleteRouterConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteRouterConflict creates DeleteRouterConflict with default headers values
func NewDeleteRouterConflict() *DeleteRouterConflict {

	return &DeleteRouterConflict{}
}

// WithPayload adds the payload to the delete router conflict response
func (o *DeleteRouterConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteRouterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete router conflict response
func (o *DeleteRouterConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRouterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRouterTooManyRequestsCode is the HTTP code returned for type DeleteRouterTooManyRequests
const DeleteRouterTooManyRequestsCode int = 429

/*DeleteRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteRouterTooManyRequests
*/
type DeleteRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteRouterTooManyRequests creates DeleteRouterTooManyRequests with default headers values
func NewDeleteRouterTooManyRequests() *DeleteRouterTooManyRequests {

	return &DeleteRouterTooManyRequests{}
}

// WithPayload adds the payload to the delete router too many requests response
func (o *DeleteRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete router too many requests response
func (o *DeleteRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRouterServiceUnavailableCode is the HTTP code returned for type DeleteRouterServiceUnavailable
const DeleteRouterServiceUnavailableCode int = 503

/*DeleteRouterServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response deleteRouterServiceUnavailable
*/
type DeleteRouterServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteRouterServiceUnavailable creates DeleteRouterServiceUnavailable with default headers values
func NewDeleteRouterServiceUnavailable() *DeleteRouterServiceUnavailable {

	return &DeleteRouterServiceUnavailable{}
}

// WithPayload adds the payload to the delete router service unavailable response
func (o *DeleteRouterServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteRouterServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete router service unavailable response
func (o *DeleteRouterServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRouterServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

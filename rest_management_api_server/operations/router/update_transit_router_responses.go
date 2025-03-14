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

// UpdateTransitRouterOKCode is the HTTP code returned for type UpdateTransitRouterOK
const UpdateTransitRouterOKCode int = 200

/*UpdateTransitRouterOK The update request was successful and the resource has been altered

swagger:response updateTransitRouterOK
*/
type UpdateTransitRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdateTransitRouterOK creates UpdateTransitRouterOK with default headers values
func NewUpdateTransitRouterOK() *UpdateTransitRouterOK {

	return &UpdateTransitRouterOK{}
}

// WithPayload adds the payload to the update transit router o k response
func (o *UpdateTransitRouterOK) WithPayload(payload *rest_model.Empty) *UpdateTransitRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update transit router o k response
func (o *UpdateTransitRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTransitRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTransitRouterBadRequestCode is the HTTP code returned for type UpdateTransitRouterBadRequest
const UpdateTransitRouterBadRequestCode int = 400

/*UpdateTransitRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updateTransitRouterBadRequest
*/
type UpdateTransitRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateTransitRouterBadRequest creates UpdateTransitRouterBadRequest with default headers values
func NewUpdateTransitRouterBadRequest() *UpdateTransitRouterBadRequest {

	return &UpdateTransitRouterBadRequest{}
}

// WithPayload adds the payload to the update transit router bad request response
func (o *UpdateTransitRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateTransitRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update transit router bad request response
func (o *UpdateTransitRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTransitRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTransitRouterUnauthorizedCode is the HTTP code returned for type UpdateTransitRouterUnauthorized
const UpdateTransitRouterUnauthorizedCode int = 401

/*UpdateTransitRouterUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response updateTransitRouterUnauthorized
*/
type UpdateTransitRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateTransitRouterUnauthorized creates UpdateTransitRouterUnauthorized with default headers values
func NewUpdateTransitRouterUnauthorized() *UpdateTransitRouterUnauthorized {

	return &UpdateTransitRouterUnauthorized{}
}

// WithPayload adds the payload to the update transit router unauthorized response
func (o *UpdateTransitRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateTransitRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update transit router unauthorized response
func (o *UpdateTransitRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTransitRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTransitRouterNotFoundCode is the HTTP code returned for type UpdateTransitRouterNotFound
const UpdateTransitRouterNotFoundCode int = 404

/*UpdateTransitRouterNotFound The requested resource does not exist

swagger:response updateTransitRouterNotFound
*/
type UpdateTransitRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateTransitRouterNotFound creates UpdateTransitRouterNotFound with default headers values
func NewUpdateTransitRouterNotFound() *UpdateTransitRouterNotFound {

	return &UpdateTransitRouterNotFound{}
}

// WithPayload adds the payload to the update transit router not found response
func (o *UpdateTransitRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateTransitRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update transit router not found response
func (o *UpdateTransitRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTransitRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTransitRouterTooManyRequestsCode is the HTTP code returned for type UpdateTransitRouterTooManyRequests
const UpdateTransitRouterTooManyRequestsCode int = 429

/*UpdateTransitRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response updateTransitRouterTooManyRequests
*/
type UpdateTransitRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateTransitRouterTooManyRequests creates UpdateTransitRouterTooManyRequests with default headers values
func NewUpdateTransitRouterTooManyRequests() *UpdateTransitRouterTooManyRequests {

	return &UpdateTransitRouterTooManyRequests{}
}

// WithPayload adds the payload to the update transit router too many requests response
func (o *UpdateTransitRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateTransitRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update transit router too many requests response
func (o *UpdateTransitRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTransitRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTransitRouterServiceUnavailableCode is the HTTP code returned for type UpdateTransitRouterServiceUnavailable
const UpdateTransitRouterServiceUnavailableCode int = 503

/*UpdateTransitRouterServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response updateTransitRouterServiceUnavailable
*/
type UpdateTransitRouterServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateTransitRouterServiceUnavailable creates UpdateTransitRouterServiceUnavailable with default headers values
func NewUpdateTransitRouterServiceUnavailable() *UpdateTransitRouterServiceUnavailable {

	return &UpdateTransitRouterServiceUnavailable{}
}

// WithPayload adds the payload to the update transit router service unavailable response
func (o *UpdateTransitRouterServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateTransitRouterServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update transit router service unavailable response
func (o *UpdateTransitRouterServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTransitRouterServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

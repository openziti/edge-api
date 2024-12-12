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

	"github.com/openziti/edge-api/rest_model"
)

// DeleteServiceOKCode is the HTTP code returned for type DeleteServiceOK
const DeleteServiceOKCode int = 200

/*DeleteServiceOK The delete request was successful and the resource has been removed

swagger:response deleteServiceOK
*/
type DeleteServiceOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteServiceOK creates DeleteServiceOK with default headers values
func NewDeleteServiceOK() *DeleteServiceOK {

	return &DeleteServiceOK{}
}

// WithPayload adds the payload to the delete service o k response
func (o *DeleteServiceOK) WithPayload(payload *rest_model.Empty) *DeleteServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service o k response
func (o *DeleteServiceOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceBadRequestCode is the HTTP code returned for type DeleteServiceBadRequest
const DeleteServiceBadRequestCode int = 400

/*DeleteServiceBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteServiceBadRequest
*/
type DeleteServiceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceBadRequest creates DeleteServiceBadRequest with default headers values
func NewDeleteServiceBadRequest() *DeleteServiceBadRequest {

	return &DeleteServiceBadRequest{}
}

// WithPayload adds the payload to the delete service bad request response
func (o *DeleteServiceBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service bad request response
func (o *DeleteServiceBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceUnauthorizedCode is the HTTP code returned for type DeleteServiceUnauthorized
const DeleteServiceUnauthorizedCode int = 401

/*DeleteServiceUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteServiceUnauthorized
*/
type DeleteServiceUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceUnauthorized creates DeleteServiceUnauthorized with default headers values
func NewDeleteServiceUnauthorized() *DeleteServiceUnauthorized {

	return &DeleteServiceUnauthorized{}
}

// WithPayload adds the payload to the delete service unauthorized response
func (o *DeleteServiceUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service unauthorized response
func (o *DeleteServiceUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceNotFoundCode is the HTTP code returned for type DeleteServiceNotFound
const DeleteServiceNotFoundCode int = 404

/*DeleteServiceNotFound The requested resource does not exist

swagger:response deleteServiceNotFound
*/
type DeleteServiceNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceNotFound creates DeleteServiceNotFound with default headers values
func NewDeleteServiceNotFound() *DeleteServiceNotFound {

	return &DeleteServiceNotFound{}
}

// WithPayload adds the payload to the delete service not found response
func (o *DeleteServiceNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service not found response
func (o *DeleteServiceNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceConflictCode is the HTTP code returned for type DeleteServiceConflict
const DeleteServiceConflictCode int = 409

/*DeleteServiceConflict The resource requested to be removed/altered cannot be as it is referenced by another object.

swagger:response deleteServiceConflict
*/
type DeleteServiceConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceConflict creates DeleteServiceConflict with default headers values
func NewDeleteServiceConflict() *DeleteServiceConflict {

	return &DeleteServiceConflict{}
}

// WithPayload adds the payload to the delete service conflict response
func (o *DeleteServiceConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service conflict response
func (o *DeleteServiceConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceTooManyRequestsCode is the HTTP code returned for type DeleteServiceTooManyRequests
const DeleteServiceTooManyRequestsCode int = 429

/*DeleteServiceTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteServiceTooManyRequests
*/
type DeleteServiceTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceTooManyRequests creates DeleteServiceTooManyRequests with default headers values
func NewDeleteServiceTooManyRequests() *DeleteServiceTooManyRequests {

	return &DeleteServiceTooManyRequests{}
}

// WithPayload adds the payload to the delete service too many requests response
func (o *DeleteServiceTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service too many requests response
func (o *DeleteServiceTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceServiceUnavailableCode is the HTTP code returned for type DeleteServiceServiceUnavailable
const DeleteServiceServiceUnavailableCode int = 503

/*DeleteServiceServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response deleteServiceServiceUnavailable
*/
type DeleteServiceServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceServiceUnavailable creates DeleteServiceServiceUnavailable with default headers values
func NewDeleteServiceServiceUnavailable() *DeleteServiceServiceUnavailable {

	return &DeleteServiceServiceUnavailable{}
}

// WithPayload adds the payload to the delete service service unavailable response
func (o *DeleteServiceServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service service unavailable response
func (o *DeleteServiceServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

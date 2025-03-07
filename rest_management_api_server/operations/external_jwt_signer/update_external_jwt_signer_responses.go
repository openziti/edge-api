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

	"github.com/openziti/edge-api/rest_model"
)

// UpdateExternalJWTSignerOKCode is the HTTP code returned for type UpdateExternalJWTSignerOK
const UpdateExternalJWTSignerOKCode int = 200

/*UpdateExternalJWTSignerOK The update request was successful and the resource has been altered

swagger:response updateExternalJwtSignerOK
*/
type UpdateExternalJWTSignerOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdateExternalJWTSignerOK creates UpdateExternalJWTSignerOK with default headers values
func NewUpdateExternalJWTSignerOK() *UpdateExternalJWTSignerOK {

	return &UpdateExternalJWTSignerOK{}
}

// WithPayload adds the payload to the update external Jwt signer o k response
func (o *UpdateExternalJWTSignerOK) WithPayload(payload *rest_model.Empty) *UpdateExternalJWTSignerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update external Jwt signer o k response
func (o *UpdateExternalJWTSignerOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExternalJWTSignerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateExternalJWTSignerBadRequestCode is the HTTP code returned for type UpdateExternalJWTSignerBadRequest
const UpdateExternalJWTSignerBadRequestCode int = 400

/*UpdateExternalJWTSignerBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updateExternalJwtSignerBadRequest
*/
type UpdateExternalJWTSignerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateExternalJWTSignerBadRequest creates UpdateExternalJWTSignerBadRequest with default headers values
func NewUpdateExternalJWTSignerBadRequest() *UpdateExternalJWTSignerBadRequest {

	return &UpdateExternalJWTSignerBadRequest{}
}

// WithPayload adds the payload to the update external Jwt signer bad request response
func (o *UpdateExternalJWTSignerBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateExternalJWTSignerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update external Jwt signer bad request response
func (o *UpdateExternalJWTSignerBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExternalJWTSignerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateExternalJWTSignerUnauthorizedCode is the HTTP code returned for type UpdateExternalJWTSignerUnauthorized
const UpdateExternalJWTSignerUnauthorizedCode int = 401

/*UpdateExternalJWTSignerUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response updateExternalJwtSignerUnauthorized
*/
type UpdateExternalJWTSignerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateExternalJWTSignerUnauthorized creates UpdateExternalJWTSignerUnauthorized with default headers values
func NewUpdateExternalJWTSignerUnauthorized() *UpdateExternalJWTSignerUnauthorized {

	return &UpdateExternalJWTSignerUnauthorized{}
}

// WithPayload adds the payload to the update external Jwt signer unauthorized response
func (o *UpdateExternalJWTSignerUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateExternalJWTSignerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update external Jwt signer unauthorized response
func (o *UpdateExternalJWTSignerUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExternalJWTSignerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateExternalJWTSignerNotFoundCode is the HTTP code returned for type UpdateExternalJWTSignerNotFound
const UpdateExternalJWTSignerNotFoundCode int = 404

/*UpdateExternalJWTSignerNotFound The requested resource does not exist

swagger:response updateExternalJwtSignerNotFound
*/
type UpdateExternalJWTSignerNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateExternalJWTSignerNotFound creates UpdateExternalJWTSignerNotFound with default headers values
func NewUpdateExternalJWTSignerNotFound() *UpdateExternalJWTSignerNotFound {

	return &UpdateExternalJWTSignerNotFound{}
}

// WithPayload adds the payload to the update external Jwt signer not found response
func (o *UpdateExternalJWTSignerNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateExternalJWTSignerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update external Jwt signer not found response
func (o *UpdateExternalJWTSignerNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExternalJWTSignerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateExternalJWTSignerTooManyRequestsCode is the HTTP code returned for type UpdateExternalJWTSignerTooManyRequests
const UpdateExternalJWTSignerTooManyRequestsCode int = 429

/*UpdateExternalJWTSignerTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response updateExternalJwtSignerTooManyRequests
*/
type UpdateExternalJWTSignerTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateExternalJWTSignerTooManyRequests creates UpdateExternalJWTSignerTooManyRequests with default headers values
func NewUpdateExternalJWTSignerTooManyRequests() *UpdateExternalJWTSignerTooManyRequests {

	return &UpdateExternalJWTSignerTooManyRequests{}
}

// WithPayload adds the payload to the update external Jwt signer too many requests response
func (o *UpdateExternalJWTSignerTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateExternalJWTSignerTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update external Jwt signer too many requests response
func (o *UpdateExternalJWTSignerTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExternalJWTSignerTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateExternalJWTSignerServiceUnavailableCode is the HTTP code returned for type UpdateExternalJWTSignerServiceUnavailable
const UpdateExternalJWTSignerServiceUnavailableCode int = 503

/*UpdateExternalJWTSignerServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response updateExternalJwtSignerServiceUnavailable
*/
type UpdateExternalJWTSignerServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateExternalJWTSignerServiceUnavailable creates UpdateExternalJWTSignerServiceUnavailable with default headers values
func NewUpdateExternalJWTSignerServiceUnavailable() *UpdateExternalJWTSignerServiceUnavailable {

	return &UpdateExternalJWTSignerServiceUnavailable{}
}

// WithPayload adds the payload to the update external Jwt signer service unavailable response
func (o *UpdateExternalJWTSignerServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateExternalJWTSignerServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update external Jwt signer service unavailable response
func (o *UpdateExternalJWTSignerServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExternalJWTSignerServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

	"github.com/openziti/edge-api/rest_model"
)

// CreateAuthenticatorCreatedCode is the HTTP code returned for type CreateAuthenticatorCreated
const CreateAuthenticatorCreatedCode int = 201

/*CreateAuthenticatorCreated The create request was successful and the resource has been added at the following location

swagger:response createAuthenticatorCreated
*/
type CreateAuthenticatorCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateAuthenticatorCreated creates CreateAuthenticatorCreated with default headers values
func NewCreateAuthenticatorCreated() *CreateAuthenticatorCreated {

	return &CreateAuthenticatorCreated{}
}

// WithPayload adds the payload to the create authenticator created response
func (o *CreateAuthenticatorCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateAuthenticatorCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authenticator created response
func (o *CreateAuthenticatorCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticatorCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticatorBadRequestCode is the HTTP code returned for type CreateAuthenticatorBadRequest
const CreateAuthenticatorBadRequestCode int = 400

/*CreateAuthenticatorBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createAuthenticatorBadRequest
*/
type CreateAuthenticatorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateAuthenticatorBadRequest creates CreateAuthenticatorBadRequest with default headers values
func NewCreateAuthenticatorBadRequest() *CreateAuthenticatorBadRequest {

	return &CreateAuthenticatorBadRequest{}
}

// WithPayload adds the payload to the create authenticator bad request response
func (o *CreateAuthenticatorBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateAuthenticatorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authenticator bad request response
func (o *CreateAuthenticatorBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticatorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticatorUnauthorizedCode is the HTTP code returned for type CreateAuthenticatorUnauthorized
const CreateAuthenticatorUnauthorizedCode int = 401

/*CreateAuthenticatorUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response createAuthenticatorUnauthorized
*/
type CreateAuthenticatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateAuthenticatorUnauthorized creates CreateAuthenticatorUnauthorized with default headers values
func NewCreateAuthenticatorUnauthorized() *CreateAuthenticatorUnauthorized {

	return &CreateAuthenticatorUnauthorized{}
}

// WithPayload adds the payload to the create authenticator unauthorized response
func (o *CreateAuthenticatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateAuthenticatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authenticator unauthorized response
func (o *CreateAuthenticatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticatorTooManyRequestsCode is the HTTP code returned for type CreateAuthenticatorTooManyRequests
const CreateAuthenticatorTooManyRequestsCode int = 429

/*CreateAuthenticatorTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createAuthenticatorTooManyRequests
*/
type CreateAuthenticatorTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateAuthenticatorTooManyRequests creates CreateAuthenticatorTooManyRequests with default headers values
func NewCreateAuthenticatorTooManyRequests() *CreateAuthenticatorTooManyRequests {

	return &CreateAuthenticatorTooManyRequests{}
}

// WithPayload adds the payload to the create authenticator too many requests response
func (o *CreateAuthenticatorTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateAuthenticatorTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authenticator too many requests response
func (o *CreateAuthenticatorTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticatorTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticatorServiceUnavailableCode is the HTTP code returned for type CreateAuthenticatorServiceUnavailable
const CreateAuthenticatorServiceUnavailableCode int = 503

/*CreateAuthenticatorServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response createAuthenticatorServiceUnavailable
*/
type CreateAuthenticatorServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateAuthenticatorServiceUnavailable creates CreateAuthenticatorServiceUnavailable with default headers values
func NewCreateAuthenticatorServiceUnavailable() *CreateAuthenticatorServiceUnavailable {

	return &CreateAuthenticatorServiceUnavailable{}
}

// WithPayload adds the payload to the create authenticator service unavailable response
func (o *CreateAuthenticatorServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateAuthenticatorServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authenticator service unavailable response
func (o *CreateAuthenticatorServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticatorServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

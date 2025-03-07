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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// CreateConfigTypeCreatedCode is the HTTP code returned for type CreateConfigTypeCreated
const CreateConfigTypeCreatedCode int = 201

/*CreateConfigTypeCreated The create request was successful and the resource has been added at the following location

swagger:response createConfigTypeCreated
*/
type CreateConfigTypeCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateConfigTypeCreated creates CreateConfigTypeCreated with default headers values
func NewCreateConfigTypeCreated() *CreateConfigTypeCreated {

	return &CreateConfigTypeCreated{}
}

// WithPayload adds the payload to the create config type created response
func (o *CreateConfigTypeCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateConfigTypeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config type created response
func (o *CreateConfigTypeCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigTypeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConfigTypeBadRequestCode is the HTTP code returned for type CreateConfigTypeBadRequest
const CreateConfigTypeBadRequestCode int = 400

/*CreateConfigTypeBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createConfigTypeBadRequest
*/
type CreateConfigTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateConfigTypeBadRequest creates CreateConfigTypeBadRequest with default headers values
func NewCreateConfigTypeBadRequest() *CreateConfigTypeBadRequest {

	return &CreateConfigTypeBadRequest{}
}

// WithPayload adds the payload to the create config type bad request response
func (o *CreateConfigTypeBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateConfigTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config type bad request response
func (o *CreateConfigTypeBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConfigTypeUnauthorizedCode is the HTTP code returned for type CreateConfigTypeUnauthorized
const CreateConfigTypeUnauthorizedCode int = 401

/*CreateConfigTypeUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response createConfigTypeUnauthorized
*/
type CreateConfigTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateConfigTypeUnauthorized creates CreateConfigTypeUnauthorized with default headers values
func NewCreateConfigTypeUnauthorized() *CreateConfigTypeUnauthorized {

	return &CreateConfigTypeUnauthorized{}
}

// WithPayload adds the payload to the create config type unauthorized response
func (o *CreateConfigTypeUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateConfigTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config type unauthorized response
func (o *CreateConfigTypeUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConfigTypeTooManyRequestsCode is the HTTP code returned for type CreateConfigTypeTooManyRequests
const CreateConfigTypeTooManyRequestsCode int = 429

/*CreateConfigTypeTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createConfigTypeTooManyRequests
*/
type CreateConfigTypeTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateConfigTypeTooManyRequests creates CreateConfigTypeTooManyRequests with default headers values
func NewCreateConfigTypeTooManyRequests() *CreateConfigTypeTooManyRequests {

	return &CreateConfigTypeTooManyRequests{}
}

// WithPayload adds the payload to the create config type too many requests response
func (o *CreateConfigTypeTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateConfigTypeTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config type too many requests response
func (o *CreateConfigTypeTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigTypeTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConfigTypeServiceUnavailableCode is the HTTP code returned for type CreateConfigTypeServiceUnavailable
const CreateConfigTypeServiceUnavailableCode int = 503

/*CreateConfigTypeServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response createConfigTypeServiceUnavailable
*/
type CreateConfigTypeServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateConfigTypeServiceUnavailable creates CreateConfigTypeServiceUnavailable with default headers values
func NewCreateConfigTypeServiceUnavailable() *CreateConfigTypeServiceUnavailable {

	return &CreateConfigTypeServiceUnavailable{}
}

// WithPayload adds the payload to the create config type service unavailable response
func (o *CreateConfigTypeServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateConfigTypeServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config type service unavailable response
func (o *CreateConfigTypeServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigTypeServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

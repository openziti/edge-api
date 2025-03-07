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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// GetCaJWTOKCode is the HTTP code returned for type GetCaJWTOK
const GetCaJWTOKCode int = 200

/*GetCaJWTOK The result is the JWT text to validate the CA

swagger:response getCaJwtOK
*/
type GetCaJWTOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetCaJWTOK creates GetCaJWTOK with default headers values
func NewGetCaJWTOK() *GetCaJWTOK {

	return &GetCaJWTOK{}
}

// WithPayload adds the payload to the get ca Jwt o k response
func (o *GetCaJWTOK) WithPayload(payload string) *GetCaJWTOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca Jwt o k response
func (o *GetCaJWTOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJWTOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCaJWTUnauthorizedCode is the HTTP code returned for type GetCaJWTUnauthorized
const GetCaJWTUnauthorizedCode int = 401

/*GetCaJWTUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response getCaJwtUnauthorized
*/
type GetCaJWTUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCaJWTUnauthorized creates GetCaJWTUnauthorized with default headers values
func NewGetCaJWTUnauthorized() *GetCaJWTUnauthorized {

	return &GetCaJWTUnauthorized{}
}

// WithPayload adds the payload to the get ca Jwt unauthorized response
func (o *GetCaJWTUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCaJWTUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca Jwt unauthorized response
func (o *GetCaJWTUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJWTUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCaJWTNotFoundCode is the HTTP code returned for type GetCaJWTNotFound
const GetCaJWTNotFoundCode int = 404

/*GetCaJWTNotFound The requested resource does not exist

swagger:response getCaJwtNotFound
*/
type GetCaJWTNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCaJWTNotFound creates GetCaJWTNotFound with default headers values
func NewGetCaJWTNotFound() *GetCaJWTNotFound {

	return &GetCaJWTNotFound{}
}

// WithPayload adds the payload to the get ca Jwt not found response
func (o *GetCaJWTNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCaJWTNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca Jwt not found response
func (o *GetCaJWTNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJWTNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCaJWTTooManyRequestsCode is the HTTP code returned for type GetCaJWTTooManyRequests
const GetCaJWTTooManyRequestsCode int = 429

/*GetCaJWTTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response getCaJwtTooManyRequests
*/
type GetCaJWTTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCaJWTTooManyRequests creates GetCaJWTTooManyRequests with default headers values
func NewGetCaJWTTooManyRequests() *GetCaJWTTooManyRequests {

	return &GetCaJWTTooManyRequests{}
}

// WithPayload adds the payload to the get ca Jwt too many requests response
func (o *GetCaJWTTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCaJWTTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca Jwt too many requests response
func (o *GetCaJWTTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJWTTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCaJWTServiceUnavailableCode is the HTTP code returned for type GetCaJWTServiceUnavailable
const GetCaJWTServiceUnavailableCode int = 503

/*GetCaJWTServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response getCaJwtServiceUnavailable
*/
type GetCaJWTServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCaJWTServiceUnavailable creates GetCaJWTServiceUnavailable with default headers values
func NewGetCaJWTServiceUnavailable() *GetCaJWTServiceUnavailable {

	return &GetCaJWTServiceUnavailable{}
}

// WithPayload adds the payload to the get ca Jwt service unavailable response
func (o *GetCaJWTServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCaJWTServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca Jwt service unavailable response
func (o *GetCaJWTServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJWTServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

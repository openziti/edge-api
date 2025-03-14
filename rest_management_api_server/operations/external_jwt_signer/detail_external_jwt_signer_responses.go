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

// DetailExternalJWTSignerOKCode is the HTTP code returned for type DetailExternalJWTSignerOK
const DetailExternalJWTSignerOKCode int = 200

/*DetailExternalJWTSignerOK A singular External JWT Signer resource

swagger:response detailExternalJwtSignerOK
*/
type DetailExternalJWTSignerOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailExternalJWTSignerEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJWTSignerOK creates DetailExternalJWTSignerOK with default headers values
func NewDetailExternalJWTSignerOK() *DetailExternalJWTSignerOK {

	return &DetailExternalJWTSignerOK{}
}

// WithPayload adds the payload to the detail external Jwt signer o k response
func (o *DetailExternalJWTSignerOK) WithPayload(payload *rest_model.DetailExternalJWTSignerEnvelope) *DetailExternalJWTSignerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external Jwt signer o k response
func (o *DetailExternalJWTSignerOK) SetPayload(payload *rest_model.DetailExternalJWTSignerEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJWTSignerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailExternalJWTSignerUnauthorizedCode is the HTTP code returned for type DetailExternalJWTSignerUnauthorized
const DetailExternalJWTSignerUnauthorizedCode int = 401

/*DetailExternalJWTSignerUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailExternalJwtSignerUnauthorized
*/
type DetailExternalJWTSignerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJWTSignerUnauthorized creates DetailExternalJWTSignerUnauthorized with default headers values
func NewDetailExternalJWTSignerUnauthorized() *DetailExternalJWTSignerUnauthorized {

	return &DetailExternalJWTSignerUnauthorized{}
}

// WithPayload adds the payload to the detail external Jwt signer unauthorized response
func (o *DetailExternalJWTSignerUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailExternalJWTSignerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external Jwt signer unauthorized response
func (o *DetailExternalJWTSignerUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJWTSignerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailExternalJWTSignerNotFoundCode is the HTTP code returned for type DetailExternalJWTSignerNotFound
const DetailExternalJWTSignerNotFoundCode int = 404

/*DetailExternalJWTSignerNotFound The requested resource does not exist

swagger:response detailExternalJwtSignerNotFound
*/
type DetailExternalJWTSignerNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJWTSignerNotFound creates DetailExternalJWTSignerNotFound with default headers values
func NewDetailExternalJWTSignerNotFound() *DetailExternalJWTSignerNotFound {

	return &DetailExternalJWTSignerNotFound{}
}

// WithPayload adds the payload to the detail external Jwt signer not found response
func (o *DetailExternalJWTSignerNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailExternalJWTSignerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external Jwt signer not found response
func (o *DetailExternalJWTSignerNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJWTSignerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailExternalJWTSignerTooManyRequestsCode is the HTTP code returned for type DetailExternalJWTSignerTooManyRequests
const DetailExternalJWTSignerTooManyRequestsCode int = 429

/*DetailExternalJWTSignerTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailExternalJwtSignerTooManyRequests
*/
type DetailExternalJWTSignerTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJWTSignerTooManyRequests creates DetailExternalJWTSignerTooManyRequests with default headers values
func NewDetailExternalJWTSignerTooManyRequests() *DetailExternalJWTSignerTooManyRequests {

	return &DetailExternalJWTSignerTooManyRequests{}
}

// WithPayload adds the payload to the detail external Jwt signer too many requests response
func (o *DetailExternalJWTSignerTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailExternalJWTSignerTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external Jwt signer too many requests response
func (o *DetailExternalJWTSignerTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJWTSignerTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailExternalJWTSignerServiceUnavailableCode is the HTTP code returned for type DetailExternalJWTSignerServiceUnavailable
const DetailExternalJWTSignerServiceUnavailableCode int = 503

/*DetailExternalJWTSignerServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response detailExternalJwtSignerServiceUnavailable
*/
type DetailExternalJWTSignerServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJWTSignerServiceUnavailable creates DetailExternalJWTSignerServiceUnavailable with default headers values
func NewDetailExternalJWTSignerServiceUnavailable() *DetailExternalJWTSignerServiceUnavailable {

	return &DetailExternalJWTSignerServiceUnavailable{}
}

// WithPayload adds the payload to the detail external Jwt signer service unavailable response
func (o *DetailExternalJWTSignerServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailExternalJWTSignerServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external Jwt signer service unavailable response
func (o *DetailExternalJWTSignerServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJWTSignerServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

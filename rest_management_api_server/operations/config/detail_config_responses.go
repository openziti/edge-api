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

// DetailConfigOKCode is the HTTP code returned for type DetailConfigOK
const DetailConfigOKCode int = 200

/*DetailConfigOK A singular config resource

swagger:response detailConfigOK
*/
type DetailConfigOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailConfigEnvelope `json:"body,omitempty"`
}

// NewDetailConfigOK creates DetailConfigOK with default headers values
func NewDetailConfigOK() *DetailConfigOK {

	return &DetailConfigOK{}
}

// WithPayload adds the payload to the detail config o k response
func (o *DetailConfigOK) WithPayload(payload *rest_model.DetailConfigEnvelope) *DetailConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail config o k response
func (o *DetailConfigOK) SetPayload(payload *rest_model.DetailConfigEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailConfigUnauthorizedCode is the HTTP code returned for type DetailConfigUnauthorized
const DetailConfigUnauthorizedCode int = 401

/*DetailConfigUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailConfigUnauthorized
*/
type DetailConfigUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailConfigUnauthorized creates DetailConfigUnauthorized with default headers values
func NewDetailConfigUnauthorized() *DetailConfigUnauthorized {

	return &DetailConfigUnauthorized{}
}

// WithPayload adds the payload to the detail config unauthorized response
func (o *DetailConfigUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailConfigUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail config unauthorized response
func (o *DetailConfigUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailConfigUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailConfigNotFoundCode is the HTTP code returned for type DetailConfigNotFound
const DetailConfigNotFoundCode int = 404

/*DetailConfigNotFound The requested resource does not exist

swagger:response detailConfigNotFound
*/
type DetailConfigNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailConfigNotFound creates DetailConfigNotFound with default headers values
func NewDetailConfigNotFound() *DetailConfigNotFound {

	return &DetailConfigNotFound{}
}

// WithPayload adds the payload to the detail config not found response
func (o *DetailConfigNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailConfigNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail config not found response
func (o *DetailConfigNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailConfigTooManyRequestsCode is the HTTP code returned for type DetailConfigTooManyRequests
const DetailConfigTooManyRequestsCode int = 429

/*DetailConfigTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailConfigTooManyRequests
*/
type DetailConfigTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailConfigTooManyRequests creates DetailConfigTooManyRequests with default headers values
func NewDetailConfigTooManyRequests() *DetailConfigTooManyRequests {

	return &DetailConfigTooManyRequests{}
}

// WithPayload adds the payload to the detail config too many requests response
func (o *DetailConfigTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailConfigTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail config too many requests response
func (o *DetailConfigTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailConfigTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailConfigServiceUnavailableCode is the HTTP code returned for type DetailConfigServiceUnavailable
const DetailConfigServiceUnavailableCode int = 503

/*DetailConfigServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response detailConfigServiceUnavailable
*/
type DetailConfigServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailConfigServiceUnavailable creates DetailConfigServiceUnavailable with default headers values
func NewDetailConfigServiceUnavailable() *DetailConfigServiceUnavailable {

	return &DetailConfigServiceUnavailable{}
}

// WithPayload adds the payload to the detail config service unavailable response
func (o *DetailConfigServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailConfigServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail config service unavailable response
func (o *DetailConfigServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailConfigServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

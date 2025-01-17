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

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// EnrollOttOKCode is the HTTP code returned for type EnrollOttOK
const EnrollOttOKCode int = 200

/*EnrollOttOK A response containing and identities client certificate chains

swagger:response enrollOttOK
*/
type EnrollOttOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.EnrollmentCertsEnvelope `json:"body,omitempty"`
}

// NewEnrollOttOK creates EnrollOttOK with default headers values
func NewEnrollOttOK() *EnrollOttOK {

	return &EnrollOttOK{}
}

// WithPayload adds the payload to the enroll ott o k response
func (o *EnrollOttOK) WithPayload(payload *rest_model.EnrollmentCertsEnvelope) *EnrollOttOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ott o k response
func (o *EnrollOttOK) SetPayload(payload *rest_model.EnrollmentCertsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollOttOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollOttNotFoundCode is the HTTP code returned for type EnrollOttNotFound
const EnrollOttNotFoundCode int = 404

/*EnrollOttNotFound The requested resource does not exist

swagger:response enrollOttNotFound
*/
type EnrollOttNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollOttNotFound creates EnrollOttNotFound with default headers values
func NewEnrollOttNotFound() *EnrollOttNotFound {

	return &EnrollOttNotFound{}
}

// WithPayload adds the payload to the enroll ott not found response
func (o *EnrollOttNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollOttNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ott not found response
func (o *EnrollOttNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollOttNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollOttTooManyRequestsCode is the HTTP code returned for type EnrollOttTooManyRequests
const EnrollOttTooManyRequestsCode int = 429

/*EnrollOttTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response enrollOttTooManyRequests
*/
type EnrollOttTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollOttTooManyRequests creates EnrollOttTooManyRequests with default headers values
func NewEnrollOttTooManyRequests() *EnrollOttTooManyRequests {

	return &EnrollOttTooManyRequests{}
}

// WithPayload adds the payload to the enroll ott too many requests response
func (o *EnrollOttTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollOttTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ott too many requests response
func (o *EnrollOttTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollOttTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollOttServiceUnavailableCode is the HTTP code returned for type EnrollOttServiceUnavailable
const EnrollOttServiceUnavailableCode int = 503

/*EnrollOttServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response enrollOttServiceUnavailable
*/
type EnrollOttServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollOttServiceUnavailable creates EnrollOttServiceUnavailable with default headers values
func NewEnrollOttServiceUnavailable() *EnrollOttServiceUnavailable {

	return &EnrollOttServiceUnavailable{}
}

// WithPayload adds the payload to the enroll ott service unavailable response
func (o *EnrollOttServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollOttServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ott service unavailable response
func (o *EnrollOttServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollOttServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

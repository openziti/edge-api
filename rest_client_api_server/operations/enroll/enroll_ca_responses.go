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

// EnrollCaOKCode is the HTTP code returned for type EnrollCaOK
const EnrollCaOKCode int = 200

/*EnrollCaOK Base empty response

swagger:response enrollCaOK
*/
type EnrollCaOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewEnrollCaOK creates EnrollCaOK with default headers values
func NewEnrollCaOK() *EnrollCaOK {

	return &EnrollCaOK{}
}

// WithPayload adds the payload to the enroll ca o k response
func (o *EnrollCaOK) WithPayload(payload *rest_model.Empty) *EnrollCaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ca o k response
func (o *EnrollCaOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollCaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollCaNotFoundCode is the HTTP code returned for type EnrollCaNotFound
const EnrollCaNotFoundCode int = 404

/*EnrollCaNotFound The requested resource does not exist

swagger:response enrollCaNotFound
*/
type EnrollCaNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollCaNotFound creates EnrollCaNotFound with default headers values
func NewEnrollCaNotFound() *EnrollCaNotFound {

	return &EnrollCaNotFound{}
}

// WithPayload adds the payload to the enroll ca not found response
func (o *EnrollCaNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollCaNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ca not found response
func (o *EnrollCaNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollCaNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollCaTooManyRequestsCode is the HTTP code returned for type EnrollCaTooManyRequests
const EnrollCaTooManyRequestsCode int = 429

/*EnrollCaTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response enrollCaTooManyRequests
*/
type EnrollCaTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollCaTooManyRequests creates EnrollCaTooManyRequests with default headers values
func NewEnrollCaTooManyRequests() *EnrollCaTooManyRequests {

	return &EnrollCaTooManyRequests{}
}

// WithPayload adds the payload to the enroll ca too many requests response
func (o *EnrollCaTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollCaTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll ca too many requests response
func (o *EnrollCaTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollCaTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

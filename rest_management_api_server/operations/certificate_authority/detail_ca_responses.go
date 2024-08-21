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

// DetailCaOKCode is the HTTP code returned for type DetailCaOK
const DetailCaOKCode int = 200

/*DetailCaOK A singular Certificate Authority (CA) resource

swagger:response detailCaOK
*/
type DetailCaOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailCaEnvelope `json:"body,omitempty"`
}

// NewDetailCaOK creates DetailCaOK with default headers values
func NewDetailCaOK() *DetailCaOK {

	return &DetailCaOK{}
}

// WithPayload adds the payload to the detail ca o k response
func (o *DetailCaOK) WithPayload(payload *rest_model.DetailCaEnvelope) *DetailCaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail ca o k response
func (o *DetailCaOK) SetPayload(payload *rest_model.DetailCaEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCaUnauthorizedCode is the HTTP code returned for type DetailCaUnauthorized
const DetailCaUnauthorizedCode int = 401

/*DetailCaUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailCaUnauthorized
*/
type DetailCaUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCaUnauthorized creates DetailCaUnauthorized with default headers values
func NewDetailCaUnauthorized() *DetailCaUnauthorized {

	return &DetailCaUnauthorized{}
}

// WithPayload adds the payload to the detail ca unauthorized response
func (o *DetailCaUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCaUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail ca unauthorized response
func (o *DetailCaUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCaNotFoundCode is the HTTP code returned for type DetailCaNotFound
const DetailCaNotFoundCode int = 404

/*DetailCaNotFound The requested resource does not exist

swagger:response detailCaNotFound
*/
type DetailCaNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCaNotFound creates DetailCaNotFound with default headers values
func NewDetailCaNotFound() *DetailCaNotFound {

	return &DetailCaNotFound{}
}

// WithPayload adds the payload to the detail ca not found response
func (o *DetailCaNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCaNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail ca not found response
func (o *DetailCaNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCaNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCaTooManyRequestsCode is the HTTP code returned for type DetailCaTooManyRequests
const DetailCaTooManyRequestsCode int = 429

/*DetailCaTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailCaTooManyRequests
*/
type DetailCaTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCaTooManyRequests creates DetailCaTooManyRequests with default headers values
func NewDetailCaTooManyRequests() *DetailCaTooManyRequests {

	return &DetailCaTooManyRequests{}
}

// WithPayload adds the payload to the detail ca too many requests response
func (o *DetailCaTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCaTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail ca too many requests response
func (o *DetailCaTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCaTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

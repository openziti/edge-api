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

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteEnrollmentOKCode is the HTTP code returned for type DeleteEnrollmentOK
const DeleteEnrollmentOKCode int = 200

/*DeleteEnrollmentOK The delete request was successful and the resource has been removed

swagger:response deleteEnrollmentOK
*/
type DeleteEnrollmentOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteEnrollmentOK creates DeleteEnrollmentOK with default headers values
func NewDeleteEnrollmentOK() *DeleteEnrollmentOK {

	return &DeleteEnrollmentOK{}
}

// WithPayload adds the payload to the delete enrollment o k response
func (o *DeleteEnrollmentOK) WithPayload(payload *rest_model.Empty) *DeleteEnrollmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete enrollment o k response
func (o *DeleteEnrollmentOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEnrollmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEnrollmentBadRequestCode is the HTTP code returned for type DeleteEnrollmentBadRequest
const DeleteEnrollmentBadRequestCode int = 400

/*DeleteEnrollmentBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteEnrollmentBadRequest
*/
type DeleteEnrollmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEnrollmentBadRequest creates DeleteEnrollmentBadRequest with default headers values
func NewDeleteEnrollmentBadRequest() *DeleteEnrollmentBadRequest {

	return &DeleteEnrollmentBadRequest{}
}

// WithPayload adds the payload to the delete enrollment bad request response
func (o *DeleteEnrollmentBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEnrollmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete enrollment bad request response
func (o *DeleteEnrollmentBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEnrollmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEnrollmentUnauthorizedCode is the HTTP code returned for type DeleteEnrollmentUnauthorized
const DeleteEnrollmentUnauthorizedCode int = 401

/*DeleteEnrollmentUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteEnrollmentUnauthorized
*/
type DeleteEnrollmentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEnrollmentUnauthorized creates DeleteEnrollmentUnauthorized with default headers values
func NewDeleteEnrollmentUnauthorized() *DeleteEnrollmentUnauthorized {

	return &DeleteEnrollmentUnauthorized{}
}

// WithPayload adds the payload to the delete enrollment unauthorized response
func (o *DeleteEnrollmentUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEnrollmentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete enrollment unauthorized response
func (o *DeleteEnrollmentUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEnrollmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEnrollmentNotFoundCode is the HTTP code returned for type DeleteEnrollmentNotFound
const DeleteEnrollmentNotFoundCode int = 404

/*DeleteEnrollmentNotFound The requested resource does not exist

swagger:response deleteEnrollmentNotFound
*/
type DeleteEnrollmentNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEnrollmentNotFound creates DeleteEnrollmentNotFound with default headers values
func NewDeleteEnrollmentNotFound() *DeleteEnrollmentNotFound {

	return &DeleteEnrollmentNotFound{}
}

// WithPayload adds the payload to the delete enrollment not found response
func (o *DeleteEnrollmentNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEnrollmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete enrollment not found response
func (o *DeleteEnrollmentNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEnrollmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEnrollmentTooManyRequestsCode is the HTTP code returned for type DeleteEnrollmentTooManyRequests
const DeleteEnrollmentTooManyRequestsCode int = 429

/*DeleteEnrollmentTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteEnrollmentTooManyRequests
*/
type DeleteEnrollmentTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEnrollmentTooManyRequests creates DeleteEnrollmentTooManyRequests with default headers values
func NewDeleteEnrollmentTooManyRequests() *DeleteEnrollmentTooManyRequests {

	return &DeleteEnrollmentTooManyRequests{}
}

// WithPayload adds the payload to the delete enrollment too many requests response
func (o *DeleteEnrollmentTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEnrollmentTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete enrollment too many requests response
func (o *DeleteEnrollmentTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEnrollmentTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

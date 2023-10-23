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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// UpdatePostureCheckOKCode is the HTTP code returned for type UpdatePostureCheckOK
const UpdatePostureCheckOKCode int = 200

/*UpdatePostureCheckOK The update request was successful and the resource has been altered

swagger:response updatePostureCheckOK
*/
type UpdatePostureCheckOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdatePostureCheckOK creates UpdatePostureCheckOK with default headers values
func NewUpdatePostureCheckOK() *UpdatePostureCheckOK {

	return &UpdatePostureCheckOK{}
}

// WithPayload adds the payload to the update posture check o k response
func (o *UpdatePostureCheckOK) WithPayload(payload *rest_model.Empty) *UpdatePostureCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update posture check o k response
func (o *UpdatePostureCheckOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePostureCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePostureCheckBadRequestCode is the HTTP code returned for type UpdatePostureCheckBadRequest
const UpdatePostureCheckBadRequestCode int = 400

/*UpdatePostureCheckBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updatePostureCheckBadRequest
*/
type UpdatePostureCheckBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdatePostureCheckBadRequest creates UpdatePostureCheckBadRequest with default headers values
func NewUpdatePostureCheckBadRequest() *UpdatePostureCheckBadRequest {

	return &UpdatePostureCheckBadRequest{}
}

// WithPayload adds the payload to the update posture check bad request response
func (o *UpdatePostureCheckBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdatePostureCheckBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update posture check bad request response
func (o *UpdatePostureCheckBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePostureCheckBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePostureCheckUnauthorizedCode is the HTTP code returned for type UpdatePostureCheckUnauthorized
const UpdatePostureCheckUnauthorizedCode int = 401

/*UpdatePostureCheckUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response updatePostureCheckUnauthorized
*/
type UpdatePostureCheckUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdatePostureCheckUnauthorized creates UpdatePostureCheckUnauthorized with default headers values
func NewUpdatePostureCheckUnauthorized() *UpdatePostureCheckUnauthorized {

	return &UpdatePostureCheckUnauthorized{}
}

// WithPayload adds the payload to the update posture check unauthorized response
func (o *UpdatePostureCheckUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdatePostureCheckUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update posture check unauthorized response
func (o *UpdatePostureCheckUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePostureCheckUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePostureCheckNotFoundCode is the HTTP code returned for type UpdatePostureCheckNotFound
const UpdatePostureCheckNotFoundCode int = 404

/*UpdatePostureCheckNotFound The requested resource does not exist

swagger:response updatePostureCheckNotFound
*/
type UpdatePostureCheckNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdatePostureCheckNotFound creates UpdatePostureCheckNotFound with default headers values
func NewUpdatePostureCheckNotFound() *UpdatePostureCheckNotFound {

	return &UpdatePostureCheckNotFound{}
}

// WithPayload adds the payload to the update posture check not found response
func (o *UpdatePostureCheckNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdatePostureCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update posture check not found response
func (o *UpdatePostureCheckNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePostureCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePostureCheckTooManyRequestsCode is the HTTP code returned for type UpdatePostureCheckTooManyRequests
const UpdatePostureCheckTooManyRequestsCode int = 429

/*UpdatePostureCheckTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response updatePostureCheckTooManyRequests
*/
type UpdatePostureCheckTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdatePostureCheckTooManyRequests creates UpdatePostureCheckTooManyRequests with default headers values
func NewUpdatePostureCheckTooManyRequests() *UpdatePostureCheckTooManyRequests {

	return &UpdatePostureCheckTooManyRequests{}
}

// WithPayload adds the payload to the update posture check too many requests response
func (o *UpdatePostureCheckTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdatePostureCheckTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update posture check too many requests response
func (o *UpdatePostureCheckTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePostureCheckTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

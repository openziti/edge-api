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

// ListPostureCheckTypesOKCode is the HTTP code returned for type ListPostureCheckTypesOK
const ListPostureCheckTypesOKCode int = 200

/*ListPostureCheckTypesOK A list of posture check types

swagger:response listPostureCheckTypesOK
*/
type ListPostureCheckTypesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListPostureCheckTypesEnvelope `json:"body,omitempty"`
}

// NewListPostureCheckTypesOK creates ListPostureCheckTypesOK with default headers values
func NewListPostureCheckTypesOK() *ListPostureCheckTypesOK {

	return &ListPostureCheckTypesOK{}
}

// WithPayload adds the payload to the list posture check types o k response
func (o *ListPostureCheckTypesOK) WithPayload(payload *rest_model.ListPostureCheckTypesEnvelope) *ListPostureCheckTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list posture check types o k response
func (o *ListPostureCheckTypesOK) SetPayload(payload *rest_model.ListPostureCheckTypesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPostureCheckTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPostureCheckTypesBadRequestCode is the HTTP code returned for type ListPostureCheckTypesBadRequest
const ListPostureCheckTypesBadRequestCode int = 400

/*ListPostureCheckTypesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listPostureCheckTypesBadRequest
*/
type ListPostureCheckTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListPostureCheckTypesBadRequest creates ListPostureCheckTypesBadRequest with default headers values
func NewListPostureCheckTypesBadRequest() *ListPostureCheckTypesBadRequest {

	return &ListPostureCheckTypesBadRequest{}
}

// WithPayload adds the payload to the list posture check types bad request response
func (o *ListPostureCheckTypesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListPostureCheckTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list posture check types bad request response
func (o *ListPostureCheckTypesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPostureCheckTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPostureCheckTypesUnauthorizedCode is the HTTP code returned for type ListPostureCheckTypesUnauthorized
const ListPostureCheckTypesUnauthorizedCode int = 401

/*ListPostureCheckTypesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listPostureCheckTypesUnauthorized
*/
type ListPostureCheckTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListPostureCheckTypesUnauthorized creates ListPostureCheckTypesUnauthorized with default headers values
func NewListPostureCheckTypesUnauthorized() *ListPostureCheckTypesUnauthorized {

	return &ListPostureCheckTypesUnauthorized{}
}

// WithPayload adds the payload to the list posture check types unauthorized response
func (o *ListPostureCheckTypesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListPostureCheckTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list posture check types unauthorized response
func (o *ListPostureCheckTypesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPostureCheckTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPostureCheckTypesTooManyRequestsCode is the HTTP code returned for type ListPostureCheckTypesTooManyRequests
const ListPostureCheckTypesTooManyRequestsCode int = 429

/*ListPostureCheckTypesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listPostureCheckTypesTooManyRequests
*/
type ListPostureCheckTypesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListPostureCheckTypesTooManyRequests creates ListPostureCheckTypesTooManyRequests with default headers values
func NewListPostureCheckTypesTooManyRequests() *ListPostureCheckTypesTooManyRequests {

	return &ListPostureCheckTypesTooManyRequests{}
}

// WithPayload adds the payload to the list posture check types too many requests response
func (o *ListPostureCheckTypesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListPostureCheckTypesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list posture check types too many requests response
func (o *ListPostureCheckTypesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPostureCheckTypesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

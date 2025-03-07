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

package informational

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListSummaryOKCode is the HTTP code returned for type ListSummaryOK
const ListSummaryOKCode int = 200

/*ListSummaryOK Entity counts scopped to the current identitie's access

swagger:response listSummaryOK
*/
type ListSummaryOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListSummaryCountsEnvelope `json:"body,omitempty"`
}

// NewListSummaryOK creates ListSummaryOK with default headers values
func NewListSummaryOK() *ListSummaryOK {

	return &ListSummaryOK{}
}

// WithPayload adds the payload to the list summary o k response
func (o *ListSummaryOK) WithPayload(payload *rest_model.ListSummaryCountsEnvelope) *ListSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list summary o k response
func (o *ListSummaryOK) SetPayload(payload *rest_model.ListSummaryCountsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSummaryUnauthorizedCode is the HTTP code returned for type ListSummaryUnauthorized
const ListSummaryUnauthorizedCode int = 401

/*ListSummaryUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listSummaryUnauthorized
*/
type ListSummaryUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListSummaryUnauthorized creates ListSummaryUnauthorized with default headers values
func NewListSummaryUnauthorized() *ListSummaryUnauthorized {

	return &ListSummaryUnauthorized{}
}

// WithPayload adds the payload to the list summary unauthorized response
func (o *ListSummaryUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListSummaryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list summary unauthorized response
func (o *ListSummaryUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSummaryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSummaryTooManyRequestsCode is the HTTP code returned for type ListSummaryTooManyRequests
const ListSummaryTooManyRequestsCode int = 429

/*ListSummaryTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listSummaryTooManyRequests
*/
type ListSummaryTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListSummaryTooManyRequests creates ListSummaryTooManyRequests with default headers values
func NewListSummaryTooManyRequests() *ListSummaryTooManyRequests {

	return &ListSummaryTooManyRequests{}
}

// WithPayload adds the payload to the list summary too many requests response
func (o *ListSummaryTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListSummaryTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list summary too many requests response
func (o *ListSummaryTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSummaryTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSummaryServiceUnavailableCode is the HTTP code returned for type ListSummaryServiceUnavailable
const ListSummaryServiceUnavailableCode int = 503

/*ListSummaryServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response listSummaryServiceUnavailable
*/
type ListSummaryServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListSummaryServiceUnavailable creates ListSummaryServiceUnavailable with default headers values
func NewListSummaryServiceUnavailable() *ListSummaryServiceUnavailable {

	return &ListSummaryServiceUnavailable{}
}

// WithPayload adds the payload to the list summary service unavailable response
func (o *ListSummaryServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *ListSummaryServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list summary service unavailable response
func (o *ListSummaryServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSummaryServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

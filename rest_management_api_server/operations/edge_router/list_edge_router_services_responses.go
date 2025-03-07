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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListEdgeRouterServicesOKCode is the HTTP code returned for type ListEdgeRouterServicesOK
const ListEdgeRouterServicesOKCode int = 200

/*ListEdgeRouterServicesOK A list of services

swagger:response listEdgeRouterServicesOK
*/
type ListEdgeRouterServicesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServicesEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServicesOK creates ListEdgeRouterServicesOK with default headers values
func NewListEdgeRouterServicesOK() *ListEdgeRouterServicesOK {

	return &ListEdgeRouterServicesOK{}
}

// WithPayload adds the payload to the list edge router services o k response
func (o *ListEdgeRouterServicesOK) WithPayload(payload *rest_model.ListServicesEnvelope) *ListEdgeRouterServicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router services o k response
func (o *ListEdgeRouterServicesOK) SetPayload(payload *rest_model.ListServicesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterServicesUnauthorizedCode is the HTTP code returned for type ListEdgeRouterServicesUnauthorized
const ListEdgeRouterServicesUnauthorizedCode int = 401

/*ListEdgeRouterServicesUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listEdgeRouterServicesUnauthorized
*/
type ListEdgeRouterServicesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServicesUnauthorized creates ListEdgeRouterServicesUnauthorized with default headers values
func NewListEdgeRouterServicesUnauthorized() *ListEdgeRouterServicesUnauthorized {

	return &ListEdgeRouterServicesUnauthorized{}
}

// WithPayload adds the payload to the list edge router services unauthorized response
func (o *ListEdgeRouterServicesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterServicesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router services unauthorized response
func (o *ListEdgeRouterServicesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServicesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterServicesNotFoundCode is the HTTP code returned for type ListEdgeRouterServicesNotFound
const ListEdgeRouterServicesNotFoundCode int = 404

/*ListEdgeRouterServicesNotFound The requested resource does not exist

swagger:response listEdgeRouterServicesNotFound
*/
type ListEdgeRouterServicesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServicesNotFound creates ListEdgeRouterServicesNotFound with default headers values
func NewListEdgeRouterServicesNotFound() *ListEdgeRouterServicesNotFound {

	return &ListEdgeRouterServicesNotFound{}
}

// WithPayload adds the payload to the list edge router services not found response
func (o *ListEdgeRouterServicesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterServicesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router services not found response
func (o *ListEdgeRouterServicesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServicesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterServicesTooManyRequestsCode is the HTTP code returned for type ListEdgeRouterServicesTooManyRequests
const ListEdgeRouterServicesTooManyRequestsCode int = 429

/*ListEdgeRouterServicesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listEdgeRouterServicesTooManyRequests
*/
type ListEdgeRouterServicesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServicesTooManyRequests creates ListEdgeRouterServicesTooManyRequests with default headers values
func NewListEdgeRouterServicesTooManyRequests() *ListEdgeRouterServicesTooManyRequests {

	return &ListEdgeRouterServicesTooManyRequests{}
}

// WithPayload adds the payload to the list edge router services too many requests response
func (o *ListEdgeRouterServicesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterServicesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router services too many requests response
func (o *ListEdgeRouterServicesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServicesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterServicesServiceUnavailableCode is the HTTP code returned for type ListEdgeRouterServicesServiceUnavailable
const ListEdgeRouterServicesServiceUnavailableCode int = 503

/*ListEdgeRouterServicesServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response listEdgeRouterServicesServiceUnavailable
*/
type ListEdgeRouterServicesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServicesServiceUnavailable creates ListEdgeRouterServicesServiceUnavailable with default headers values
func NewListEdgeRouterServicesServiceUnavailable() *ListEdgeRouterServicesServiceUnavailable {

	return &ListEdgeRouterServicesServiceUnavailable{}
}

// WithPayload adds the payload to the list edge router services service unavailable response
func (o *ListEdgeRouterServicesServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterServicesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router services service unavailable response
func (o *ListEdgeRouterServicesServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServicesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

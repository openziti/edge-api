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

package service_edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListServiceEdgeRouterPolicyServicesOKCode is the HTTP code returned for type ListServiceEdgeRouterPolicyServicesOK
const ListServiceEdgeRouterPolicyServicesOKCode int = 200

/*ListServiceEdgeRouterPolicyServicesOK A list of services

swagger:response listServiceEdgeRouterPolicyServicesOK
*/
type ListServiceEdgeRouterPolicyServicesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServicesEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPolicyServicesOK creates ListServiceEdgeRouterPolicyServicesOK with default headers values
func NewListServiceEdgeRouterPolicyServicesOK() *ListServiceEdgeRouterPolicyServicesOK {

	return &ListServiceEdgeRouterPolicyServicesOK{}
}

// WithPayload adds the payload to the list service edge router policy services o k response
func (o *ListServiceEdgeRouterPolicyServicesOK) WithPayload(payload *rest_model.ListServicesEnvelope) *ListServiceEdgeRouterPolicyServicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policy services o k response
func (o *ListServiceEdgeRouterPolicyServicesOK) SetPayload(payload *rest_model.ListServicesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPolicyServicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPolicyServicesUnauthorizedCode is the HTTP code returned for type ListServiceEdgeRouterPolicyServicesUnauthorized
const ListServiceEdgeRouterPolicyServicesUnauthorizedCode int = 401

/*ListServiceEdgeRouterPolicyServicesUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listServiceEdgeRouterPolicyServicesUnauthorized
*/
type ListServiceEdgeRouterPolicyServicesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPolicyServicesUnauthorized creates ListServiceEdgeRouterPolicyServicesUnauthorized with default headers values
func NewListServiceEdgeRouterPolicyServicesUnauthorized() *ListServiceEdgeRouterPolicyServicesUnauthorized {

	return &ListServiceEdgeRouterPolicyServicesUnauthorized{}
}

// WithPayload adds the payload to the list service edge router policy services unauthorized response
func (o *ListServiceEdgeRouterPolicyServicesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPolicyServicesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policy services unauthorized response
func (o *ListServiceEdgeRouterPolicyServicesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPolicyServicesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPolicyServicesNotFoundCode is the HTTP code returned for type ListServiceEdgeRouterPolicyServicesNotFound
const ListServiceEdgeRouterPolicyServicesNotFoundCode int = 404

/*ListServiceEdgeRouterPolicyServicesNotFound The requested resource does not exist

swagger:response listServiceEdgeRouterPolicyServicesNotFound
*/
type ListServiceEdgeRouterPolicyServicesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPolicyServicesNotFound creates ListServiceEdgeRouterPolicyServicesNotFound with default headers values
func NewListServiceEdgeRouterPolicyServicesNotFound() *ListServiceEdgeRouterPolicyServicesNotFound {

	return &ListServiceEdgeRouterPolicyServicesNotFound{}
}

// WithPayload adds the payload to the list service edge router policy services not found response
func (o *ListServiceEdgeRouterPolicyServicesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPolicyServicesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policy services not found response
func (o *ListServiceEdgeRouterPolicyServicesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPolicyServicesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPolicyServicesTooManyRequestsCode is the HTTP code returned for type ListServiceEdgeRouterPolicyServicesTooManyRequests
const ListServiceEdgeRouterPolicyServicesTooManyRequestsCode int = 429

/*ListServiceEdgeRouterPolicyServicesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listServiceEdgeRouterPolicyServicesTooManyRequests
*/
type ListServiceEdgeRouterPolicyServicesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPolicyServicesTooManyRequests creates ListServiceEdgeRouterPolicyServicesTooManyRequests with default headers values
func NewListServiceEdgeRouterPolicyServicesTooManyRequests() *ListServiceEdgeRouterPolicyServicesTooManyRequests {

	return &ListServiceEdgeRouterPolicyServicesTooManyRequests{}
}

// WithPayload adds the payload to the list service edge router policy services too many requests response
func (o *ListServiceEdgeRouterPolicyServicesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPolicyServicesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policy services too many requests response
func (o *ListServiceEdgeRouterPolicyServicesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPolicyServicesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPolicyServicesServiceUnavailableCode is the HTTP code returned for type ListServiceEdgeRouterPolicyServicesServiceUnavailable
const ListServiceEdgeRouterPolicyServicesServiceUnavailableCode int = 503

/*ListServiceEdgeRouterPolicyServicesServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response listServiceEdgeRouterPolicyServicesServiceUnavailable
*/
type ListServiceEdgeRouterPolicyServicesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPolicyServicesServiceUnavailable creates ListServiceEdgeRouterPolicyServicesServiceUnavailable with default headers values
func NewListServiceEdgeRouterPolicyServicesServiceUnavailable() *ListServiceEdgeRouterPolicyServicesServiceUnavailable {

	return &ListServiceEdgeRouterPolicyServicesServiceUnavailable{}
}

// WithPayload adds the payload to the list service edge router policy services service unavailable response
func (o *ListServiceEdgeRouterPolicyServicesServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPolicyServicesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policy services service unavailable response
func (o *ListServiceEdgeRouterPolicyServicesServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPolicyServicesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

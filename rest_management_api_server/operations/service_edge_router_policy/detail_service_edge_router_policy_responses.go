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

// DetailServiceEdgeRouterPolicyOKCode is the HTTP code returned for type DetailServiceEdgeRouterPolicyOK
const DetailServiceEdgeRouterPolicyOKCode int = 200

/*DetailServiceEdgeRouterPolicyOK A single service edge router policy

swagger:response detailServiceEdgeRouterPolicyOK
*/
type DetailServiceEdgeRouterPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailServiceEdgePolicyEnvelope `json:"body,omitempty"`
}

// NewDetailServiceEdgeRouterPolicyOK creates DetailServiceEdgeRouterPolicyOK with default headers values
func NewDetailServiceEdgeRouterPolicyOK() *DetailServiceEdgeRouterPolicyOK {

	return &DetailServiceEdgeRouterPolicyOK{}
}

// WithPayload adds the payload to the detail service edge router policy o k response
func (o *DetailServiceEdgeRouterPolicyOK) WithPayload(payload *rest_model.DetailServiceEdgePolicyEnvelope) *DetailServiceEdgeRouterPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail service edge router policy o k response
func (o *DetailServiceEdgeRouterPolicyOK) SetPayload(payload *rest_model.DetailServiceEdgePolicyEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailServiceEdgeRouterPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailServiceEdgeRouterPolicyUnauthorizedCode is the HTTP code returned for type DetailServiceEdgeRouterPolicyUnauthorized
const DetailServiceEdgeRouterPolicyUnauthorizedCode int = 401

/*DetailServiceEdgeRouterPolicyUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailServiceEdgeRouterPolicyUnauthorized
*/
type DetailServiceEdgeRouterPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailServiceEdgeRouterPolicyUnauthorized creates DetailServiceEdgeRouterPolicyUnauthorized with default headers values
func NewDetailServiceEdgeRouterPolicyUnauthorized() *DetailServiceEdgeRouterPolicyUnauthorized {

	return &DetailServiceEdgeRouterPolicyUnauthorized{}
}

// WithPayload adds the payload to the detail service edge router policy unauthorized response
func (o *DetailServiceEdgeRouterPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailServiceEdgeRouterPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail service edge router policy unauthorized response
func (o *DetailServiceEdgeRouterPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailServiceEdgeRouterPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailServiceEdgeRouterPolicyNotFoundCode is the HTTP code returned for type DetailServiceEdgeRouterPolicyNotFound
const DetailServiceEdgeRouterPolicyNotFoundCode int = 404

/*DetailServiceEdgeRouterPolicyNotFound The requested resource does not exist

swagger:response detailServiceEdgeRouterPolicyNotFound
*/
type DetailServiceEdgeRouterPolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailServiceEdgeRouterPolicyNotFound creates DetailServiceEdgeRouterPolicyNotFound with default headers values
func NewDetailServiceEdgeRouterPolicyNotFound() *DetailServiceEdgeRouterPolicyNotFound {

	return &DetailServiceEdgeRouterPolicyNotFound{}
}

// WithPayload adds the payload to the detail service edge router policy not found response
func (o *DetailServiceEdgeRouterPolicyNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailServiceEdgeRouterPolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail service edge router policy not found response
func (o *DetailServiceEdgeRouterPolicyNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailServiceEdgeRouterPolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailServiceEdgeRouterPolicyTooManyRequestsCode is the HTTP code returned for type DetailServiceEdgeRouterPolicyTooManyRequests
const DetailServiceEdgeRouterPolicyTooManyRequestsCode int = 429

/*DetailServiceEdgeRouterPolicyTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailServiceEdgeRouterPolicyTooManyRequests
*/
type DetailServiceEdgeRouterPolicyTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailServiceEdgeRouterPolicyTooManyRequests creates DetailServiceEdgeRouterPolicyTooManyRequests with default headers values
func NewDetailServiceEdgeRouterPolicyTooManyRequests() *DetailServiceEdgeRouterPolicyTooManyRequests {

	return &DetailServiceEdgeRouterPolicyTooManyRequests{}
}

// WithPayload adds the payload to the detail service edge router policy too many requests response
func (o *DetailServiceEdgeRouterPolicyTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailServiceEdgeRouterPolicyTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail service edge router policy too many requests response
func (o *DetailServiceEdgeRouterPolicyTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailServiceEdgeRouterPolicyTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailServiceEdgeRouterPolicyServiceUnavailableCode is the HTTP code returned for type DetailServiceEdgeRouterPolicyServiceUnavailable
const DetailServiceEdgeRouterPolicyServiceUnavailableCode int = 503

/*DetailServiceEdgeRouterPolicyServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response detailServiceEdgeRouterPolicyServiceUnavailable
*/
type DetailServiceEdgeRouterPolicyServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailServiceEdgeRouterPolicyServiceUnavailable creates DetailServiceEdgeRouterPolicyServiceUnavailable with default headers values
func NewDetailServiceEdgeRouterPolicyServiceUnavailable() *DetailServiceEdgeRouterPolicyServiceUnavailable {

	return &DetailServiceEdgeRouterPolicyServiceUnavailable{}
}

// WithPayload adds the payload to the detail service edge router policy service unavailable response
func (o *DetailServiceEdgeRouterPolicyServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailServiceEdgeRouterPolicyServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail service edge router policy service unavailable response
func (o *DetailServiceEdgeRouterPolicyServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailServiceEdgeRouterPolicyServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

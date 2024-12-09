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

// UpdateServiceEdgeRouterPolicyOKCode is the HTTP code returned for type UpdateServiceEdgeRouterPolicyOK
const UpdateServiceEdgeRouterPolicyOKCode int = 200

/*UpdateServiceEdgeRouterPolicyOK The update request was successful and the resource has been altered

swagger:response updateServiceEdgeRouterPolicyOK
*/
type UpdateServiceEdgeRouterPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdateServiceEdgeRouterPolicyOK creates UpdateServiceEdgeRouterPolicyOK with default headers values
func NewUpdateServiceEdgeRouterPolicyOK() *UpdateServiceEdgeRouterPolicyOK {

	return &UpdateServiceEdgeRouterPolicyOK{}
}

// WithPayload adds the payload to the update service edge router policy o k response
func (o *UpdateServiceEdgeRouterPolicyOK) WithPayload(payload *rest_model.Empty) *UpdateServiceEdgeRouterPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service edge router policy o k response
func (o *UpdateServiceEdgeRouterPolicyOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceEdgeRouterPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateServiceEdgeRouterPolicyBadRequestCode is the HTTP code returned for type UpdateServiceEdgeRouterPolicyBadRequest
const UpdateServiceEdgeRouterPolicyBadRequestCode int = 400

/*UpdateServiceEdgeRouterPolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updateServiceEdgeRouterPolicyBadRequest
*/
type UpdateServiceEdgeRouterPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateServiceEdgeRouterPolicyBadRequest creates UpdateServiceEdgeRouterPolicyBadRequest with default headers values
func NewUpdateServiceEdgeRouterPolicyBadRequest() *UpdateServiceEdgeRouterPolicyBadRequest {

	return &UpdateServiceEdgeRouterPolicyBadRequest{}
}

// WithPayload adds the payload to the update service edge router policy bad request response
func (o *UpdateServiceEdgeRouterPolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateServiceEdgeRouterPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service edge router policy bad request response
func (o *UpdateServiceEdgeRouterPolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceEdgeRouterPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateServiceEdgeRouterPolicyUnauthorizedCode is the HTTP code returned for type UpdateServiceEdgeRouterPolicyUnauthorized
const UpdateServiceEdgeRouterPolicyUnauthorizedCode int = 401

/*UpdateServiceEdgeRouterPolicyUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response updateServiceEdgeRouterPolicyUnauthorized
*/
type UpdateServiceEdgeRouterPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateServiceEdgeRouterPolicyUnauthorized creates UpdateServiceEdgeRouterPolicyUnauthorized with default headers values
func NewUpdateServiceEdgeRouterPolicyUnauthorized() *UpdateServiceEdgeRouterPolicyUnauthorized {

	return &UpdateServiceEdgeRouterPolicyUnauthorized{}
}

// WithPayload adds the payload to the update service edge router policy unauthorized response
func (o *UpdateServiceEdgeRouterPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateServiceEdgeRouterPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service edge router policy unauthorized response
func (o *UpdateServiceEdgeRouterPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceEdgeRouterPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateServiceEdgeRouterPolicyNotFoundCode is the HTTP code returned for type UpdateServiceEdgeRouterPolicyNotFound
const UpdateServiceEdgeRouterPolicyNotFoundCode int = 404

/*UpdateServiceEdgeRouterPolicyNotFound The requested resource does not exist

swagger:response updateServiceEdgeRouterPolicyNotFound
*/
type UpdateServiceEdgeRouterPolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateServiceEdgeRouterPolicyNotFound creates UpdateServiceEdgeRouterPolicyNotFound with default headers values
func NewUpdateServiceEdgeRouterPolicyNotFound() *UpdateServiceEdgeRouterPolicyNotFound {

	return &UpdateServiceEdgeRouterPolicyNotFound{}
}

// WithPayload adds the payload to the update service edge router policy not found response
func (o *UpdateServiceEdgeRouterPolicyNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateServiceEdgeRouterPolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service edge router policy not found response
func (o *UpdateServiceEdgeRouterPolicyNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceEdgeRouterPolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateServiceEdgeRouterPolicyTooManyRequestsCode is the HTTP code returned for type UpdateServiceEdgeRouterPolicyTooManyRequests
const UpdateServiceEdgeRouterPolicyTooManyRequestsCode int = 429

/*UpdateServiceEdgeRouterPolicyTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response updateServiceEdgeRouterPolicyTooManyRequests
*/
type UpdateServiceEdgeRouterPolicyTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateServiceEdgeRouterPolicyTooManyRequests creates UpdateServiceEdgeRouterPolicyTooManyRequests with default headers values
func NewUpdateServiceEdgeRouterPolicyTooManyRequests() *UpdateServiceEdgeRouterPolicyTooManyRequests {

	return &UpdateServiceEdgeRouterPolicyTooManyRequests{}
}

// WithPayload adds the payload to the update service edge router policy too many requests response
func (o *UpdateServiceEdgeRouterPolicyTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateServiceEdgeRouterPolicyTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service edge router policy too many requests response
func (o *UpdateServiceEdgeRouterPolicyTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceEdgeRouterPolicyTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateServiceEdgeRouterPolicyServiceUnavailableCode is the HTTP code returned for type UpdateServiceEdgeRouterPolicyServiceUnavailable
const UpdateServiceEdgeRouterPolicyServiceUnavailableCode int = 503

/*UpdateServiceEdgeRouterPolicyServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response updateServiceEdgeRouterPolicyServiceUnavailable
*/
type UpdateServiceEdgeRouterPolicyServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateServiceEdgeRouterPolicyServiceUnavailable creates UpdateServiceEdgeRouterPolicyServiceUnavailable with default headers values
func NewUpdateServiceEdgeRouterPolicyServiceUnavailable() *UpdateServiceEdgeRouterPolicyServiceUnavailable {

	return &UpdateServiceEdgeRouterPolicyServiceUnavailable{}
}

// WithPayload adds the payload to the update service edge router policy service unavailable response
func (o *UpdateServiceEdgeRouterPolicyServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateServiceEdgeRouterPolicyServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service edge router policy service unavailable response
func (o *UpdateServiceEdgeRouterPolicyServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceEdgeRouterPolicyServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

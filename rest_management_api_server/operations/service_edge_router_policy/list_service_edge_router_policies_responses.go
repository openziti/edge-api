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

// ListServiceEdgeRouterPoliciesOKCode is the HTTP code returned for type ListServiceEdgeRouterPoliciesOK
const ListServiceEdgeRouterPoliciesOKCode int = 200

/*ListServiceEdgeRouterPoliciesOK A list of service edge router policies

swagger:response listServiceEdgeRouterPoliciesOK
*/
type ListServiceEdgeRouterPoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPoliciesOK creates ListServiceEdgeRouterPoliciesOK with default headers values
func NewListServiceEdgeRouterPoliciesOK() *ListServiceEdgeRouterPoliciesOK {

	return &ListServiceEdgeRouterPoliciesOK{}
}

// WithPayload adds the payload to the list service edge router policies o k response
func (o *ListServiceEdgeRouterPoliciesOK) WithPayload(payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope) *ListServiceEdgeRouterPoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policies o k response
func (o *ListServiceEdgeRouterPoliciesOK) SetPayload(payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPoliciesBadRequestCode is the HTTP code returned for type ListServiceEdgeRouterPoliciesBadRequest
const ListServiceEdgeRouterPoliciesBadRequestCode int = 400

/*ListServiceEdgeRouterPoliciesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listServiceEdgeRouterPoliciesBadRequest
*/
type ListServiceEdgeRouterPoliciesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPoliciesBadRequest creates ListServiceEdgeRouterPoliciesBadRequest with default headers values
func NewListServiceEdgeRouterPoliciesBadRequest() *ListServiceEdgeRouterPoliciesBadRequest {

	return &ListServiceEdgeRouterPoliciesBadRequest{}
}

// WithPayload adds the payload to the list service edge router policies bad request response
func (o *ListServiceEdgeRouterPoliciesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPoliciesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policies bad request response
func (o *ListServiceEdgeRouterPoliciesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPoliciesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPoliciesUnauthorizedCode is the HTTP code returned for type ListServiceEdgeRouterPoliciesUnauthorized
const ListServiceEdgeRouterPoliciesUnauthorizedCode int = 401

/*ListServiceEdgeRouterPoliciesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServiceEdgeRouterPoliciesUnauthorized
*/
type ListServiceEdgeRouterPoliciesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPoliciesUnauthorized creates ListServiceEdgeRouterPoliciesUnauthorized with default headers values
func NewListServiceEdgeRouterPoliciesUnauthorized() *ListServiceEdgeRouterPoliciesUnauthorized {

	return &ListServiceEdgeRouterPoliciesUnauthorized{}
}

// WithPayload adds the payload to the list service edge router policies unauthorized response
func (o *ListServiceEdgeRouterPoliciesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPoliciesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policies unauthorized response
func (o *ListServiceEdgeRouterPoliciesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPoliciesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRouterPoliciesTooManyRequestsCode is the HTTP code returned for type ListServiceEdgeRouterPoliciesTooManyRequests
const ListServiceEdgeRouterPoliciesTooManyRequestsCode int = 429

/*ListServiceEdgeRouterPoliciesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listServiceEdgeRouterPoliciesTooManyRequests
*/
type ListServiceEdgeRouterPoliciesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRouterPoliciesTooManyRequests creates ListServiceEdgeRouterPoliciesTooManyRequests with default headers values
func NewListServiceEdgeRouterPoliciesTooManyRequests() *ListServiceEdgeRouterPoliciesTooManyRequests {

	return &ListServiceEdgeRouterPoliciesTooManyRequests{}
}

// WithPayload adds the payload to the list service edge router policies too many requests response
func (o *ListServiceEdgeRouterPoliciesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRouterPoliciesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge router policies too many requests response
func (o *ListServiceEdgeRouterPoliciesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRouterPoliciesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

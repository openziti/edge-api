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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListServiceEdgeRoutersOKCode is the HTTP code returned for type ListServiceEdgeRoutersOK
const ListServiceEdgeRoutersOKCode int = 200

/*ListServiceEdgeRoutersOK A list of edge routers

swagger:response listServiceEdgeRoutersOK
*/
type ListServiceEdgeRoutersOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListEdgeRoutersEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRoutersOK creates ListServiceEdgeRoutersOK with default headers values
func NewListServiceEdgeRoutersOK() *ListServiceEdgeRoutersOK {

	return &ListServiceEdgeRoutersOK{}
}

// WithPayload adds the payload to the list service edge routers o k response
func (o *ListServiceEdgeRoutersOK) WithPayload(payload *rest_model.ListEdgeRoutersEnvelope) *ListServiceEdgeRoutersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge routers o k response
func (o *ListServiceEdgeRoutersOK) SetPayload(payload *rest_model.ListEdgeRoutersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRoutersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRoutersBadRequestCode is the HTTP code returned for type ListServiceEdgeRoutersBadRequest
const ListServiceEdgeRoutersBadRequestCode int = 400

/*ListServiceEdgeRoutersBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listServiceEdgeRoutersBadRequest
*/
type ListServiceEdgeRoutersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRoutersBadRequest creates ListServiceEdgeRoutersBadRequest with default headers values
func NewListServiceEdgeRoutersBadRequest() *ListServiceEdgeRoutersBadRequest {

	return &ListServiceEdgeRoutersBadRequest{}
}

// WithPayload adds the payload to the list service edge routers bad request response
func (o *ListServiceEdgeRoutersBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRoutersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge routers bad request response
func (o *ListServiceEdgeRoutersBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRoutersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRoutersUnauthorizedCode is the HTTP code returned for type ListServiceEdgeRoutersUnauthorized
const ListServiceEdgeRoutersUnauthorizedCode int = 401

/*ListServiceEdgeRoutersUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listServiceEdgeRoutersUnauthorized
*/
type ListServiceEdgeRoutersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRoutersUnauthorized creates ListServiceEdgeRoutersUnauthorized with default headers values
func NewListServiceEdgeRoutersUnauthorized() *ListServiceEdgeRoutersUnauthorized {

	return &ListServiceEdgeRoutersUnauthorized{}
}

// WithPayload adds the payload to the list service edge routers unauthorized response
func (o *ListServiceEdgeRoutersUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRoutersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge routers unauthorized response
func (o *ListServiceEdgeRoutersUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRoutersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceEdgeRoutersTooManyRequestsCode is the HTTP code returned for type ListServiceEdgeRoutersTooManyRequests
const ListServiceEdgeRoutersTooManyRequestsCode int = 429

/*ListServiceEdgeRoutersTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listServiceEdgeRoutersTooManyRequests
*/
type ListServiceEdgeRoutersTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceEdgeRoutersTooManyRequests creates ListServiceEdgeRoutersTooManyRequests with default headers values
func NewListServiceEdgeRoutersTooManyRequests() *ListServiceEdgeRoutersTooManyRequests {

	return &ListServiceEdgeRoutersTooManyRequests{}
}

// WithPayload adds the payload to the list service edge routers too many requests response
func (o *ListServiceEdgeRoutersTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceEdgeRoutersTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service edge routers too many requests response
func (o *ListServiceEdgeRoutersTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceEdgeRoutersTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

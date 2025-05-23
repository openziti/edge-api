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

// ListServicesOKCode is the HTTP code returned for type ListServicesOK
const ListServicesOKCode int = 200

/*ListServicesOK A list of services

swagger:response listServicesOK
*/
type ListServicesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServicesEnvelope `json:"body,omitempty"`
}

// NewListServicesOK creates ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {

	return &ListServicesOK{}
}

// WithPayload adds the payload to the list services o k response
func (o *ListServicesOK) WithPayload(payload *rest_model.ListServicesEnvelope) *ListServicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list services o k response
func (o *ListServicesOK) SetPayload(payload *rest_model.ListServicesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicesBadRequestCode is the HTTP code returned for type ListServicesBadRequest
const ListServicesBadRequestCode int = 400

/*ListServicesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listServicesBadRequest
*/
type ListServicesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicesBadRequest creates ListServicesBadRequest with default headers values
func NewListServicesBadRequest() *ListServicesBadRequest {

	return &ListServicesBadRequest{}
}

// WithPayload adds the payload to the list services bad request response
func (o *ListServicesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list services bad request response
func (o *ListServicesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicesUnauthorizedCode is the HTTP code returned for type ListServicesUnauthorized
const ListServicesUnauthorizedCode int = 401

/*ListServicesUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listServicesUnauthorized
*/
type ListServicesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicesUnauthorized creates ListServicesUnauthorized with default headers values
func NewListServicesUnauthorized() *ListServicesUnauthorized {

	return &ListServicesUnauthorized{}
}

// WithPayload adds the payload to the list services unauthorized response
func (o *ListServicesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list services unauthorized response
func (o *ListServicesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicesTooManyRequestsCode is the HTTP code returned for type ListServicesTooManyRequests
const ListServicesTooManyRequestsCode int = 429

/*ListServicesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listServicesTooManyRequests
*/
type ListServicesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicesTooManyRequests creates ListServicesTooManyRequests with default headers values
func NewListServicesTooManyRequests() *ListServicesTooManyRequests {

	return &ListServicesTooManyRequests{}
}

// WithPayload adds the payload to the list services too many requests response
func (o *ListServicesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list services too many requests response
func (o *ListServicesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicesServiceUnavailableCode is the HTTP code returned for type ListServicesServiceUnavailable
const ListServicesServiceUnavailableCode int = 503

/*ListServicesServiceUnavailable The request could not be completed due to the server being busy or in a temporarily bad state

swagger:response listServicesServiceUnavailable
*/
type ListServicesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicesServiceUnavailable creates ListServicesServiceUnavailable with default headers values
func NewListServicesServiceUnavailable() *ListServicesServiceUnavailable {

	return &ListServicesServiceUnavailable{}
}

// WithPayload adds the payload to the list services service unavailable response
func (o *ListServicesServiceUnavailable) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list services service unavailable response
func (o *ListServicesServiceUnavailable) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

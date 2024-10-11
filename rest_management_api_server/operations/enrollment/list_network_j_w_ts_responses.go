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

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListNetworkJWTsOKCode is the HTTP code returned for type ListNetworkJWTsOK
const ListNetworkJWTsOKCode int = 200

/*ListNetworkJWTsOK A list of network JWTs

swagger:response listNetworkJWTsOK
*/
type ListNetworkJWTsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListNetworkJWTsEnvelope `json:"body,omitempty"`
}

// NewListNetworkJWTsOK creates ListNetworkJWTsOK with default headers values
func NewListNetworkJWTsOK() *ListNetworkJWTsOK {

	return &ListNetworkJWTsOK{}
}

// WithPayload adds the payload to the list network j w ts o k response
func (o *ListNetworkJWTsOK) WithPayload(payload *rest_model.ListNetworkJWTsEnvelope) *ListNetworkJWTsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list network j w ts o k response
func (o *ListNetworkJWTsOK) SetPayload(payload *rest_model.ListNetworkJWTsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNetworkJWTsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListNetworkJWTsBadRequestCode is the HTTP code returned for type ListNetworkJWTsBadRequest
const ListNetworkJWTsBadRequestCode int = 400

/*ListNetworkJWTsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listNetworkJWTsBadRequest
*/
type ListNetworkJWTsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListNetworkJWTsBadRequest creates ListNetworkJWTsBadRequest with default headers values
func NewListNetworkJWTsBadRequest() *ListNetworkJWTsBadRequest {

	return &ListNetworkJWTsBadRequest{}
}

// WithPayload adds the payload to the list network j w ts bad request response
func (o *ListNetworkJWTsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListNetworkJWTsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list network j w ts bad request response
func (o *ListNetworkJWTsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNetworkJWTsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListNetworkJWTsTooManyRequestsCode is the HTTP code returned for type ListNetworkJWTsTooManyRequests
const ListNetworkJWTsTooManyRequestsCode int = 429

/*ListNetworkJWTsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listNetworkJWTsTooManyRequests
*/
type ListNetworkJWTsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListNetworkJWTsTooManyRequests creates ListNetworkJWTsTooManyRequests with default headers values
func NewListNetworkJWTsTooManyRequests() *ListNetworkJWTsTooManyRequests {

	return &ListNetworkJWTsTooManyRequests{}
}

// WithPayload adds the payload to the list network j w ts too many requests response
func (o *ListNetworkJWTsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListNetworkJWTsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list network j w ts too many requests response
func (o *ListNetworkJWTsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNetworkJWTsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

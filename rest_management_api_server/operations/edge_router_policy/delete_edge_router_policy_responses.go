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

package edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteEdgeRouterPolicyOKCode is the HTTP code returned for type DeleteEdgeRouterPolicyOK
const DeleteEdgeRouterPolicyOKCode int = 200

/*DeleteEdgeRouterPolicyOK The delete request was successful and the resource has been removed

swagger:response deleteEdgeRouterPolicyOK
*/
type DeleteEdgeRouterPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteEdgeRouterPolicyOK creates DeleteEdgeRouterPolicyOK with default headers values
func NewDeleteEdgeRouterPolicyOK() *DeleteEdgeRouterPolicyOK {

	return &DeleteEdgeRouterPolicyOK{}
}

// WithPayload adds the payload to the delete edge router policy o k response
func (o *DeleteEdgeRouterPolicyOK) WithPayload(payload *rest_model.Empty) *DeleteEdgeRouterPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete edge router policy o k response
func (o *DeleteEdgeRouterPolicyOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEdgeRouterPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEdgeRouterPolicyBadRequestCode is the HTTP code returned for type DeleteEdgeRouterPolicyBadRequest
const DeleteEdgeRouterPolicyBadRequestCode int = 400

/*DeleteEdgeRouterPolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteEdgeRouterPolicyBadRequest
*/
type DeleteEdgeRouterPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEdgeRouterPolicyBadRequest creates DeleteEdgeRouterPolicyBadRequest with default headers values
func NewDeleteEdgeRouterPolicyBadRequest() *DeleteEdgeRouterPolicyBadRequest {

	return &DeleteEdgeRouterPolicyBadRequest{}
}

// WithPayload adds the payload to the delete edge router policy bad request response
func (o *DeleteEdgeRouterPolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEdgeRouterPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete edge router policy bad request response
func (o *DeleteEdgeRouterPolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEdgeRouterPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEdgeRouterPolicyUnauthorizedCode is the HTTP code returned for type DeleteEdgeRouterPolicyUnauthorized
const DeleteEdgeRouterPolicyUnauthorizedCode int = 401

/*DeleteEdgeRouterPolicyUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response deleteEdgeRouterPolicyUnauthorized
*/
type DeleteEdgeRouterPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEdgeRouterPolicyUnauthorized creates DeleteEdgeRouterPolicyUnauthorized with default headers values
func NewDeleteEdgeRouterPolicyUnauthorized() *DeleteEdgeRouterPolicyUnauthorized {

	return &DeleteEdgeRouterPolicyUnauthorized{}
}

// WithPayload adds the payload to the delete edge router policy unauthorized response
func (o *DeleteEdgeRouterPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEdgeRouterPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete edge router policy unauthorized response
func (o *DeleteEdgeRouterPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEdgeRouterPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEdgeRouterPolicyConflictCode is the HTTP code returned for type DeleteEdgeRouterPolicyConflict
const DeleteEdgeRouterPolicyConflictCode int = 409

/*DeleteEdgeRouterPolicyConflict The resource requested to be removed/altered cannot be as it is referenced by another object.

swagger:response deleteEdgeRouterPolicyConflict
*/
type DeleteEdgeRouterPolicyConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEdgeRouterPolicyConflict creates DeleteEdgeRouterPolicyConflict with default headers values
func NewDeleteEdgeRouterPolicyConflict() *DeleteEdgeRouterPolicyConflict {

	return &DeleteEdgeRouterPolicyConflict{}
}

// WithPayload adds the payload to the delete edge router policy conflict response
func (o *DeleteEdgeRouterPolicyConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEdgeRouterPolicyConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete edge router policy conflict response
func (o *DeleteEdgeRouterPolicyConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEdgeRouterPolicyConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEdgeRouterPolicyTooManyRequestsCode is the HTTP code returned for type DeleteEdgeRouterPolicyTooManyRequests
const DeleteEdgeRouterPolicyTooManyRequestsCode int = 429

/*DeleteEdgeRouterPolicyTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteEdgeRouterPolicyTooManyRequests
*/
type DeleteEdgeRouterPolicyTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteEdgeRouterPolicyTooManyRequests creates DeleteEdgeRouterPolicyTooManyRequests with default headers values
func NewDeleteEdgeRouterPolicyTooManyRequests() *DeleteEdgeRouterPolicyTooManyRequests {

	return &DeleteEdgeRouterPolicyTooManyRequests{}
}

// WithPayload adds the payload to the delete edge router policy too many requests response
func (o *DeleteEdgeRouterPolicyTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteEdgeRouterPolicyTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete edge router policy too many requests response
func (o *DeleteEdgeRouterPolicyTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEdgeRouterPolicyTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

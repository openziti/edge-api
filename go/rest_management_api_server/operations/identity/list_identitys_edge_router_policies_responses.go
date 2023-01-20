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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/go/rest_model"
)

// ListIdentitysEdgeRouterPoliciesOKCode is the HTTP code returned for type ListIdentitysEdgeRouterPoliciesOK
const ListIdentitysEdgeRouterPoliciesOKCode int = 200

/*ListIdentitysEdgeRouterPoliciesOK A list of edge router policies

swagger:response listIdentitysEdgeRouterPoliciesOK
*/
type ListIdentitysEdgeRouterPoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListEdgeRouterPoliciesEnvelope `json:"body,omitempty"`
}

// NewListIdentitysEdgeRouterPoliciesOK creates ListIdentitysEdgeRouterPoliciesOK with default headers values
func NewListIdentitysEdgeRouterPoliciesOK() *ListIdentitysEdgeRouterPoliciesOK {

	return &ListIdentitysEdgeRouterPoliciesOK{}
}

// WithPayload adds the payload to the list identitys edge router policies o k response
func (o *ListIdentitysEdgeRouterPoliciesOK) WithPayload(payload *rest_model.ListEdgeRouterPoliciesEnvelope) *ListIdentitysEdgeRouterPoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identitys edge router policies o k response
func (o *ListIdentitysEdgeRouterPoliciesOK) SetPayload(payload *rest_model.ListEdgeRouterPoliciesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitysEdgeRouterPoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitysEdgeRouterPoliciesUnauthorizedCode is the HTTP code returned for type ListIdentitysEdgeRouterPoliciesUnauthorized
const ListIdentitysEdgeRouterPoliciesUnauthorizedCode int = 401

/*ListIdentitysEdgeRouterPoliciesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listIdentitysEdgeRouterPoliciesUnauthorized
*/
type ListIdentitysEdgeRouterPoliciesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitysEdgeRouterPoliciesUnauthorized creates ListIdentitysEdgeRouterPoliciesUnauthorized with default headers values
func NewListIdentitysEdgeRouterPoliciesUnauthorized() *ListIdentitysEdgeRouterPoliciesUnauthorized {

	return &ListIdentitysEdgeRouterPoliciesUnauthorized{}
}

// WithPayload adds the payload to the list identitys edge router policies unauthorized response
func (o *ListIdentitysEdgeRouterPoliciesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitysEdgeRouterPoliciesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identitys edge router policies unauthorized response
func (o *ListIdentitysEdgeRouterPoliciesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitysEdgeRouterPoliciesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitysEdgeRouterPoliciesNotFoundCode is the HTTP code returned for type ListIdentitysEdgeRouterPoliciesNotFound
const ListIdentitysEdgeRouterPoliciesNotFoundCode int = 404

/*ListIdentitysEdgeRouterPoliciesNotFound The requested resource does not exist

swagger:response listIdentitysEdgeRouterPoliciesNotFound
*/
type ListIdentitysEdgeRouterPoliciesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitysEdgeRouterPoliciesNotFound creates ListIdentitysEdgeRouterPoliciesNotFound with default headers values
func NewListIdentitysEdgeRouterPoliciesNotFound() *ListIdentitysEdgeRouterPoliciesNotFound {

	return &ListIdentitysEdgeRouterPoliciesNotFound{}
}

// WithPayload adds the payload to the list identitys edge router policies not found response
func (o *ListIdentitysEdgeRouterPoliciesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitysEdgeRouterPoliciesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identitys edge router policies not found response
func (o *ListIdentitysEdgeRouterPoliciesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitysEdgeRouterPoliciesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

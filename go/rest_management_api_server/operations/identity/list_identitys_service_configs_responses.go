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

// ListIdentitysServiceConfigsOKCode is the HTTP code returned for type ListIdentitysServiceConfigsOK
const ListIdentitysServiceConfigsOKCode int = 200

/*ListIdentitysServiceConfigsOK A list of service configs

swagger:response listIdentitysServiceConfigsOK
*/
type ListIdentitysServiceConfigsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServiceConfigsEnvelope `json:"body,omitempty"`
}

// NewListIdentitysServiceConfigsOK creates ListIdentitysServiceConfigsOK with default headers values
func NewListIdentitysServiceConfigsOK() *ListIdentitysServiceConfigsOK {

	return &ListIdentitysServiceConfigsOK{}
}

// WithPayload adds the payload to the list identitys service configs o k response
func (o *ListIdentitysServiceConfigsOK) WithPayload(payload *rest_model.ListServiceConfigsEnvelope) *ListIdentitysServiceConfigsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identitys service configs o k response
func (o *ListIdentitysServiceConfigsOK) SetPayload(payload *rest_model.ListServiceConfigsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitysServiceConfigsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitysServiceConfigsUnauthorizedCode is the HTTP code returned for type ListIdentitysServiceConfigsUnauthorized
const ListIdentitysServiceConfigsUnauthorizedCode int = 401

/*ListIdentitysServiceConfigsUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listIdentitysServiceConfigsUnauthorized
*/
type ListIdentitysServiceConfigsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitysServiceConfigsUnauthorized creates ListIdentitysServiceConfigsUnauthorized with default headers values
func NewListIdentitysServiceConfigsUnauthorized() *ListIdentitysServiceConfigsUnauthorized {

	return &ListIdentitysServiceConfigsUnauthorized{}
}

// WithPayload adds the payload to the list identitys service configs unauthorized response
func (o *ListIdentitysServiceConfigsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitysServiceConfigsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identitys service configs unauthorized response
func (o *ListIdentitysServiceConfigsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitysServiceConfigsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitysServiceConfigsNotFoundCode is the HTTP code returned for type ListIdentitysServiceConfigsNotFound
const ListIdentitysServiceConfigsNotFoundCode int = 404

/*ListIdentitysServiceConfigsNotFound The requested resource does not exist

swagger:response listIdentitysServiceConfigsNotFound
*/
type ListIdentitysServiceConfigsNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitysServiceConfigsNotFound creates ListIdentitysServiceConfigsNotFound with default headers values
func NewListIdentitysServiceConfigsNotFound() *ListIdentitysServiceConfigsNotFound {

	return &ListIdentitysServiceConfigsNotFound{}
}

// WithPayload adds the payload to the list identitys service configs not found response
func (o *ListIdentitysServiceConfigsNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitysServiceConfigsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identitys service configs not found response
func (o *ListIdentitysServiceConfigsNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitysServiceConfigsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

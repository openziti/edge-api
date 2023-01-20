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

// DisableIdentityOKCode is the HTTP code returned for type DisableIdentityOK
const DisableIdentityOKCode int = 200

/*DisableIdentityOK Base empty response

swagger:response disableIdentityOK
*/
type DisableIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDisableIdentityOK creates DisableIdentityOK with default headers values
func NewDisableIdentityOK() *DisableIdentityOK {

	return &DisableIdentityOK{}
}

// WithPayload adds the payload to the disable identity o k response
func (o *DisableIdentityOK) WithPayload(payload *rest_model.Empty) *DisableIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the disable identity o k response
func (o *DisableIdentityOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DisableIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DisableIdentityUnauthorizedCode is the HTTP code returned for type DisableIdentityUnauthorized
const DisableIdentityUnauthorizedCode int = 401

/*DisableIdentityUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response disableIdentityUnauthorized
*/
type DisableIdentityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDisableIdentityUnauthorized creates DisableIdentityUnauthorized with default headers values
func NewDisableIdentityUnauthorized() *DisableIdentityUnauthorized {

	return &DisableIdentityUnauthorized{}
}

// WithPayload adds the payload to the disable identity unauthorized response
func (o *DisableIdentityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DisableIdentityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the disable identity unauthorized response
func (o *DisableIdentityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DisableIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DisableIdentityNotFoundCode is the HTTP code returned for type DisableIdentityNotFound
const DisableIdentityNotFoundCode int = 404

/*DisableIdentityNotFound The requested resource does not exist

swagger:response disableIdentityNotFound
*/
type DisableIdentityNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDisableIdentityNotFound creates DisableIdentityNotFound with default headers values
func NewDisableIdentityNotFound() *DisableIdentityNotFound {

	return &DisableIdentityNotFound{}
}

// WithPayload adds the payload to the disable identity not found response
func (o *DisableIdentityNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DisableIdentityNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the disable identity not found response
func (o *DisableIdentityNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DisableIdentityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

// GetIdentityEnrollmentsOKCode is the HTTP code returned for type GetIdentityEnrollmentsOK
const GetIdentityEnrollmentsOKCode int = 200

/*GetIdentityEnrollmentsOK A list of enrollments

swagger:response getIdentityEnrollmentsOK
*/
type GetIdentityEnrollmentsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListEnrollmentsEnvelope `json:"body,omitempty"`
}

// NewGetIdentityEnrollmentsOK creates GetIdentityEnrollmentsOK with default headers values
func NewGetIdentityEnrollmentsOK() *GetIdentityEnrollmentsOK {

	return &GetIdentityEnrollmentsOK{}
}

// WithPayload adds the payload to the get identity enrollments o k response
func (o *GetIdentityEnrollmentsOK) WithPayload(payload *rest_model.ListEnrollmentsEnvelope) *GetIdentityEnrollmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity enrollments o k response
func (o *GetIdentityEnrollmentsOK) SetPayload(payload *rest_model.ListEnrollmentsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityEnrollmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIdentityEnrollmentsUnauthorizedCode is the HTTP code returned for type GetIdentityEnrollmentsUnauthorized
const GetIdentityEnrollmentsUnauthorizedCode int = 401

/*GetIdentityEnrollmentsUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response getIdentityEnrollmentsUnauthorized
*/
type GetIdentityEnrollmentsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetIdentityEnrollmentsUnauthorized creates GetIdentityEnrollmentsUnauthorized with default headers values
func NewGetIdentityEnrollmentsUnauthorized() *GetIdentityEnrollmentsUnauthorized {

	return &GetIdentityEnrollmentsUnauthorized{}
}

// WithPayload adds the payload to the get identity enrollments unauthorized response
func (o *GetIdentityEnrollmentsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *GetIdentityEnrollmentsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity enrollments unauthorized response
func (o *GetIdentityEnrollmentsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityEnrollmentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIdentityEnrollmentsNotFoundCode is the HTTP code returned for type GetIdentityEnrollmentsNotFound
const GetIdentityEnrollmentsNotFoundCode int = 404

/*GetIdentityEnrollmentsNotFound The requested resource does not exist

swagger:response getIdentityEnrollmentsNotFound
*/
type GetIdentityEnrollmentsNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetIdentityEnrollmentsNotFound creates GetIdentityEnrollmentsNotFound with default headers values
func NewGetIdentityEnrollmentsNotFound() *GetIdentityEnrollmentsNotFound {

	return &GetIdentityEnrollmentsNotFound{}
}

// WithPayload adds the payload to the get identity enrollments not found response
func (o *GetIdentityEnrollmentsNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *GetIdentityEnrollmentsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity enrollments not found response
func (o *GetIdentityEnrollmentsNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityEnrollmentsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

// DetailIdentityTypeOKCode is the HTTP code returned for type DetailIdentityTypeOK
const DetailIdentityTypeOKCode int = 200

/*DetailIdentityTypeOK A single identity type

swagger:response detailIdentityTypeOK
*/
type DetailIdentityTypeOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailIdentityTypeEnvelope `json:"body,omitempty"`
}

// NewDetailIdentityTypeOK creates DetailIdentityTypeOK with default headers values
func NewDetailIdentityTypeOK() *DetailIdentityTypeOK {

	return &DetailIdentityTypeOK{}
}

// WithPayload adds the payload to the detail identity type o k response
func (o *DetailIdentityTypeOK) WithPayload(payload *rest_model.DetailIdentityTypeEnvelope) *DetailIdentityTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail identity type o k response
func (o *DetailIdentityTypeOK) SetPayload(payload *rest_model.DetailIdentityTypeEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailIdentityTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailIdentityTypeUnauthorizedCode is the HTTP code returned for type DetailIdentityTypeUnauthorized
const DetailIdentityTypeUnauthorizedCode int = 401

/*DetailIdentityTypeUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailIdentityTypeUnauthorized
*/
type DetailIdentityTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailIdentityTypeUnauthorized creates DetailIdentityTypeUnauthorized with default headers values
func NewDetailIdentityTypeUnauthorized() *DetailIdentityTypeUnauthorized {

	return &DetailIdentityTypeUnauthorized{}
}

// WithPayload adds the payload to the detail identity type unauthorized response
func (o *DetailIdentityTypeUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailIdentityTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail identity type unauthorized response
func (o *DetailIdentityTypeUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailIdentityTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailIdentityTypeNotFoundCode is the HTTP code returned for type DetailIdentityTypeNotFound
const DetailIdentityTypeNotFoundCode int = 404

/*DetailIdentityTypeNotFound The requested resource does not exist

swagger:response detailIdentityTypeNotFound
*/
type DetailIdentityTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailIdentityTypeNotFound creates DetailIdentityTypeNotFound with default headers values
func NewDetailIdentityTypeNotFound() *DetailIdentityTypeNotFound {

	return &DetailIdentityTypeNotFound{}
}

// WithPayload adds the payload to the detail identity type not found response
func (o *DetailIdentityTypeNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailIdentityTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail identity type not found response
func (o *DetailIdentityTypeNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailIdentityTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

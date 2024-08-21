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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// UpdateCurrentIdentityAuthenticatorOKCode is the HTTP code returned for type UpdateCurrentIdentityAuthenticatorOK
const UpdateCurrentIdentityAuthenticatorOKCode int = 200

/*UpdateCurrentIdentityAuthenticatorOK The update request was successful and the resource has been altered

swagger:response updateCurrentIdentityAuthenticatorOK
*/
type UpdateCurrentIdentityAuthenticatorOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdateCurrentIdentityAuthenticatorOK creates UpdateCurrentIdentityAuthenticatorOK with default headers values
func NewUpdateCurrentIdentityAuthenticatorOK() *UpdateCurrentIdentityAuthenticatorOK {

	return &UpdateCurrentIdentityAuthenticatorOK{}
}

// WithPayload adds the payload to the update current identity authenticator o k response
func (o *UpdateCurrentIdentityAuthenticatorOK) WithPayload(payload *rest_model.Empty) *UpdateCurrentIdentityAuthenticatorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update current identity authenticator o k response
func (o *UpdateCurrentIdentityAuthenticatorOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCurrentIdentityAuthenticatorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCurrentIdentityAuthenticatorBadRequestCode is the HTTP code returned for type UpdateCurrentIdentityAuthenticatorBadRequest
const UpdateCurrentIdentityAuthenticatorBadRequestCode int = 400

/*UpdateCurrentIdentityAuthenticatorBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updateCurrentIdentityAuthenticatorBadRequest
*/
type UpdateCurrentIdentityAuthenticatorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateCurrentIdentityAuthenticatorBadRequest creates UpdateCurrentIdentityAuthenticatorBadRequest with default headers values
func NewUpdateCurrentIdentityAuthenticatorBadRequest() *UpdateCurrentIdentityAuthenticatorBadRequest {

	return &UpdateCurrentIdentityAuthenticatorBadRequest{}
}

// WithPayload adds the payload to the update current identity authenticator bad request response
func (o *UpdateCurrentIdentityAuthenticatorBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateCurrentIdentityAuthenticatorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update current identity authenticator bad request response
func (o *UpdateCurrentIdentityAuthenticatorBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCurrentIdentityAuthenticatorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCurrentIdentityAuthenticatorUnauthorizedCode is the HTTP code returned for type UpdateCurrentIdentityAuthenticatorUnauthorized
const UpdateCurrentIdentityAuthenticatorUnauthorizedCode int = 401

/*UpdateCurrentIdentityAuthenticatorUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response updateCurrentIdentityAuthenticatorUnauthorized
*/
type UpdateCurrentIdentityAuthenticatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateCurrentIdentityAuthenticatorUnauthorized creates UpdateCurrentIdentityAuthenticatorUnauthorized with default headers values
func NewUpdateCurrentIdentityAuthenticatorUnauthorized() *UpdateCurrentIdentityAuthenticatorUnauthorized {

	return &UpdateCurrentIdentityAuthenticatorUnauthorized{}
}

// WithPayload adds the payload to the update current identity authenticator unauthorized response
func (o *UpdateCurrentIdentityAuthenticatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateCurrentIdentityAuthenticatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update current identity authenticator unauthorized response
func (o *UpdateCurrentIdentityAuthenticatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCurrentIdentityAuthenticatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCurrentIdentityAuthenticatorNotFoundCode is the HTTP code returned for type UpdateCurrentIdentityAuthenticatorNotFound
const UpdateCurrentIdentityAuthenticatorNotFoundCode int = 404

/*UpdateCurrentIdentityAuthenticatorNotFound The requested resource does not exist

swagger:response updateCurrentIdentityAuthenticatorNotFound
*/
type UpdateCurrentIdentityAuthenticatorNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateCurrentIdentityAuthenticatorNotFound creates UpdateCurrentIdentityAuthenticatorNotFound with default headers values
func NewUpdateCurrentIdentityAuthenticatorNotFound() *UpdateCurrentIdentityAuthenticatorNotFound {

	return &UpdateCurrentIdentityAuthenticatorNotFound{}
}

// WithPayload adds the payload to the update current identity authenticator not found response
func (o *UpdateCurrentIdentityAuthenticatorNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateCurrentIdentityAuthenticatorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update current identity authenticator not found response
func (o *UpdateCurrentIdentityAuthenticatorNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCurrentIdentityAuthenticatorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// UpdateControllerSettingOKCode is the HTTP code returned for type UpdateControllerSettingOK
const UpdateControllerSettingOKCode int = 200

/*UpdateControllerSettingOK The update request was successful and the resource has been altered

swagger:response updateControllerSettingOK
*/
type UpdateControllerSettingOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewUpdateControllerSettingOK creates UpdateControllerSettingOK with default headers values
func NewUpdateControllerSettingOK() *UpdateControllerSettingOK {

	return &UpdateControllerSettingOK{}
}

// WithPayload adds the payload to the update controller setting o k response
func (o *UpdateControllerSettingOK) WithPayload(payload *rest_model.Empty) *UpdateControllerSettingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update controller setting o k response
func (o *UpdateControllerSettingOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateControllerSettingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateControllerSettingBadRequestCode is the HTTP code returned for type UpdateControllerSettingBadRequest
const UpdateControllerSettingBadRequestCode int = 400

/*UpdateControllerSettingBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response updateControllerSettingBadRequest
*/
type UpdateControllerSettingBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateControllerSettingBadRequest creates UpdateControllerSettingBadRequest with default headers values
func NewUpdateControllerSettingBadRequest() *UpdateControllerSettingBadRequest {

	return &UpdateControllerSettingBadRequest{}
}

// WithPayload adds the payload to the update controller setting bad request response
func (o *UpdateControllerSettingBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateControllerSettingBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update controller setting bad request response
func (o *UpdateControllerSettingBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateControllerSettingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateControllerSettingUnauthorizedCode is the HTTP code returned for type UpdateControllerSettingUnauthorized
const UpdateControllerSettingUnauthorizedCode int = 401

/*UpdateControllerSettingUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response updateControllerSettingUnauthorized
*/
type UpdateControllerSettingUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateControllerSettingUnauthorized creates UpdateControllerSettingUnauthorized with default headers values
func NewUpdateControllerSettingUnauthorized() *UpdateControllerSettingUnauthorized {

	return &UpdateControllerSettingUnauthorized{}
}

// WithPayload adds the payload to the update controller setting unauthorized response
func (o *UpdateControllerSettingUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateControllerSettingUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update controller setting unauthorized response
func (o *UpdateControllerSettingUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateControllerSettingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateControllerSettingNotFoundCode is the HTTP code returned for type UpdateControllerSettingNotFound
const UpdateControllerSettingNotFoundCode int = 404

/*UpdateControllerSettingNotFound The requested resource does not exist

swagger:response updateControllerSettingNotFound
*/
type UpdateControllerSettingNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateControllerSettingNotFound creates UpdateControllerSettingNotFound with default headers values
func NewUpdateControllerSettingNotFound() *UpdateControllerSettingNotFound {

	return &UpdateControllerSettingNotFound{}
}

// WithPayload adds the payload to the update controller setting not found response
func (o *UpdateControllerSettingNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateControllerSettingNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update controller setting not found response
func (o *UpdateControllerSettingNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateControllerSettingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateControllerSettingTooManyRequestsCode is the HTTP code returned for type UpdateControllerSettingTooManyRequests
const UpdateControllerSettingTooManyRequestsCode int = 429

/*UpdateControllerSettingTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response updateControllerSettingTooManyRequests
*/
type UpdateControllerSettingTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewUpdateControllerSettingTooManyRequests creates UpdateControllerSettingTooManyRequests with default headers values
func NewUpdateControllerSettingTooManyRequests() *UpdateControllerSettingTooManyRequests {

	return &UpdateControllerSettingTooManyRequests{}
}

// WithPayload adds the payload to the update controller setting too many requests response
func (o *UpdateControllerSettingTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *UpdateControllerSettingTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update controller setting too many requests response
func (o *UpdateControllerSettingTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateControllerSettingTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

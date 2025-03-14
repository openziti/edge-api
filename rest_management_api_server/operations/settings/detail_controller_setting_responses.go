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

// DetailControllerSettingOKCode is the HTTP code returned for type DetailControllerSettingOK
const DetailControllerSettingOKCode int = 200

/*DetailControllerSettingOK A singular controller setting object

swagger:response detailControllerSettingOK
*/
type DetailControllerSettingOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailControllerSettingEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingOK creates DetailControllerSettingOK with default headers values
func NewDetailControllerSettingOK() *DetailControllerSettingOK {

	return &DetailControllerSettingOK{}
}

// WithPayload adds the payload to the detail controller setting o k response
func (o *DetailControllerSettingOK) WithPayload(payload *rest_model.DetailControllerSettingEnvelope) *DetailControllerSettingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting o k response
func (o *DetailControllerSettingOK) SetPayload(payload *rest_model.DetailControllerSettingEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailControllerSettingUnauthorizedCode is the HTTP code returned for type DetailControllerSettingUnauthorized
const DetailControllerSettingUnauthorizedCode int = 401

/*DetailControllerSettingUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailControllerSettingUnauthorized
*/
type DetailControllerSettingUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingUnauthorized creates DetailControllerSettingUnauthorized with default headers values
func NewDetailControllerSettingUnauthorized() *DetailControllerSettingUnauthorized {

	return &DetailControllerSettingUnauthorized{}
}

// WithPayload adds the payload to the detail controller setting unauthorized response
func (o *DetailControllerSettingUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailControllerSettingUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting unauthorized response
func (o *DetailControllerSettingUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailControllerSettingNotFoundCode is the HTTP code returned for type DetailControllerSettingNotFound
const DetailControllerSettingNotFoundCode int = 404

/*DetailControllerSettingNotFound The requested resource does not exist

swagger:response detailControllerSettingNotFound
*/
type DetailControllerSettingNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingNotFound creates DetailControllerSettingNotFound with default headers values
func NewDetailControllerSettingNotFound() *DetailControllerSettingNotFound {

	return &DetailControllerSettingNotFound{}
}

// WithPayload adds the payload to the detail controller setting not found response
func (o *DetailControllerSettingNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailControllerSettingNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting not found response
func (o *DetailControllerSettingNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailControllerSettingTooManyRequestsCode is the HTTP code returned for type DetailControllerSettingTooManyRequests
const DetailControllerSettingTooManyRequestsCode int = 429

/*DetailControllerSettingTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailControllerSettingTooManyRequests
*/
type DetailControllerSettingTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingTooManyRequests creates DetailControllerSettingTooManyRequests with default headers values
func NewDetailControllerSettingTooManyRequests() *DetailControllerSettingTooManyRequests {

	return &DetailControllerSettingTooManyRequests{}
}

// WithPayload adds the payload to the detail controller setting too many requests response
func (o *DetailControllerSettingTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailControllerSettingTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting too many requests response
func (o *DetailControllerSettingTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

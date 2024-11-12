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

// DetailControllerSettingEffectiveOKCode is the HTTP code returned for type DetailControllerSettingEffectiveOK
const DetailControllerSettingEffectiveOKCode int = 200

/*DetailControllerSettingEffectiveOK A singular controller's effective setting object

swagger:response detailControllerSettingEffectiveOK
*/
type DetailControllerSettingEffectiveOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailControllerSettingEffectiveEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingEffectiveOK creates DetailControllerSettingEffectiveOK with default headers values
func NewDetailControllerSettingEffectiveOK() *DetailControllerSettingEffectiveOK {

	return &DetailControllerSettingEffectiveOK{}
}

// WithPayload adds the payload to the detail controller setting effective o k response
func (o *DetailControllerSettingEffectiveOK) WithPayload(payload *rest_model.DetailControllerSettingEffectiveEnvelope) *DetailControllerSettingEffectiveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting effective o k response
func (o *DetailControllerSettingEffectiveOK) SetPayload(payload *rest_model.DetailControllerSettingEffectiveEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingEffectiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailControllerSettingEffectiveUnauthorizedCode is the HTTP code returned for type DetailControllerSettingEffectiveUnauthorized
const DetailControllerSettingEffectiveUnauthorizedCode int = 401

/*DetailControllerSettingEffectiveUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailControllerSettingEffectiveUnauthorized
*/
type DetailControllerSettingEffectiveUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingEffectiveUnauthorized creates DetailControllerSettingEffectiveUnauthorized with default headers values
func NewDetailControllerSettingEffectiveUnauthorized() *DetailControllerSettingEffectiveUnauthorized {

	return &DetailControllerSettingEffectiveUnauthorized{}
}

// WithPayload adds the payload to the detail controller setting effective unauthorized response
func (o *DetailControllerSettingEffectiveUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailControllerSettingEffectiveUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting effective unauthorized response
func (o *DetailControllerSettingEffectiveUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingEffectiveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailControllerSettingEffectiveNotFoundCode is the HTTP code returned for type DetailControllerSettingEffectiveNotFound
const DetailControllerSettingEffectiveNotFoundCode int = 404

/*DetailControllerSettingEffectiveNotFound The requested resource does not exist

swagger:response detailControllerSettingEffectiveNotFound
*/
type DetailControllerSettingEffectiveNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingEffectiveNotFound creates DetailControllerSettingEffectiveNotFound with default headers values
func NewDetailControllerSettingEffectiveNotFound() *DetailControllerSettingEffectiveNotFound {

	return &DetailControllerSettingEffectiveNotFound{}
}

// WithPayload adds the payload to the detail controller setting effective not found response
func (o *DetailControllerSettingEffectiveNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailControllerSettingEffectiveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting effective not found response
func (o *DetailControllerSettingEffectiveNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingEffectiveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailControllerSettingEffectiveTooManyRequestsCode is the HTTP code returned for type DetailControllerSettingEffectiveTooManyRequests
const DetailControllerSettingEffectiveTooManyRequestsCode int = 429

/*DetailControllerSettingEffectiveTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailControllerSettingEffectiveTooManyRequests
*/
type DetailControllerSettingEffectiveTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailControllerSettingEffectiveTooManyRequests creates DetailControllerSettingEffectiveTooManyRequests with default headers values
func NewDetailControllerSettingEffectiveTooManyRequests() *DetailControllerSettingEffectiveTooManyRequests {

	return &DetailControllerSettingEffectiveTooManyRequests{}
}

// WithPayload adds the payload to the detail controller setting effective too many requests response
func (o *DetailControllerSettingEffectiveTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailControllerSettingEffectiveTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail controller setting effective too many requests response
func (o *DetailControllerSettingEffectiveTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailControllerSettingEffectiveTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

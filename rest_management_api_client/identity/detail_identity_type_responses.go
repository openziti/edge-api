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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DetailIdentityTypeReader is a Reader for the DetailIdentityType structure.
type DetailIdentityTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailIdentityTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailIdentityTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailIdentityTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailIdentityTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailIdentityTypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDetailIdentityTypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailIdentityTypeOK creates a DetailIdentityTypeOK with default headers values
func NewDetailIdentityTypeOK() *DetailIdentityTypeOK {
	return &DetailIdentityTypeOK{}
}

/* DetailIdentityTypeOK describes a response with status code 200, with default header values.

A single identity type
*/
type DetailIdentityTypeOK struct {
	Payload *rest_model.DetailIdentityTypeEnvelope
}

func (o *DetailIdentityTypeOK) Error() string {
	return fmt.Sprintf("[GET /identity-types/{id}][%d] detailIdentityTypeOK  %+v", 200, o.Payload)
}
func (o *DetailIdentityTypeOK) GetPayload() *rest_model.DetailIdentityTypeEnvelope {
	return o.Payload
}

func (o *DetailIdentityTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailIdentityTypeEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailIdentityTypeUnauthorized creates a DetailIdentityTypeUnauthorized with default headers values
func NewDetailIdentityTypeUnauthorized() *DetailIdentityTypeUnauthorized {
	return &DetailIdentityTypeUnauthorized{}
}

/* DetailIdentityTypeUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailIdentityTypeUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailIdentityTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identity-types/{id}][%d] detailIdentityTypeUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailIdentityTypeUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailIdentityTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailIdentityTypeNotFound creates a DetailIdentityTypeNotFound with default headers values
func NewDetailIdentityTypeNotFound() *DetailIdentityTypeNotFound {
	return &DetailIdentityTypeNotFound{}
}

/* DetailIdentityTypeNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailIdentityTypeNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailIdentityTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /identity-types/{id}][%d] detailIdentityTypeNotFound  %+v", 404, o.Payload)
}
func (o *DetailIdentityTypeNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailIdentityTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailIdentityTypeTooManyRequests creates a DetailIdentityTypeTooManyRequests with default headers values
func NewDetailIdentityTypeTooManyRequests() *DetailIdentityTypeTooManyRequests {
	return &DetailIdentityTypeTooManyRequests{}
}

/* DetailIdentityTypeTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailIdentityTypeTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailIdentityTypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /identity-types/{id}][%d] detailIdentityTypeTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailIdentityTypeTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailIdentityTypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailIdentityTypeServiceUnavailable creates a DetailIdentityTypeServiceUnavailable with default headers values
func NewDetailIdentityTypeServiceUnavailable() *DetailIdentityTypeServiceUnavailable {
	return &DetailIdentityTypeServiceUnavailable{}
}

/* DetailIdentityTypeServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type DetailIdentityTypeServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailIdentityTypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /identity-types/{id}][%d] detailIdentityTypeServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DetailIdentityTypeServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailIdentityTypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

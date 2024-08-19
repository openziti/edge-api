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

package auth_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DetailAuthPolicyReader is a Reader for the DetailAuthPolicy structure.
type DetailAuthPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailAuthPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailAuthPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailAuthPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailAuthPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailAuthPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailAuthPolicyOK creates a DetailAuthPolicyOK with default headers values
func NewDetailAuthPolicyOK() *DetailAuthPolicyOK {
	return &DetailAuthPolicyOK{}
}

/*
DetailAuthPolicyOK describes a response with status code 200, with default header values.

A singular Auth Policy resource
*/
type DetailAuthPolicyOK struct {
	Payload *rest_model.DetailAuthPolicyEnvelope
}

func (o *DetailAuthPolicyOK) Error() string {
	return fmt.Sprintf("[GET /auth-policies/{id}][%d] detailAuthPolicyOK  %+v", 200, o.Payload)
}
func (o *DetailAuthPolicyOK) GetPayload() *rest_model.DetailAuthPolicyEnvelope {
	return o.Payload
}

func (o *DetailAuthPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailAuthPolicyEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailAuthPolicyUnauthorized creates a DetailAuthPolicyUnauthorized with default headers values
func NewDetailAuthPolicyUnauthorized() *DetailAuthPolicyUnauthorized {
	return &DetailAuthPolicyUnauthorized{}
}

/*
DetailAuthPolicyUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailAuthPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailAuthPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth-policies/{id}][%d] detailAuthPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailAuthPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailAuthPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailAuthPolicyNotFound creates a DetailAuthPolicyNotFound with default headers values
func NewDetailAuthPolicyNotFound() *DetailAuthPolicyNotFound {
	return &DetailAuthPolicyNotFound{}
}

/*
DetailAuthPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailAuthPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailAuthPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /auth-policies/{id}][%d] detailAuthPolicyNotFound  %+v", 404, o.Payload)
}
func (o *DetailAuthPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailAuthPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailAuthPolicyTooManyRequests creates a DetailAuthPolicyTooManyRequests with default headers values
func NewDetailAuthPolicyTooManyRequests() *DetailAuthPolicyTooManyRequests {
	return &DetailAuthPolicyTooManyRequests{}
}

/*
DetailAuthPolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailAuthPolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailAuthPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /auth-policies/{id}][%d] detailAuthPolicyTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailAuthPolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailAuthPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

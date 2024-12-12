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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListServicePolicyPostureChecksReader is a Reader for the ListServicePolicyPostureChecks structure.
type ListServicePolicyPostureChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicePolicyPostureChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicePolicyPostureChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServicePolicyPostureChecksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListServicePolicyPostureChecksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServicePolicyPostureChecksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListServicePolicyPostureChecksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServicePolicyPostureChecksOK creates a ListServicePolicyPostureChecksOK with default headers values
func NewListServicePolicyPostureChecksOK() *ListServicePolicyPostureChecksOK {
	return &ListServicePolicyPostureChecksOK{}
}

/* ListServicePolicyPostureChecksOK describes a response with status code 200, with default header values.

A list of posture checks
*/
type ListServicePolicyPostureChecksOK struct {
	Payload *rest_model.ListPostureCheckEnvelope
}

func (o *ListServicePolicyPostureChecksOK) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/posture-checks][%d] listServicePolicyPostureChecksOK  %+v", 200, o.Payload)
}
func (o *ListServicePolicyPostureChecksOK) GetPayload() *rest_model.ListPostureCheckEnvelope {
	return o.Payload
}

func (o *ListServicePolicyPostureChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListPostureCheckEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyPostureChecksBadRequest creates a ListServicePolicyPostureChecksBadRequest with default headers values
func NewListServicePolicyPostureChecksBadRequest() *ListServicePolicyPostureChecksBadRequest {
	return &ListServicePolicyPostureChecksBadRequest{}
}

/* ListServicePolicyPostureChecksBadRequest describes a response with status code 400, with default header values.

The requested resource does not exist
*/
type ListServicePolicyPostureChecksBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyPostureChecksBadRequest) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/posture-checks][%d] listServicePolicyPostureChecksBadRequest  %+v", 400, o.Payload)
}
func (o *ListServicePolicyPostureChecksBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyPostureChecksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyPostureChecksUnauthorized creates a ListServicePolicyPostureChecksUnauthorized with default headers values
func NewListServicePolicyPostureChecksUnauthorized() *ListServicePolicyPostureChecksUnauthorized {
	return &ListServicePolicyPostureChecksUnauthorized{}
}

/* ListServicePolicyPostureChecksUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListServicePolicyPostureChecksUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyPostureChecksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/posture-checks][%d] listServicePolicyPostureChecksUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServicePolicyPostureChecksUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyPostureChecksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyPostureChecksTooManyRequests creates a ListServicePolicyPostureChecksTooManyRequests with default headers values
func NewListServicePolicyPostureChecksTooManyRequests() *ListServicePolicyPostureChecksTooManyRequests {
	return &ListServicePolicyPostureChecksTooManyRequests{}
}

/* ListServicePolicyPostureChecksTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListServicePolicyPostureChecksTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyPostureChecksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/posture-checks][%d] listServicePolicyPostureChecksTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServicePolicyPostureChecksTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyPostureChecksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyPostureChecksServiceUnavailable creates a ListServicePolicyPostureChecksServiceUnavailable with default headers values
func NewListServicePolicyPostureChecksServiceUnavailable() *ListServicePolicyPostureChecksServiceUnavailable {
	return &ListServicePolicyPostureChecksServiceUnavailable{}
}

/* ListServicePolicyPostureChecksServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type ListServicePolicyPostureChecksServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyPostureChecksServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/posture-checks][%d] listServicePolicyPostureChecksServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListServicePolicyPostureChecksServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyPostureChecksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

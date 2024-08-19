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

package edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// UpdateEdgeRouterPolicyReader is a Reader for the UpdateEdgeRouterPolicy structure.
type UpdateEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEdgeRouterPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateEdgeRouterPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEdgeRouterPolicyOK creates a UpdateEdgeRouterPolicyOK with default headers values
func NewUpdateEdgeRouterPolicyOK() *UpdateEdgeRouterPolicyOK {
	return &UpdateEdgeRouterPolicyOK{}
}

/*
UpdateEdgeRouterPolicyOK describes a response with status code 200, with default header values.

The update request was successful and the resource has been altered
*/
type UpdateEdgeRouterPolicyOK struct {
	Payload *rest_model.Empty
}

func (o *UpdateEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /edge-router-policies/{id}][%d] updateEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *UpdateEdgeRouterPolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdateEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeRouterPolicyBadRequest creates a UpdateEdgeRouterPolicyBadRequest with default headers values
func NewUpdateEdgeRouterPolicyBadRequest() *UpdateEdgeRouterPolicyBadRequest {
	return &UpdateEdgeRouterPolicyBadRequest{}
}

/*
UpdateEdgeRouterPolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdateEdgeRouterPolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateEdgeRouterPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /edge-router-policies/{id}][%d] updateEdgeRouterPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateEdgeRouterPolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateEdgeRouterPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeRouterPolicyUnauthorized creates a UpdateEdgeRouterPolicyUnauthorized with default headers values
func NewUpdateEdgeRouterPolicyUnauthorized() *UpdateEdgeRouterPolicyUnauthorized {
	return &UpdateEdgeRouterPolicyUnauthorized{}
}

/*
UpdateEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type UpdateEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /edge-router-policies/{id}][%d] updateEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeRouterPolicyNotFound creates a UpdateEdgeRouterPolicyNotFound with default headers values
func NewUpdateEdgeRouterPolicyNotFound() *UpdateEdgeRouterPolicyNotFound {
	return &UpdateEdgeRouterPolicyNotFound{}
}

/*
UpdateEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type UpdateEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /edge-router-policies/{id}][%d] updateEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeRouterPolicyTooManyRequests creates a UpdateEdgeRouterPolicyTooManyRequests with default headers values
func NewUpdateEdgeRouterPolicyTooManyRequests() *UpdateEdgeRouterPolicyTooManyRequests {
	return &UpdateEdgeRouterPolicyTooManyRequests{}
}

/*
UpdateEdgeRouterPolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type UpdateEdgeRouterPolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateEdgeRouterPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /edge-router-policies/{id}][%d] updateEdgeRouterPolicyTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateEdgeRouterPolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateEdgeRouterPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

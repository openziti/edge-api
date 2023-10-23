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

// DetailEdgeRouterPolicyReader is a Reader for the DetailEdgeRouterPolicy structure.
type DetailEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailEdgeRouterPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailEdgeRouterPolicyOK creates a DetailEdgeRouterPolicyOK with default headers values
func NewDetailEdgeRouterPolicyOK() *DetailEdgeRouterPolicyOK {
	return &DetailEdgeRouterPolicyOK{}
}

/* DetailEdgeRouterPolicyOK describes a response with status code 200, with default header values.

A single edge router policy
*/
type DetailEdgeRouterPolicyOK struct {
	Payload *rest_model.DetailEdgeRouterPolicyEnvelope
}

func (o *DetailEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}][%d] detailEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *DetailEdgeRouterPolicyOK) GetPayload() *rest_model.DetailEdgeRouterPolicyEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailEdgeRouterPolicyEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailEdgeRouterPolicyUnauthorized creates a DetailEdgeRouterPolicyUnauthorized with default headers values
func NewDetailEdgeRouterPolicyUnauthorized() *DetailEdgeRouterPolicyUnauthorized {
	return &DetailEdgeRouterPolicyUnauthorized{}
}

/* DetailEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}][%d] detailEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailEdgeRouterPolicyNotFound creates a DetailEdgeRouterPolicyNotFound with default headers values
func NewDetailEdgeRouterPolicyNotFound() *DetailEdgeRouterPolicyNotFound {
	return &DetailEdgeRouterPolicyNotFound{}
}

/* DetailEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}][%d] detailEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *DetailEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailEdgeRouterPolicyTooManyRequests creates a DetailEdgeRouterPolicyTooManyRequests with default headers values
func NewDetailEdgeRouterPolicyTooManyRequests() *DetailEdgeRouterPolicyTooManyRequests {
	return &DetailEdgeRouterPolicyTooManyRequests{}
}

/* DetailEdgeRouterPolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailEdgeRouterPolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailEdgeRouterPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}][%d] detailEdgeRouterPolicyTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailEdgeRouterPolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

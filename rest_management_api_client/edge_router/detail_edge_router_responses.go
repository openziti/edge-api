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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DetailEdgeRouterReader is a Reader for the DetailEdgeRouter structure.
type DetailEdgeRouterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailEdgeRouterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailEdgeRouterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailEdgeRouterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailEdgeRouterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailEdgeRouterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailEdgeRouterOK creates a DetailEdgeRouterOK with default headers values
func NewDetailEdgeRouterOK() *DetailEdgeRouterOK {
	return &DetailEdgeRouterOK{}
}

/* DetailEdgeRouterOK describes a response with status code 200, with default header values.

A singular edge router resource
*/
type DetailEdgeRouterOK struct {
	Payload *rest_model.DetailedEdgeRouterEnvelope
}

func (o *DetailEdgeRouterOK) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}][%d] detailEdgeRouterOK  %+v", 200, o.Payload)
}
func (o *DetailEdgeRouterOK) GetPayload() *rest_model.DetailedEdgeRouterEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailedEdgeRouterEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailEdgeRouterUnauthorized creates a DetailEdgeRouterUnauthorized with default headers values
func NewDetailEdgeRouterUnauthorized() *DetailEdgeRouterUnauthorized {
	return &DetailEdgeRouterUnauthorized{}
}

/* DetailEdgeRouterUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailEdgeRouterUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailEdgeRouterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}][%d] detailEdgeRouterUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailEdgeRouterUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailEdgeRouterNotFound creates a DetailEdgeRouterNotFound with default headers values
func NewDetailEdgeRouterNotFound() *DetailEdgeRouterNotFound {
	return &DetailEdgeRouterNotFound{}
}

/* DetailEdgeRouterNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailEdgeRouterNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailEdgeRouterNotFound) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}][%d] detailEdgeRouterNotFound  %+v", 404, o.Payload)
}
func (o *DetailEdgeRouterNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailEdgeRouterTooManyRequests creates a DetailEdgeRouterTooManyRequests with default headers values
func NewDetailEdgeRouterTooManyRequests() *DetailEdgeRouterTooManyRequests {
	return &DetailEdgeRouterTooManyRequests{}
}

/* DetailEdgeRouterTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailEdgeRouterTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailEdgeRouterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}][%d] detailEdgeRouterTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailEdgeRouterTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailEdgeRouterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

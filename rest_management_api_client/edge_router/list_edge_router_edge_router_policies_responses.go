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

// ListEdgeRouterEdgeRouterPoliciesReader is a Reader for the ListEdgeRouterEdgeRouterPolicies structure.
type ListEdgeRouterEdgeRouterPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEdgeRouterEdgeRouterPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEdgeRouterEdgeRouterPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEdgeRouterEdgeRouterPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListEdgeRouterEdgeRouterPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListEdgeRouterEdgeRouterPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEdgeRouterEdgeRouterPoliciesOK creates a ListEdgeRouterEdgeRouterPoliciesOK with default headers values
func NewListEdgeRouterEdgeRouterPoliciesOK() *ListEdgeRouterEdgeRouterPoliciesOK {
	return &ListEdgeRouterEdgeRouterPoliciesOK{}
}

/*
ListEdgeRouterEdgeRouterPoliciesOK describes a response with status code 200, with default header values.

A list of edge router policies
*/
type ListEdgeRouterEdgeRouterPoliciesOK struct {
	Payload *rest_model.ListEdgeRouterPoliciesEnvelope
}

func (o *ListEdgeRouterEdgeRouterPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}/edge-router-policies][%d] listEdgeRouterEdgeRouterPoliciesOK  %+v", 200, o.Payload)
}
func (o *ListEdgeRouterEdgeRouterPoliciesOK) GetPayload() *rest_model.ListEdgeRouterPoliciesEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterEdgeRouterPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRouterPoliciesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterEdgeRouterPoliciesUnauthorized creates a ListEdgeRouterEdgeRouterPoliciesUnauthorized with default headers values
func NewListEdgeRouterEdgeRouterPoliciesUnauthorized() *ListEdgeRouterEdgeRouterPoliciesUnauthorized {
	return &ListEdgeRouterEdgeRouterPoliciesUnauthorized{}
}

/*
ListEdgeRouterEdgeRouterPoliciesUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListEdgeRouterEdgeRouterPoliciesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterEdgeRouterPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}/edge-router-policies][%d] listEdgeRouterEdgeRouterPoliciesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListEdgeRouterEdgeRouterPoliciesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterEdgeRouterPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterEdgeRouterPoliciesNotFound creates a ListEdgeRouterEdgeRouterPoliciesNotFound with default headers values
func NewListEdgeRouterEdgeRouterPoliciesNotFound() *ListEdgeRouterEdgeRouterPoliciesNotFound {
	return &ListEdgeRouterEdgeRouterPoliciesNotFound{}
}

/*
ListEdgeRouterEdgeRouterPoliciesNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ListEdgeRouterEdgeRouterPoliciesNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterEdgeRouterPoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}/edge-router-policies][%d] listEdgeRouterEdgeRouterPoliciesNotFound  %+v", 404, o.Payload)
}
func (o *ListEdgeRouterEdgeRouterPoliciesNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterEdgeRouterPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterEdgeRouterPoliciesTooManyRequests creates a ListEdgeRouterEdgeRouterPoliciesTooManyRequests with default headers values
func NewListEdgeRouterEdgeRouterPoliciesTooManyRequests() *ListEdgeRouterEdgeRouterPoliciesTooManyRequests {
	return &ListEdgeRouterEdgeRouterPoliciesTooManyRequests{}
}

/*
ListEdgeRouterEdgeRouterPoliciesTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListEdgeRouterEdgeRouterPoliciesTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterEdgeRouterPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /edge-routers/{id}/edge-router-policies][%d] listEdgeRouterEdgeRouterPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListEdgeRouterEdgeRouterPoliciesTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterEdgeRouterPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

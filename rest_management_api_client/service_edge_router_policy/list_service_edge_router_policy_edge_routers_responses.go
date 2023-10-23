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

package service_edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListServiceEdgeRouterPolicyEdgeRoutersReader is a Reader for the ListServiceEdgeRouterPolicyEdgeRouters structure.
type ListServiceEdgeRouterPolicyEdgeRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceEdgeRouterPolicyEdgeRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceEdgeRouterPolicyEdgeRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServiceEdgeRouterPolicyEdgeRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListServiceEdgeRouterPolicyEdgeRoutersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceEdgeRouterPolicyEdgeRoutersOK creates a ListServiceEdgeRouterPolicyEdgeRoutersOK with default headers values
func NewListServiceEdgeRouterPolicyEdgeRoutersOK() *ListServiceEdgeRouterPolicyEdgeRoutersOK {
	return &ListServiceEdgeRouterPolicyEdgeRoutersOK{}
}

/* ListServiceEdgeRouterPolicyEdgeRoutersOK describes a response with status code 200, with default header values.

A list of edge routers
*/
type ListServiceEdgeRouterPolicyEdgeRoutersOK struct {
	Payload *rest_model.ListEdgeRoutersEnvelope
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersOK) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/edge-routers][%d] listServiceEdgeRouterPolicyEdgeRoutersOK  %+v", 200, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyEdgeRoutersOK) GetPayload() *rest_model.ListEdgeRoutersEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPolicyEdgeRoutersUnauthorized creates a ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized with default headers values
func NewListServiceEdgeRouterPolicyEdgeRoutersUnauthorized() *ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized {
	return &ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized{}
}

/* ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/edge-routers][%d] listServiceEdgeRouterPolicyEdgeRoutersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPolicyEdgeRoutersNotFound creates a ListServiceEdgeRouterPolicyEdgeRoutersNotFound with default headers values
func NewListServiceEdgeRouterPolicyEdgeRoutersNotFound() *ListServiceEdgeRouterPolicyEdgeRoutersNotFound {
	return &ListServiceEdgeRouterPolicyEdgeRoutersNotFound{}
}

/* ListServiceEdgeRouterPolicyEdgeRoutersNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ListServiceEdgeRouterPolicyEdgeRoutersNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersNotFound) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/edge-routers][%d] listServiceEdgeRouterPolicyEdgeRoutersNotFound  %+v", 404, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyEdgeRoutersNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests creates a ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests with default headers values
func NewListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests() *ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests {
	return &ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests{}
}

/* ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/edge-routers][%d] listServiceEdgeRouterPolicyEdgeRoutersTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyEdgeRoutersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

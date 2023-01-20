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

	"github.com/openziti/edge-api/go/rest_model"
)

// ListEdgeRouterPolicyEdgeRoutersReader is a Reader for the ListEdgeRouterPolicyEdgeRouters structure.
type ListEdgeRouterPolicyEdgeRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEdgeRouterPolicyEdgeRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEdgeRouterPolicyEdgeRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEdgeRouterPolicyEdgeRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListEdgeRouterPolicyEdgeRoutersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEdgeRouterPolicyEdgeRoutersOK creates a ListEdgeRouterPolicyEdgeRoutersOK with default headers values
func NewListEdgeRouterPolicyEdgeRoutersOK() *ListEdgeRouterPolicyEdgeRoutersOK {
	return &ListEdgeRouterPolicyEdgeRoutersOK{}
}

/* ListEdgeRouterPolicyEdgeRoutersOK describes a response with status code 200, with default header values.

A list of edge routers
*/
type ListEdgeRouterPolicyEdgeRoutersOK struct {
	Payload *rest_model.ListEdgeRoutersEnvelope
}

func (o *ListEdgeRouterPolicyEdgeRoutersOK) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}/edge-routers][%d] listEdgeRouterPolicyEdgeRoutersOK  %+v", 200, o.Payload)
}
func (o *ListEdgeRouterPolicyEdgeRoutersOK) GetPayload() *rest_model.ListEdgeRoutersEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPolicyEdgeRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterPolicyEdgeRoutersUnauthorized creates a ListEdgeRouterPolicyEdgeRoutersUnauthorized with default headers values
func NewListEdgeRouterPolicyEdgeRoutersUnauthorized() *ListEdgeRouterPolicyEdgeRoutersUnauthorized {
	return &ListEdgeRouterPolicyEdgeRoutersUnauthorized{}
}

/* ListEdgeRouterPolicyEdgeRoutersUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListEdgeRouterPolicyEdgeRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterPolicyEdgeRoutersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}/edge-routers][%d] listEdgeRouterPolicyEdgeRoutersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListEdgeRouterPolicyEdgeRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPolicyEdgeRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterPolicyEdgeRoutersNotFound creates a ListEdgeRouterPolicyEdgeRoutersNotFound with default headers values
func NewListEdgeRouterPolicyEdgeRoutersNotFound() *ListEdgeRouterPolicyEdgeRoutersNotFound {
	return &ListEdgeRouterPolicyEdgeRoutersNotFound{}
}

/* ListEdgeRouterPolicyEdgeRoutersNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ListEdgeRouterPolicyEdgeRoutersNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterPolicyEdgeRoutersNotFound) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies/{id}/edge-routers][%d] listEdgeRouterPolicyEdgeRoutersNotFound  %+v", 404, o.Payload)
}
func (o *ListEdgeRouterPolicyEdgeRoutersNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPolicyEdgeRoutersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

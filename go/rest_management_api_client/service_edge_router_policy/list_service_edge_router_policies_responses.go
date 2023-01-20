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

	"github.com/openziti/edge-api/go/rest_model"
)

// ListServiceEdgeRouterPoliciesReader is a Reader for the ListServiceEdgeRouterPolicies structure.
type ListServiceEdgeRouterPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceEdgeRouterPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceEdgeRouterPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServiceEdgeRouterPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListServiceEdgeRouterPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceEdgeRouterPoliciesOK creates a ListServiceEdgeRouterPoliciesOK with default headers values
func NewListServiceEdgeRouterPoliciesOK() *ListServiceEdgeRouterPoliciesOK {
	return &ListServiceEdgeRouterPoliciesOK{}
}

/* ListServiceEdgeRouterPoliciesOK describes a response with status code 200, with default header values.

A list of service edge router policies
*/
type ListServiceEdgeRouterPoliciesOK struct {
	Payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope
}

func (o *ListServiceEdgeRouterPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies][%d] listServiceEdgeRouterPoliciesOK  %+v", 200, o.Payload)
}
func (o *ListServiceEdgeRouterPoliciesOK) GetPayload() *rest_model.ListServiceEdgeRouterPoliciesEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListServiceEdgeRouterPoliciesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPoliciesBadRequest creates a ListServiceEdgeRouterPoliciesBadRequest with default headers values
func NewListServiceEdgeRouterPoliciesBadRequest() *ListServiceEdgeRouterPoliciesBadRequest {
	return &ListServiceEdgeRouterPoliciesBadRequest{}
}

/* ListServiceEdgeRouterPoliciesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListServiceEdgeRouterPoliciesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies][%d] listServiceEdgeRouterPoliciesBadRequest  %+v", 400, o.Payload)
}
func (o *ListServiceEdgeRouterPoliciesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPoliciesUnauthorized creates a ListServiceEdgeRouterPoliciesUnauthorized with default headers values
func NewListServiceEdgeRouterPoliciesUnauthorized() *ListServiceEdgeRouterPoliciesUnauthorized {
	return &ListServiceEdgeRouterPoliciesUnauthorized{}
}

/* ListServiceEdgeRouterPoliciesUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListServiceEdgeRouterPoliciesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies][%d] listServiceEdgeRouterPoliciesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServiceEdgeRouterPoliciesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

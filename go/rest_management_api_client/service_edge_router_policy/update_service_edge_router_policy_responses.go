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

// UpdateServiceEdgeRouterPolicyReader is a Reader for the UpdateServiceEdgeRouterPolicy structure.
type UpdateServiceEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceEdgeRouterPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateServiceEdgeRouterPolicyOK creates a UpdateServiceEdgeRouterPolicyOK with default headers values
func NewUpdateServiceEdgeRouterPolicyOK() *UpdateServiceEdgeRouterPolicyOK {
	return &UpdateServiceEdgeRouterPolicyOK{}
}

/* UpdateServiceEdgeRouterPolicyOK describes a response with status code 200, with default header values.

The update request was successful and the resource has been altered
*/
type UpdateServiceEdgeRouterPolicyOK struct {
	Payload *rest_model.Empty
}

func (o *UpdateServiceEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /service-edge-router-policies/{id}][%d] updateServiceEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *UpdateServiceEdgeRouterPolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdateServiceEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceEdgeRouterPolicyBadRequest creates a UpdateServiceEdgeRouterPolicyBadRequest with default headers values
func NewUpdateServiceEdgeRouterPolicyBadRequest() *UpdateServiceEdgeRouterPolicyBadRequest {
	return &UpdateServiceEdgeRouterPolicyBadRequest{}
}

/* UpdateServiceEdgeRouterPolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdateServiceEdgeRouterPolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateServiceEdgeRouterPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service-edge-router-policies/{id}][%d] updateServiceEdgeRouterPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateServiceEdgeRouterPolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceEdgeRouterPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceEdgeRouterPolicyUnauthorized creates a UpdateServiceEdgeRouterPolicyUnauthorized with default headers values
func NewUpdateServiceEdgeRouterPolicyUnauthorized() *UpdateServiceEdgeRouterPolicyUnauthorized {
	return &UpdateServiceEdgeRouterPolicyUnauthorized{}
}

/* UpdateServiceEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type UpdateServiceEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateServiceEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /service-edge-router-policies/{id}][%d] updateServiceEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateServiceEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceEdgeRouterPolicyNotFound creates a UpdateServiceEdgeRouterPolicyNotFound with default headers values
func NewUpdateServiceEdgeRouterPolicyNotFound() *UpdateServiceEdgeRouterPolicyNotFound {
	return &UpdateServiceEdgeRouterPolicyNotFound{}
}

/* UpdateServiceEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type UpdateServiceEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateServiceEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /service-edge-router-policies/{id}][%d] updateServiceEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *UpdateServiceEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

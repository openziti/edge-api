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

// PatchServiceEdgeRouterPolicyReader is a Reader for the PatchServiceEdgeRouterPolicy structure.
type PatchServiceEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchServiceEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchServiceEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchServiceEdgeRouterPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchServiceEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchServiceEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchServiceEdgeRouterPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchServiceEdgeRouterPolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchServiceEdgeRouterPolicyOK creates a PatchServiceEdgeRouterPolicyOK with default headers values
func NewPatchServiceEdgeRouterPolicyOK() *PatchServiceEdgeRouterPolicyOK {
	return &PatchServiceEdgeRouterPolicyOK{}
}

/* PatchServiceEdgeRouterPolicyOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchServiceEdgeRouterPolicyOK struct {
	Payload *rest_model.Empty
}

func (o *PatchServiceEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /service-edge-router-policies/{id}][%d] patchServiceEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *PatchServiceEdgeRouterPolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchServiceEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceEdgeRouterPolicyBadRequest creates a PatchServiceEdgeRouterPolicyBadRequest with default headers values
func NewPatchServiceEdgeRouterPolicyBadRequest() *PatchServiceEdgeRouterPolicyBadRequest {
	return &PatchServiceEdgeRouterPolicyBadRequest{}
}

/* PatchServiceEdgeRouterPolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchServiceEdgeRouterPolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchServiceEdgeRouterPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /service-edge-router-policies/{id}][%d] patchServiceEdgeRouterPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *PatchServiceEdgeRouterPolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceEdgeRouterPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceEdgeRouterPolicyUnauthorized creates a PatchServiceEdgeRouterPolicyUnauthorized with default headers values
func NewPatchServiceEdgeRouterPolicyUnauthorized() *PatchServiceEdgeRouterPolicyUnauthorized {
	return &PatchServiceEdgeRouterPolicyUnauthorized{}
}

/* PatchServiceEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type PatchServiceEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchServiceEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /service-edge-router-policies/{id}][%d] patchServiceEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchServiceEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceEdgeRouterPolicyNotFound creates a PatchServiceEdgeRouterPolicyNotFound with default headers values
func NewPatchServiceEdgeRouterPolicyNotFound() *PatchServiceEdgeRouterPolicyNotFound {
	return &PatchServiceEdgeRouterPolicyNotFound{}
}

/* PatchServiceEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchServiceEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchServiceEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[PATCH /service-edge-router-policies/{id}][%d] patchServiceEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *PatchServiceEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceEdgeRouterPolicyTooManyRequests creates a PatchServiceEdgeRouterPolicyTooManyRequests with default headers values
func NewPatchServiceEdgeRouterPolicyTooManyRequests() *PatchServiceEdgeRouterPolicyTooManyRequests {
	return &PatchServiceEdgeRouterPolicyTooManyRequests{}
}

/* PatchServiceEdgeRouterPolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchServiceEdgeRouterPolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchServiceEdgeRouterPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /service-edge-router-policies/{id}][%d] patchServiceEdgeRouterPolicyTooManyRequests  %+v", 429, o.Payload)
}
func (o *PatchServiceEdgeRouterPolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceEdgeRouterPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchServiceEdgeRouterPolicyServiceUnavailable creates a PatchServiceEdgeRouterPolicyServiceUnavailable with default headers values
func NewPatchServiceEdgeRouterPolicyServiceUnavailable() *PatchServiceEdgeRouterPolicyServiceUnavailable {
	return &PatchServiceEdgeRouterPolicyServiceUnavailable{}
}

/* PatchServiceEdgeRouterPolicyServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type PatchServiceEdgeRouterPolicyServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchServiceEdgeRouterPolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /service-edge-router-policies/{id}][%d] patchServiceEdgeRouterPolicyServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PatchServiceEdgeRouterPolicyServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchServiceEdgeRouterPolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

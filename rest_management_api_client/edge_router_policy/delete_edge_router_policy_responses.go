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

// DeleteEdgeRouterPolicyReader is a Reader for the DeleteEdgeRouterPolicy structure.
type DeleteEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEdgeRouterPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteEdgeRouterPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteEdgeRouterPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEdgeRouterPolicyOK creates a DeleteEdgeRouterPolicyOK with default headers values
func NewDeleteEdgeRouterPolicyOK() *DeleteEdgeRouterPolicyOK {
	return &DeleteEdgeRouterPolicyOK{}
}

/* DeleteEdgeRouterPolicyOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteEdgeRouterPolicyOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /edge-router-policies/{id}][%d] deleteEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *DeleteEdgeRouterPolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeRouterPolicyBadRequest creates a DeleteEdgeRouterPolicyBadRequest with default headers values
func NewDeleteEdgeRouterPolicyBadRequest() *DeleteEdgeRouterPolicyBadRequest {
	return &DeleteEdgeRouterPolicyBadRequest{}
}

/* DeleteEdgeRouterPolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteEdgeRouterPolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEdgeRouterPolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /edge-router-policies/{id}][%d] deleteEdgeRouterPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteEdgeRouterPolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEdgeRouterPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeRouterPolicyUnauthorized creates a DeleteEdgeRouterPolicyUnauthorized with default headers values
func NewDeleteEdgeRouterPolicyUnauthorized() *DeleteEdgeRouterPolicyUnauthorized {
	return &DeleteEdgeRouterPolicyUnauthorized{}
}

/* DeleteEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /edge-router-policies/{id}][%d] deleteEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeRouterPolicyNotFound creates a DeleteEdgeRouterPolicyNotFound with default headers values
func NewDeleteEdgeRouterPolicyNotFound() *DeleteEdgeRouterPolicyNotFound {
	return &DeleteEdgeRouterPolicyNotFound{}
}

/* DeleteEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DeleteEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /edge-router-policies/{id}][%d] deleteEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *DeleteEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeRouterPolicyConflict creates a DeleteEdgeRouterPolicyConflict with default headers values
func NewDeleteEdgeRouterPolicyConflict() *DeleteEdgeRouterPolicyConflict {
	return &DeleteEdgeRouterPolicyConflict{}
}

/* DeleteEdgeRouterPolicyConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteEdgeRouterPolicyConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEdgeRouterPolicyConflict) Error() string {
	return fmt.Sprintf("[DELETE /edge-router-policies/{id}][%d] deleteEdgeRouterPolicyConflict  %+v", 409, o.Payload)
}
func (o *DeleteEdgeRouterPolicyConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEdgeRouterPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeRouterPolicyTooManyRequests creates a DeleteEdgeRouterPolicyTooManyRequests with default headers values
func NewDeleteEdgeRouterPolicyTooManyRequests() *DeleteEdgeRouterPolicyTooManyRequests {
	return &DeleteEdgeRouterPolicyTooManyRequests{}
}

/* DeleteEdgeRouterPolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteEdgeRouterPolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEdgeRouterPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /edge-router-policies/{id}][%d] deleteEdgeRouterPolicyTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteEdgeRouterPolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEdgeRouterPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

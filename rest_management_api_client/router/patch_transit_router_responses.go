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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// PatchTransitRouterReader is a Reader for the PatchTransitRouter structure.
type PatchTransitRouterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTransitRouterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTransitRouterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchTransitRouterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchTransitRouterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchTransitRouterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchTransitRouterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchTransitRouterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchTransitRouterOK creates a PatchTransitRouterOK with default headers values
func NewPatchTransitRouterOK() *PatchTransitRouterOK {
	return &PatchTransitRouterOK{}
}

/* PatchTransitRouterOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchTransitRouterOK struct {
	Payload *rest_model.Empty
}

func (o *PatchTransitRouterOK) Error() string {
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterOK  %+v", 200, o.Payload)
}
func (o *PatchTransitRouterOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchTransitRouterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterBadRequest creates a PatchTransitRouterBadRequest with default headers values
func NewPatchTransitRouterBadRequest() *PatchTransitRouterBadRequest {
	return &PatchTransitRouterBadRequest{}
}

/* PatchTransitRouterBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchTransitRouterBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchTransitRouterBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterBadRequest  %+v", 400, o.Payload)
}
func (o *PatchTransitRouterBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterUnauthorized creates a PatchTransitRouterUnauthorized with default headers values
func NewPatchTransitRouterUnauthorized() *PatchTransitRouterUnauthorized {
	return &PatchTransitRouterUnauthorized{}
}

/* PatchTransitRouterUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type PatchTransitRouterUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchTransitRouterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchTransitRouterUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterNotFound creates a PatchTransitRouterNotFound with default headers values
func NewPatchTransitRouterNotFound() *PatchTransitRouterNotFound {
	return &PatchTransitRouterNotFound{}
}

/* PatchTransitRouterNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchTransitRouterNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchTransitRouterNotFound) Error() string {
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterNotFound  %+v", 404, o.Payload)
}
func (o *PatchTransitRouterNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterTooManyRequests creates a PatchTransitRouterTooManyRequests with default headers values
func NewPatchTransitRouterTooManyRequests() *PatchTransitRouterTooManyRequests {
	return &PatchTransitRouterTooManyRequests{}
}

/* PatchTransitRouterTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchTransitRouterTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchTransitRouterTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterTooManyRequests  %+v", 429, o.Payload)
}
func (o *PatchTransitRouterTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterServiceUnavailable creates a PatchTransitRouterServiceUnavailable with default headers values
func NewPatchTransitRouterServiceUnavailable() *PatchTransitRouterServiceUnavailable {
	return &PatchTransitRouterServiceUnavailable{}
}

/* PatchTransitRouterServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type PatchTransitRouterServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchTransitRouterServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PatchTransitRouterServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

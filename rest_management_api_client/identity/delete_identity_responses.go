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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteIdentityReader is a Reader for the DeleteIdentity structure.
type DeleteIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteIdentityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityOK creates a DeleteIdentityOK with default headers values
func NewDeleteIdentityOK() *DeleteIdentityOK {
	return &DeleteIdentityOK{}
}

/*
DeleteIdentityOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteIdentityOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteIdentityOK) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityOK  %+v", 200, o.Payload)
}
func (o *DeleteIdentityOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityBadRequest creates a DeleteIdentityBadRequest with default headers values
func NewDeleteIdentityBadRequest() *DeleteIdentityBadRequest {
	return &DeleteIdentityBadRequest{}
}

/*
DeleteIdentityBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteIdentityBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteIdentityBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteIdentityBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityUnauthorized creates a DeleteIdentityUnauthorized with default headers values
func NewDeleteIdentityUnauthorized() *DeleteIdentityUnauthorized {
	return &DeleteIdentityUnauthorized{}
}

/*
DeleteIdentityUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteIdentityUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteIdentityUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteIdentityUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityConflict creates a DeleteIdentityConflict with default headers values
func NewDeleteIdentityConflict() *DeleteIdentityConflict {
	return &DeleteIdentityConflict{}
}

/*
DeleteIdentityConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteIdentityConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteIdentityConflict) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityConflict  %+v", 409, o.Payload)
}
func (o *DeleteIdentityConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityTooManyRequests creates a DeleteIdentityTooManyRequests with default headers values
func NewDeleteIdentityTooManyRequests() *DeleteIdentityTooManyRequests {
	return &DeleteIdentityTooManyRequests{}
}

/*
DeleteIdentityTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteIdentityTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteIdentityTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteIdentityTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

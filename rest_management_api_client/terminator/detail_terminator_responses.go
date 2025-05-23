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

package terminator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DetailTerminatorReader is a Reader for the DetailTerminator structure.
type DetailTerminatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailTerminatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailTerminatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailTerminatorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailTerminatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailTerminatorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDetailTerminatorServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailTerminatorOK creates a DetailTerminatorOK with default headers values
func NewDetailTerminatorOK() *DetailTerminatorOK {
	return &DetailTerminatorOK{}
}

/* DetailTerminatorOK describes a response with status code 200, with default header values.

A single terminator
*/
type DetailTerminatorOK struct {
	Payload *rest_model.DetailTerminatorEnvelope
}

func (o *DetailTerminatorOK) Error() string {
	return fmt.Sprintf("[GET /terminators/{id}][%d] detailTerminatorOK  %+v", 200, o.Payload)
}
func (o *DetailTerminatorOK) GetPayload() *rest_model.DetailTerminatorEnvelope {
	return o.Payload
}

func (o *DetailTerminatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailTerminatorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailTerminatorUnauthorized creates a DetailTerminatorUnauthorized with default headers values
func NewDetailTerminatorUnauthorized() *DetailTerminatorUnauthorized {
	return &DetailTerminatorUnauthorized{}
}

/* DetailTerminatorUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailTerminatorUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailTerminatorUnauthorized) Error() string {
	return fmt.Sprintf("[GET /terminators/{id}][%d] detailTerminatorUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailTerminatorUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailTerminatorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailTerminatorNotFound creates a DetailTerminatorNotFound with default headers values
func NewDetailTerminatorNotFound() *DetailTerminatorNotFound {
	return &DetailTerminatorNotFound{}
}

/* DetailTerminatorNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailTerminatorNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailTerminatorNotFound) Error() string {
	return fmt.Sprintf("[GET /terminators/{id}][%d] detailTerminatorNotFound  %+v", 404, o.Payload)
}
func (o *DetailTerminatorNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailTerminatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailTerminatorTooManyRequests creates a DetailTerminatorTooManyRequests with default headers values
func NewDetailTerminatorTooManyRequests() *DetailTerminatorTooManyRequests {
	return &DetailTerminatorTooManyRequests{}
}

/* DetailTerminatorTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailTerminatorTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailTerminatorTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /terminators/{id}][%d] detailTerminatorTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailTerminatorTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailTerminatorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailTerminatorServiceUnavailable creates a DetailTerminatorServiceUnavailable with default headers values
func NewDetailTerminatorServiceUnavailable() *DetailTerminatorServiceUnavailable {
	return &DetailTerminatorServiceUnavailable{}
}

/* DetailTerminatorServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type DetailTerminatorServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailTerminatorServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /terminators/{id}][%d] detailTerminatorServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DetailTerminatorServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailTerminatorServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

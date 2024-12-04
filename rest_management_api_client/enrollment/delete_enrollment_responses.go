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

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteEnrollmentReader is a Reader for the DeleteEnrollment structure.
type DeleteEnrollmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnrollmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEnrollmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEnrollmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteEnrollmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEnrollmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteEnrollmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEnrollmentOK creates a DeleteEnrollmentOK with default headers values
func NewDeleteEnrollmentOK() *DeleteEnrollmentOK {
	return &DeleteEnrollmentOK{}
}

/* DeleteEnrollmentOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteEnrollmentOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteEnrollmentOK) Error() string {
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentOK  %+v", 200, o.Payload)
}
func (o *DeleteEnrollmentOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteEnrollmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentBadRequest creates a DeleteEnrollmentBadRequest with default headers values
func NewDeleteEnrollmentBadRequest() *DeleteEnrollmentBadRequest {
	return &DeleteEnrollmentBadRequest{}
}

/* DeleteEnrollmentBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteEnrollmentBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEnrollmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteEnrollmentBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentUnauthorized creates a DeleteEnrollmentUnauthorized with default headers values
func NewDeleteEnrollmentUnauthorized() *DeleteEnrollmentUnauthorized {
	return &DeleteEnrollmentUnauthorized{}
}

/* DeleteEnrollmentUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteEnrollmentUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEnrollmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteEnrollmentUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentNotFound creates a DeleteEnrollmentNotFound with default headers values
func NewDeleteEnrollmentNotFound() *DeleteEnrollmentNotFound {
	return &DeleteEnrollmentNotFound{}
}

/* DeleteEnrollmentNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DeleteEnrollmentNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEnrollmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentNotFound  %+v", 404, o.Payload)
}
func (o *DeleteEnrollmentNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentTooManyRequests creates a DeleteEnrollmentTooManyRequests with default headers values
func NewDeleteEnrollmentTooManyRequests() *DeleteEnrollmentTooManyRequests {
	return &DeleteEnrollmentTooManyRequests{}
}

/* DeleteEnrollmentTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteEnrollmentTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteEnrollmentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteEnrollmentTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

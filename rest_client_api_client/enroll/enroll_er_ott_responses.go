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

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// EnrollErOttReader is a Reader for the EnrollErOtt structure.
type EnrollErOttReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnrollErOttReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnrollErOttOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnrollErOttBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewEnrollErOttTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEnrollErOttServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnrollErOttOK creates a EnrollErOttOK with default headers values
func NewEnrollErOttOK() *EnrollErOttOK {
	return &EnrollErOttOK{}
}

/* EnrollErOttOK describes a response with status code 200, with default header values.

A response containing the edge routers signed certificates (server chain, server cert, CAs).
*/
type EnrollErOttOK struct {
	Payload *rest_model.EnrollmentCertsEnvelope
}

func (o *EnrollErOttOK) Error() string {
	return fmt.Sprintf("[POST /enroll/erott][%d] enrollErOttOK  %+v", 200, o.Payload)
}
func (o *EnrollErOttOK) GetPayload() *rest_model.EnrollmentCertsEnvelope {
	return o.Payload
}

func (o *EnrollErOttOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.EnrollmentCertsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnrollErOttBadRequest creates a EnrollErOttBadRequest with default headers values
func NewEnrollErOttBadRequest() *EnrollErOttBadRequest {
	return &EnrollErOttBadRequest{}
}

/* EnrollErOttBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type EnrollErOttBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnrollErOttBadRequest) Error() string {
	return fmt.Sprintf("[POST /enroll/erott][%d] enrollErOttBadRequest  %+v", 400, o.Payload)
}
func (o *EnrollErOttBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnrollErOttBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnrollErOttTooManyRequests creates a EnrollErOttTooManyRequests with default headers values
func NewEnrollErOttTooManyRequests() *EnrollErOttTooManyRequests {
	return &EnrollErOttTooManyRequests{}
}

/* EnrollErOttTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type EnrollErOttTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnrollErOttTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /enroll/erott][%d] enrollErOttTooManyRequests  %+v", 429, o.Payload)
}
func (o *EnrollErOttTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnrollErOttTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnrollErOttServiceUnavailable creates a EnrollErOttServiceUnavailable with default headers values
func NewEnrollErOttServiceUnavailable() *EnrollErOttServiceUnavailable {
	return &EnrollErOttServiceUnavailable{}
}

/* EnrollErOttServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type EnrollErOttServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnrollErOttServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /enroll/erott][%d] enrollErOttServiceUnavailable  %+v", 503, o.Payload)
}
func (o *EnrollErOttServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnrollErOttServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// EnrollOttCaReader is a Reader for the EnrollOttCa structure.
type EnrollOttCaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnrollOttCaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnrollOttCaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnrollOttCaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewEnrollOttCaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEnrollOttCaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnrollOttCaOK creates a EnrollOttCaOK with default headers values
func NewEnrollOttCaOK() *EnrollOttCaOK {
	return &EnrollOttCaOK{}
}

/* EnrollOttCaOK describes a response with status code 200, with default header values.

Base empty response
*/
type EnrollOttCaOK struct {
	Payload *rest_model.Empty
}

func (o *EnrollOttCaOK) Error() string {
	return fmt.Sprintf("[POST /enroll/ottca][%d] enrollOttCaOK  %+v", 200, o.Payload)
}
func (o *EnrollOttCaOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *EnrollOttCaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnrollOttCaBadRequest creates a EnrollOttCaBadRequest with default headers values
func NewEnrollOttCaBadRequest() *EnrollOttCaBadRequest {
	return &EnrollOttCaBadRequest{}
}

/* EnrollOttCaBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type EnrollOttCaBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnrollOttCaBadRequest) Error() string {
	return fmt.Sprintf("[POST /enroll/ottca][%d] enrollOttCaBadRequest  %+v", 400, o.Payload)
}
func (o *EnrollOttCaBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnrollOttCaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnrollOttCaTooManyRequests creates a EnrollOttCaTooManyRequests with default headers values
func NewEnrollOttCaTooManyRequests() *EnrollOttCaTooManyRequests {
	return &EnrollOttCaTooManyRequests{}
}

/* EnrollOttCaTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type EnrollOttCaTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnrollOttCaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /enroll/ottca][%d] enrollOttCaTooManyRequests  %+v", 429, o.Payload)
}
func (o *EnrollOttCaTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnrollOttCaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnrollOttCaServiceUnavailable creates a EnrollOttCaServiceUnavailable with default headers values
func NewEnrollOttCaServiceUnavailable() *EnrollOttCaServiceUnavailable {
	return &EnrollOttCaServiceUnavailable{}
}

/* EnrollOttCaServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type EnrollOttCaServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnrollOttCaServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /enroll/ottca][%d] enrollOttCaServiceUnavailable  %+v", 503, o.Payload)
}
func (o *EnrollOttCaServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnrollOttCaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

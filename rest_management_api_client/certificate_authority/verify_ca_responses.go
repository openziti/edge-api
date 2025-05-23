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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// VerifyCaReader is a Reader for the VerifyCa structure.
type VerifyCaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyCaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifyCaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVerifyCaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVerifyCaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVerifyCaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewVerifyCaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewVerifyCaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerifyCaOK creates a VerifyCaOK with default headers values
func NewVerifyCaOK() *VerifyCaOK {
	return &VerifyCaOK{}
}

/* VerifyCaOK describes a response with status code 200, with default header values.

Base empty response
*/
type VerifyCaOK struct {
	Payload *rest_model.Empty
}

func (o *VerifyCaOK) Error() string {
	return fmt.Sprintf("[POST /cas/{id}/verify][%d] verifyCaOK  %+v", 200, o.Payload)
}
func (o *VerifyCaOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *VerifyCaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyCaBadRequest creates a VerifyCaBadRequest with default headers values
func NewVerifyCaBadRequest() *VerifyCaBadRequest {
	return &VerifyCaBadRequest{}
}

/* VerifyCaBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type VerifyCaBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *VerifyCaBadRequest) Error() string {
	return fmt.Sprintf("[POST /cas/{id}/verify][%d] verifyCaBadRequest  %+v", 400, o.Payload)
}
func (o *VerifyCaBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *VerifyCaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyCaUnauthorized creates a VerifyCaUnauthorized with default headers values
func NewVerifyCaUnauthorized() *VerifyCaUnauthorized {
	return &VerifyCaUnauthorized{}
}

/* VerifyCaUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type VerifyCaUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *VerifyCaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cas/{id}/verify][%d] verifyCaUnauthorized  %+v", 401, o.Payload)
}
func (o *VerifyCaUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *VerifyCaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyCaNotFound creates a VerifyCaNotFound with default headers values
func NewVerifyCaNotFound() *VerifyCaNotFound {
	return &VerifyCaNotFound{}
}

/* VerifyCaNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type VerifyCaNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *VerifyCaNotFound) Error() string {
	return fmt.Sprintf("[POST /cas/{id}/verify][%d] verifyCaNotFound  %+v", 404, o.Payload)
}
func (o *VerifyCaNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *VerifyCaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyCaTooManyRequests creates a VerifyCaTooManyRequests with default headers values
func NewVerifyCaTooManyRequests() *VerifyCaTooManyRequests {
	return &VerifyCaTooManyRequests{}
}

/* VerifyCaTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type VerifyCaTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *VerifyCaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cas/{id}/verify][%d] verifyCaTooManyRequests  %+v", 429, o.Payload)
}
func (o *VerifyCaTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *VerifyCaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyCaServiceUnavailable creates a VerifyCaServiceUnavailable with default headers values
func NewVerifyCaServiceUnavailable() *VerifyCaServiceUnavailable {
	return &VerifyCaServiceUnavailable{}
}

/* VerifyCaServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type VerifyCaServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *VerifyCaServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /cas/{id}/verify][%d] verifyCaServiceUnavailable  %+v", 503, o.Payload)
}
func (o *VerifyCaServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *VerifyCaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

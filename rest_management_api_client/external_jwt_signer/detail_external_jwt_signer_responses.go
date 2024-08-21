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

package external_jwt_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DetailExternalJWTSignerReader is a Reader for the DetailExternalJWTSigner structure.
type DetailExternalJWTSignerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailExternalJWTSignerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailExternalJWTSignerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailExternalJWTSignerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailExternalJWTSignerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailExternalJWTSignerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailExternalJWTSignerOK creates a DetailExternalJWTSignerOK with default headers values
func NewDetailExternalJWTSignerOK() *DetailExternalJWTSignerOK {
	return &DetailExternalJWTSignerOK{}
}

/* DetailExternalJWTSignerOK describes a response with status code 200, with default header values.

A singular External JWT Signer resource
*/
type DetailExternalJWTSignerOK struct {
	Payload *rest_model.DetailExternalJWTSignerEnvelope
}

func (o *DetailExternalJWTSignerOK) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers/{id}][%d] detailExternalJwtSignerOK  %+v", 200, o.Payload)
}
func (o *DetailExternalJWTSignerOK) GetPayload() *rest_model.DetailExternalJWTSignerEnvelope {
	return o.Payload
}

func (o *DetailExternalJWTSignerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailExternalJWTSignerEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailExternalJWTSignerUnauthorized creates a DetailExternalJWTSignerUnauthorized with default headers values
func NewDetailExternalJWTSignerUnauthorized() *DetailExternalJWTSignerUnauthorized {
	return &DetailExternalJWTSignerUnauthorized{}
}

/* DetailExternalJWTSignerUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailExternalJWTSignerUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailExternalJWTSignerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers/{id}][%d] detailExternalJwtSignerUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailExternalJWTSignerUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailExternalJWTSignerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailExternalJWTSignerNotFound creates a DetailExternalJWTSignerNotFound with default headers values
func NewDetailExternalJWTSignerNotFound() *DetailExternalJWTSignerNotFound {
	return &DetailExternalJWTSignerNotFound{}
}

/* DetailExternalJWTSignerNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailExternalJWTSignerNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailExternalJWTSignerNotFound) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers/{id}][%d] detailExternalJwtSignerNotFound  %+v", 404, o.Payload)
}
func (o *DetailExternalJWTSignerNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailExternalJWTSignerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailExternalJWTSignerTooManyRequests creates a DetailExternalJWTSignerTooManyRequests with default headers values
func NewDetailExternalJWTSignerTooManyRequests() *DetailExternalJWTSignerTooManyRequests {
	return &DetailExternalJWTSignerTooManyRequests{}
}

/* DetailExternalJWTSignerTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailExternalJWTSignerTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailExternalJWTSignerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers/{id}][%d] detailExternalJwtSignerTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailExternalJWTSignerTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailExternalJWTSignerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

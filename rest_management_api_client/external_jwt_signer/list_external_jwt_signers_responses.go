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

// ListExternalJWTSignersReader is a Reader for the ListExternalJWTSigners structure.
type ListExternalJWTSignersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalJWTSignersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalJWTSignersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListExternalJWTSignersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListExternalJWTSignersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListExternalJWTSignersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListExternalJWTSignersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListExternalJWTSignersOK creates a ListExternalJWTSignersOK with default headers values
func NewListExternalJWTSignersOK() *ListExternalJWTSignersOK {
	return &ListExternalJWTSignersOK{}
}

/* ListExternalJWTSignersOK describes a response with status code 200, with default header values.

A list of External JWT Signers
*/
type ListExternalJWTSignersOK struct {
	Payload *rest_model.ListExternalJWTSignersEnvelope
}

func (o *ListExternalJWTSignersOK) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers][%d] listExternalJwtSignersOK  %+v", 200, o.Payload)
}
func (o *ListExternalJWTSignersOK) GetPayload() *rest_model.ListExternalJWTSignersEnvelope {
	return o.Payload
}

func (o *ListExternalJWTSignersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListExternalJWTSignersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalJWTSignersBadRequest creates a ListExternalJWTSignersBadRequest with default headers values
func NewListExternalJWTSignersBadRequest() *ListExternalJWTSignersBadRequest {
	return &ListExternalJWTSignersBadRequest{}
}

/* ListExternalJWTSignersBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListExternalJWTSignersBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListExternalJWTSignersBadRequest) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers][%d] listExternalJwtSignersBadRequest  %+v", 400, o.Payload)
}
func (o *ListExternalJWTSignersBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListExternalJWTSignersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalJWTSignersUnauthorized creates a ListExternalJWTSignersUnauthorized with default headers values
func NewListExternalJWTSignersUnauthorized() *ListExternalJWTSignersUnauthorized {
	return &ListExternalJWTSignersUnauthorized{}
}

/* ListExternalJWTSignersUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListExternalJWTSignersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListExternalJWTSignersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers][%d] listExternalJwtSignersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListExternalJWTSignersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListExternalJWTSignersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalJWTSignersTooManyRequests creates a ListExternalJWTSignersTooManyRequests with default headers values
func NewListExternalJWTSignersTooManyRequests() *ListExternalJWTSignersTooManyRequests {
	return &ListExternalJWTSignersTooManyRequests{}
}

/* ListExternalJWTSignersTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListExternalJWTSignersTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListExternalJWTSignersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers][%d] listExternalJwtSignersTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListExternalJWTSignersTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListExternalJWTSignersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalJWTSignersServiceUnavailable creates a ListExternalJWTSignersServiceUnavailable with default headers values
func NewListExternalJWTSignersServiceUnavailable() *ListExternalJWTSignersServiceUnavailable {
	return &ListExternalJWTSignersServiceUnavailable{}
}

/* ListExternalJWTSignersServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type ListExternalJWTSignersServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListExternalJWTSignersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /external-jwt-signers][%d] listExternalJwtSignersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListExternalJWTSignersServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListExternalJWTSignersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

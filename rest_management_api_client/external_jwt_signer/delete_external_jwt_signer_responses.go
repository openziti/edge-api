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

// DeleteExternalJWTSignerReader is a Reader for the DeleteExternalJWTSigner structure.
type DeleteExternalJWTSignerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalJWTSignerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExternalJWTSignerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalJWTSignerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalJWTSignerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalJWTSignerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalJWTSignerOK creates a DeleteExternalJWTSignerOK with default headers values
func NewDeleteExternalJWTSignerOK() *DeleteExternalJWTSignerOK {
	return &DeleteExternalJWTSignerOK{}
}

/*
DeleteExternalJWTSignerOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteExternalJWTSignerOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteExternalJWTSignerOK) Error() string {
	return fmt.Sprintf("[DELETE /external-jwt-signers/{id}][%d] deleteExternalJwtSignerOK  %+v", 200, o.Payload)
}
func (o *DeleteExternalJWTSignerOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteExternalJWTSignerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalJWTSignerBadRequest creates a DeleteExternalJWTSignerBadRequest with default headers values
func NewDeleteExternalJWTSignerBadRequest() *DeleteExternalJWTSignerBadRequest {
	return &DeleteExternalJWTSignerBadRequest{}
}

/*
DeleteExternalJWTSignerBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteExternalJWTSignerBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteExternalJWTSignerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /external-jwt-signers/{id}][%d] deleteExternalJwtSignerBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteExternalJWTSignerBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteExternalJWTSignerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalJWTSignerUnauthorized creates a DeleteExternalJWTSignerUnauthorized with default headers values
func NewDeleteExternalJWTSignerUnauthorized() *DeleteExternalJWTSignerUnauthorized {
	return &DeleteExternalJWTSignerUnauthorized{}
}

/*
DeleteExternalJWTSignerUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteExternalJWTSignerUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteExternalJWTSignerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /external-jwt-signers/{id}][%d] deleteExternalJwtSignerUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteExternalJWTSignerUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteExternalJWTSignerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalJWTSignerTooManyRequests creates a DeleteExternalJWTSignerTooManyRequests with default headers values
func NewDeleteExternalJWTSignerTooManyRequests() *DeleteExternalJWTSignerTooManyRequests {
	return &DeleteExternalJWTSignerTooManyRequests{}
}

/*
DeleteExternalJWTSignerTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteExternalJWTSignerTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteExternalJWTSignerTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /external-jwt-signers/{id}][%d] deleteExternalJwtSignerTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteExternalJWTSignerTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteExternalJWTSignerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

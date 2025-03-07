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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// CreateCurrentAPISessionCertificateReader is a Reader for the CreateCurrentAPISessionCertificate structure.
type CreateCurrentAPISessionCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCurrentAPISessionCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCurrentAPISessionCertificateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCurrentAPISessionCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCurrentAPISessionCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateCurrentAPISessionCertificateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateCurrentAPISessionCertificateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCurrentAPISessionCertificateCreated creates a CreateCurrentAPISessionCertificateCreated with default headers values
func NewCreateCurrentAPISessionCertificateCreated() *CreateCurrentAPISessionCertificateCreated {
	return &CreateCurrentAPISessionCertificateCreated{}
}

/* CreateCurrentAPISessionCertificateCreated describes a response with status code 201, with default header values.

A response of a create API Session certificate
*/
type CreateCurrentAPISessionCertificateCreated struct {
	Payload *rest_model.CreateCurrentAPISessionCertificateEnvelope
}

func (o *CreateCurrentAPISessionCertificateCreated) Error() string {
	return fmt.Sprintf("[POST /current-api-session/certificates][%d] createCurrentApiSessionCertificateCreated  %+v", 201, o.Payload)
}
func (o *CreateCurrentAPISessionCertificateCreated) GetPayload() *rest_model.CreateCurrentAPISessionCertificateEnvelope {
	return o.Payload
}

func (o *CreateCurrentAPISessionCertificateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateCurrentAPISessionCertificateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCurrentAPISessionCertificateBadRequest creates a CreateCurrentAPISessionCertificateBadRequest with default headers values
func NewCreateCurrentAPISessionCertificateBadRequest() *CreateCurrentAPISessionCertificateBadRequest {
	return &CreateCurrentAPISessionCertificateBadRequest{}
}

/* CreateCurrentAPISessionCertificateBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateCurrentAPISessionCertificateBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateCurrentAPISessionCertificateBadRequest) Error() string {
	return fmt.Sprintf("[POST /current-api-session/certificates][%d] createCurrentApiSessionCertificateBadRequest  %+v", 400, o.Payload)
}
func (o *CreateCurrentAPISessionCertificateBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateCurrentAPISessionCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCurrentAPISessionCertificateUnauthorized creates a CreateCurrentAPISessionCertificateUnauthorized with default headers values
func NewCreateCurrentAPISessionCertificateUnauthorized() *CreateCurrentAPISessionCertificateUnauthorized {
	return &CreateCurrentAPISessionCertificateUnauthorized{}
}

/* CreateCurrentAPISessionCertificateUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type CreateCurrentAPISessionCertificateUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateCurrentAPISessionCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /current-api-session/certificates][%d] createCurrentApiSessionCertificateUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateCurrentAPISessionCertificateUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateCurrentAPISessionCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCurrentAPISessionCertificateTooManyRequests creates a CreateCurrentAPISessionCertificateTooManyRequests with default headers values
func NewCreateCurrentAPISessionCertificateTooManyRequests() *CreateCurrentAPISessionCertificateTooManyRequests {
	return &CreateCurrentAPISessionCertificateTooManyRequests{}
}

/* CreateCurrentAPISessionCertificateTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type CreateCurrentAPISessionCertificateTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateCurrentAPISessionCertificateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /current-api-session/certificates][%d] createCurrentApiSessionCertificateTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateCurrentAPISessionCertificateTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateCurrentAPISessionCertificateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCurrentAPISessionCertificateServiceUnavailable creates a CreateCurrentAPISessionCertificateServiceUnavailable with default headers values
func NewCreateCurrentAPISessionCertificateServiceUnavailable() *CreateCurrentAPISessionCertificateServiceUnavailable {
	return &CreateCurrentAPISessionCertificateServiceUnavailable{}
}

/* CreateCurrentAPISessionCertificateServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type CreateCurrentAPISessionCertificateServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateCurrentAPISessionCertificateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /current-api-session/certificates][%d] createCurrentApiSessionCertificateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateCurrentAPISessionCertificateServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateCurrentAPISessionCertificateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

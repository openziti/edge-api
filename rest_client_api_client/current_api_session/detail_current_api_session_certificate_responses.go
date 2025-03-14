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

// DetailCurrentAPISessionCertificateReader is a Reader for the DetailCurrentAPISessionCertificate structure.
type DetailCurrentAPISessionCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailCurrentAPISessionCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailCurrentAPISessionCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailCurrentAPISessionCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailCurrentAPISessionCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailCurrentAPISessionCertificateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDetailCurrentAPISessionCertificateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailCurrentAPISessionCertificateOK creates a DetailCurrentAPISessionCertificateOK with default headers values
func NewDetailCurrentAPISessionCertificateOK() *DetailCurrentAPISessionCertificateOK {
	return &DetailCurrentAPISessionCertificateOK{}
}

/* DetailCurrentAPISessionCertificateOK describes a response with status code 200, with default header values.

A response containing a single API Session certificate
*/
type DetailCurrentAPISessionCertificateOK struct {
	Payload *rest_model.DetailCurrentAPISessionCertificateEnvelope
}

func (o *DetailCurrentAPISessionCertificateOK) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates/{id}][%d] detailCurrentApiSessionCertificateOK  %+v", 200, o.Payload)
}
func (o *DetailCurrentAPISessionCertificateOK) GetPayload() *rest_model.DetailCurrentAPISessionCertificateEnvelope {
	return o.Payload
}

func (o *DetailCurrentAPISessionCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailCurrentAPISessionCertificateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailCurrentAPISessionCertificateUnauthorized creates a DetailCurrentAPISessionCertificateUnauthorized with default headers values
func NewDetailCurrentAPISessionCertificateUnauthorized() *DetailCurrentAPISessionCertificateUnauthorized {
	return &DetailCurrentAPISessionCertificateUnauthorized{}
}

/* DetailCurrentAPISessionCertificateUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailCurrentAPISessionCertificateUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailCurrentAPISessionCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates/{id}][%d] detailCurrentApiSessionCertificateUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailCurrentAPISessionCertificateUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailCurrentAPISessionCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailCurrentAPISessionCertificateNotFound creates a DetailCurrentAPISessionCertificateNotFound with default headers values
func NewDetailCurrentAPISessionCertificateNotFound() *DetailCurrentAPISessionCertificateNotFound {
	return &DetailCurrentAPISessionCertificateNotFound{}
}

/* DetailCurrentAPISessionCertificateNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailCurrentAPISessionCertificateNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailCurrentAPISessionCertificateNotFound) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates/{id}][%d] detailCurrentApiSessionCertificateNotFound  %+v", 404, o.Payload)
}
func (o *DetailCurrentAPISessionCertificateNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailCurrentAPISessionCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailCurrentAPISessionCertificateTooManyRequests creates a DetailCurrentAPISessionCertificateTooManyRequests with default headers values
func NewDetailCurrentAPISessionCertificateTooManyRequests() *DetailCurrentAPISessionCertificateTooManyRequests {
	return &DetailCurrentAPISessionCertificateTooManyRequests{}
}

/* DetailCurrentAPISessionCertificateTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailCurrentAPISessionCertificateTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailCurrentAPISessionCertificateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates/{id}][%d] detailCurrentApiSessionCertificateTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailCurrentAPISessionCertificateTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailCurrentAPISessionCertificateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailCurrentAPISessionCertificateServiceUnavailable creates a DetailCurrentAPISessionCertificateServiceUnavailable with default headers values
func NewDetailCurrentAPISessionCertificateServiceUnavailable() *DetailCurrentAPISessionCertificateServiceUnavailable {
	return &DetailCurrentAPISessionCertificateServiceUnavailable{}
}

/* DetailCurrentAPISessionCertificateServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type DetailCurrentAPISessionCertificateServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailCurrentAPISessionCertificateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates/{id}][%d] detailCurrentApiSessionCertificateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DetailCurrentAPISessionCertificateServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailCurrentAPISessionCertificateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

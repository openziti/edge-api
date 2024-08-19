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

// ListCasReader is a Reader for the ListCas structure.
type ListCasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListCasTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCasOK creates a ListCasOK with default headers values
func NewListCasOK() *ListCasOK {
	return &ListCasOK{}
}

/*
ListCasOK describes a response with status code 200, with default header values.

A list of Certificate Authorities (CAs)
*/
type ListCasOK struct {
	Payload *rest_model.ListCasEnvelope
}

func (o *ListCasOK) Error() string {
	return fmt.Sprintf("[GET /cas][%d] listCasOK  %+v", 200, o.Payload)
}
func (o *ListCasOK) GetPayload() *rest_model.ListCasEnvelope {
	return o.Payload
}

func (o *ListCasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListCasEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCasBadRequest creates a ListCasBadRequest with default headers values
func NewListCasBadRequest() *ListCasBadRequest {
	return &ListCasBadRequest{}
}

/*
ListCasBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListCasBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListCasBadRequest) Error() string {
	return fmt.Sprintf("[GET /cas][%d] listCasBadRequest  %+v", 400, o.Payload)
}
func (o *ListCasBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCasUnauthorized creates a ListCasUnauthorized with default headers values
func NewListCasUnauthorized() *ListCasUnauthorized {
	return &ListCasUnauthorized{}
}

/*
ListCasUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListCasUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListCasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cas][%d] listCasUnauthorized  %+v", 401, o.Payload)
}
func (o *ListCasUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCasTooManyRequests creates a ListCasTooManyRequests with default headers values
func NewListCasTooManyRequests() *ListCasTooManyRequests {
	return &ListCasTooManyRequests{}
}

/*
ListCasTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListCasTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListCasTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cas][%d] listCasTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListCasTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCasTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

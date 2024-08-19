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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListPostureCheckTypesReader is a Reader for the ListPostureCheckTypes structure.
type ListPostureCheckTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPostureCheckTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPostureCheckTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPostureCheckTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListPostureCheckTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListPostureCheckTypesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPostureCheckTypesOK creates a ListPostureCheckTypesOK with default headers values
func NewListPostureCheckTypesOK() *ListPostureCheckTypesOK {
	return &ListPostureCheckTypesOK{}
}

/*
ListPostureCheckTypesOK describes a response with status code 200, with default header values.

A list of posture check types
*/
type ListPostureCheckTypesOK struct {
	Payload *rest_model.ListPostureCheckTypesEnvelope
}

func (o *ListPostureCheckTypesOK) Error() string {
	return fmt.Sprintf("[GET /posture-check-types][%d] listPostureCheckTypesOK  %+v", 200, o.Payload)
}
func (o *ListPostureCheckTypesOK) GetPayload() *rest_model.ListPostureCheckTypesEnvelope {
	return o.Payload
}

func (o *ListPostureCheckTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListPostureCheckTypesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostureCheckTypesBadRequest creates a ListPostureCheckTypesBadRequest with default headers values
func NewListPostureCheckTypesBadRequest() *ListPostureCheckTypesBadRequest {
	return &ListPostureCheckTypesBadRequest{}
}

/*
ListPostureCheckTypesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListPostureCheckTypesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListPostureCheckTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /posture-check-types][%d] listPostureCheckTypesBadRequest  %+v", 400, o.Payload)
}
func (o *ListPostureCheckTypesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListPostureCheckTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostureCheckTypesUnauthorized creates a ListPostureCheckTypesUnauthorized with default headers values
func NewListPostureCheckTypesUnauthorized() *ListPostureCheckTypesUnauthorized {
	return &ListPostureCheckTypesUnauthorized{}
}

/*
ListPostureCheckTypesUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListPostureCheckTypesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListPostureCheckTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /posture-check-types][%d] listPostureCheckTypesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListPostureCheckTypesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListPostureCheckTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostureCheckTypesTooManyRequests creates a ListPostureCheckTypesTooManyRequests with default headers values
func NewListPostureCheckTypesTooManyRequests() *ListPostureCheckTypesTooManyRequests {
	return &ListPostureCheckTypesTooManyRequests{}
}

/*
ListPostureCheckTypesTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListPostureCheckTypesTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListPostureCheckTypesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /posture-check-types][%d] listPostureCheckTypesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListPostureCheckTypesTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListPostureCheckTypesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

package controllers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListControllersReader is a Reader for the ListControllers structure.
type ListControllersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListControllersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListControllersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListControllersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListControllersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListControllersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListControllersOK creates a ListControllersOK with default headers values
func NewListControllersOK() *ListControllersOK {
	return &ListControllersOK{}
}

/* ListControllersOK describes a response with status code 200, with default header values.

A list of controllers
*/
type ListControllersOK struct {
	Payload *rest_model.ListControllersEnvelope
}

func (o *ListControllersOK) Error() string {
	return fmt.Sprintf("[GET /controllers][%d] listControllersOK  %+v", 200, o.Payload)
}
func (o *ListControllersOK) GetPayload() *rest_model.ListControllersEnvelope {
	return o.Payload
}

func (o *ListControllersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListControllersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListControllersBadRequest creates a ListControllersBadRequest with default headers values
func NewListControllersBadRequest() *ListControllersBadRequest {
	return &ListControllersBadRequest{}
}

/* ListControllersBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListControllersBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListControllersBadRequest) Error() string {
	return fmt.Sprintf("[GET /controllers][%d] listControllersBadRequest  %+v", 400, o.Payload)
}
func (o *ListControllersBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListControllersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListControllersUnauthorized creates a ListControllersUnauthorized with default headers values
func NewListControllersUnauthorized() *ListControllersUnauthorized {
	return &ListControllersUnauthorized{}
}

/* ListControllersUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListControllersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListControllersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controllers][%d] listControllersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListControllersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListControllersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListControllersTooManyRequests creates a ListControllersTooManyRequests with default headers values
func NewListControllersTooManyRequests() *ListControllersTooManyRequests {
	return &ListControllersTooManyRequests{}
}

/* ListControllersTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListControllersTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListControllersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /controllers][%d] listControllersTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListControllersTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListControllersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
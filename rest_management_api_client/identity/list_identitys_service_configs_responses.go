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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListIdentitysServiceConfigsReader is a Reader for the ListIdentitysServiceConfigs structure.
type ListIdentitysServiceConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIdentitysServiceConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIdentitysServiceConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListIdentitysServiceConfigsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListIdentitysServiceConfigsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListIdentitysServiceConfigsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListIdentitysServiceConfigsOK creates a ListIdentitysServiceConfigsOK with default headers values
func NewListIdentitysServiceConfigsOK() *ListIdentitysServiceConfigsOK {
	return &ListIdentitysServiceConfigsOK{}
}

/* ListIdentitysServiceConfigsOK describes a response with status code 200, with default header values.

A list of service configs
*/
type ListIdentitysServiceConfigsOK struct {
	Payload *rest_model.ListServiceConfigsEnvelope
}

func (o *ListIdentitysServiceConfigsOK) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/service-configs][%d] listIdentitysServiceConfigsOK  %+v", 200, o.Payload)
}
func (o *ListIdentitysServiceConfigsOK) GetPayload() *rest_model.ListServiceConfigsEnvelope {
	return o.Payload
}

func (o *ListIdentitysServiceConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListServiceConfigsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentitysServiceConfigsUnauthorized creates a ListIdentitysServiceConfigsUnauthorized with default headers values
func NewListIdentitysServiceConfigsUnauthorized() *ListIdentitysServiceConfigsUnauthorized {
	return &ListIdentitysServiceConfigsUnauthorized{}
}

/* ListIdentitysServiceConfigsUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListIdentitysServiceConfigsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentitysServiceConfigsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/service-configs][%d] listIdentitysServiceConfigsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListIdentitysServiceConfigsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentitysServiceConfigsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentitysServiceConfigsNotFound creates a ListIdentitysServiceConfigsNotFound with default headers values
func NewListIdentitysServiceConfigsNotFound() *ListIdentitysServiceConfigsNotFound {
	return &ListIdentitysServiceConfigsNotFound{}
}

/* ListIdentitysServiceConfigsNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ListIdentitysServiceConfigsNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentitysServiceConfigsNotFound) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/service-configs][%d] listIdentitysServiceConfigsNotFound  %+v", 404, o.Payload)
}
func (o *ListIdentitysServiceConfigsNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentitysServiceConfigsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentitysServiceConfigsTooManyRequests creates a ListIdentitysServiceConfigsTooManyRequests with default headers values
func NewListIdentitysServiceConfigsTooManyRequests() *ListIdentitysServiceConfigsTooManyRequests {
	return &ListIdentitysServiceConfigsTooManyRequests{}
}

/* ListIdentitysServiceConfigsTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListIdentitysServiceConfigsTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentitysServiceConfigsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/service-configs][%d] listIdentitysServiceConfigsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListIdentitysServiceConfigsTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentitysServiceConfigsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

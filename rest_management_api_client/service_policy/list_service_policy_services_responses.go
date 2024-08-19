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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListServicePolicyServicesReader is a Reader for the ListServicePolicyServices structure.
type ListServicePolicyServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicePolicyServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicePolicyServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServicePolicyServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListServicePolicyServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServicePolicyServicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServicePolicyServicesOK creates a ListServicePolicyServicesOK with default headers values
func NewListServicePolicyServicesOK() *ListServicePolicyServicesOK {
	return &ListServicePolicyServicesOK{}
}

/*
ListServicePolicyServicesOK describes a response with status code 200, with default header values.

A list of services
*/
type ListServicePolicyServicesOK struct {
	Payload *rest_model.ListServicesEnvelope
}

func (o *ListServicePolicyServicesOK) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/services][%d] listServicePolicyServicesOK  %+v", 200, o.Payload)
}
func (o *ListServicePolicyServicesOK) GetPayload() *rest_model.ListServicesEnvelope {
	return o.Payload
}

func (o *ListServicePolicyServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListServicesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyServicesBadRequest creates a ListServicePolicyServicesBadRequest with default headers values
func NewListServicePolicyServicesBadRequest() *ListServicePolicyServicesBadRequest {
	return &ListServicePolicyServicesBadRequest{}
}

/*
ListServicePolicyServicesBadRequest describes a response with status code 400, with default header values.

The requested resource does not exist
*/
type ListServicePolicyServicesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyServicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/services][%d] listServicePolicyServicesBadRequest  %+v", 400, o.Payload)
}
func (o *ListServicePolicyServicesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyServicesUnauthorized creates a ListServicePolicyServicesUnauthorized with default headers values
func NewListServicePolicyServicesUnauthorized() *ListServicePolicyServicesUnauthorized {
	return &ListServicePolicyServicesUnauthorized{}
}

/*
ListServicePolicyServicesUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListServicePolicyServicesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/services][%d] listServicePolicyServicesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServicePolicyServicesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicePolicyServicesTooManyRequests creates a ListServicePolicyServicesTooManyRequests with default headers values
func NewListServicePolicyServicesTooManyRequests() *ListServicePolicyServicesTooManyRequests {
	return &ListServicePolicyServicesTooManyRequests{}
}

/*
ListServicePolicyServicesTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListServicePolicyServicesTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServicePolicyServicesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /service-policies/{id}/services][%d] listServicePolicyServicesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServicePolicyServicesTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServicePolicyServicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

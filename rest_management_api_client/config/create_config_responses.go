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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// CreateConfigReader is a Reader for the CreateConfig structure.
type CreateConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConfigCreated creates a CreateConfigCreated with default headers values
func NewCreateConfigCreated() *CreateConfigCreated {
	return &CreateConfigCreated{}
}

/* CreateConfigCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateConfigCreated struct {
	Payload *rest_model.CreateEnvelope
}

func (o *CreateConfigCreated) Error() string {
	return fmt.Sprintf("[POST /configs][%d] createConfigCreated  %+v", 201, o.Payload)
}
func (o *CreateConfigCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigBadRequest creates a CreateConfigBadRequest with default headers values
func NewCreateConfigBadRequest() *CreateConfigBadRequest {
	return &CreateConfigBadRequest{}
}

/* CreateConfigBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateConfigBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /configs][%d] createConfigBadRequest  %+v", 400, o.Payload)
}
func (o *CreateConfigBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigUnauthorized creates a CreateConfigUnauthorized with default headers values
func NewCreateConfigUnauthorized() *CreateConfigUnauthorized {
	return &CreateConfigUnauthorized{}
}

/* CreateConfigUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type CreateConfigUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /configs][%d] createConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateConfigUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigTooManyRequests creates a CreateConfigTooManyRequests with default headers values
func NewCreateConfigTooManyRequests() *CreateConfigTooManyRequests {
	return &CreateConfigTooManyRequests{}
}

/* CreateConfigTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type CreateConfigTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /configs][%d] createConfigTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateConfigTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

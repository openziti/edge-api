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

// DeleteConfigReader is a Reader for the DeleteConfig structure.
type DeleteConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteConfigConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConfigOK creates a DeleteConfigOK with default headers values
func NewDeleteConfigOK() *DeleteConfigOK {
	return &DeleteConfigOK{}
}

/* DeleteConfigOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteConfigOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /configs/{id}][%d] deleteConfigOK  %+v", 200, o.Payload)
}
func (o *DeleteConfigOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigBadRequest creates a DeleteConfigBadRequest with default headers values
func NewDeleteConfigBadRequest() *DeleteConfigBadRequest {
	return &DeleteConfigBadRequest{}
}

/* DeleteConfigBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteConfigBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /configs/{id}][%d] deleteConfigBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteConfigBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigUnauthorized creates a DeleteConfigUnauthorized with default headers values
func NewDeleteConfigUnauthorized() *DeleteConfigUnauthorized {
	return &DeleteConfigUnauthorized{}
}

/* DeleteConfigUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteConfigUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteConfigUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /configs/{id}][%d] deleteConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteConfigUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigNotFound creates a DeleteConfigNotFound with default headers values
func NewDeleteConfigNotFound() *DeleteConfigNotFound {
	return &DeleteConfigNotFound{}
}

/* DeleteConfigNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DeleteConfigNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteConfigNotFound) Error() string {
	return fmt.Sprintf("[DELETE /configs/{id}][%d] deleteConfigNotFound  %+v", 404, o.Payload)
}
func (o *DeleteConfigNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigConflict creates a DeleteConfigConflict with default headers values
func NewDeleteConfigConflict() *DeleteConfigConflict {
	return &DeleteConfigConflict{}
}

/* DeleteConfigConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteConfigConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteConfigConflict) Error() string {
	return fmt.Sprintf("[DELETE /configs/{id}][%d] deleteConfigConflict  %+v", 409, o.Payload)
}
func (o *DeleteConfigConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigTooManyRequests creates a DeleteConfigTooManyRequests with default headers values
func NewDeleteConfigTooManyRequests() *DeleteConfigTooManyRequests {
	return &DeleteConfigTooManyRequests{}
}

/* DeleteConfigTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteConfigTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /configs/{id}][%d] deleteConfigTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteConfigTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

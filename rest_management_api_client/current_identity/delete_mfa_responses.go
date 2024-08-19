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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteMfaReader is a Reader for the DeleteMfa structure.
type DeleteMfaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMfaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMfaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMfaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMfaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMfaOK creates a DeleteMfaOK with default headers values
func NewDeleteMfaOK() *DeleteMfaOK {
	return &DeleteMfaOK{}
}

/*
DeleteMfaOK describes a response with status code 200, with default header values.

Base empty response
*/
type DeleteMfaOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteMfaOK) Error() string {
	return fmt.Sprintf("[DELETE /current-identity/mfa][%d] deleteMfaOK  %+v", 200, o.Payload)
}
func (o *DeleteMfaOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteMfaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMfaUnauthorized creates a DeleteMfaUnauthorized with default headers values
func NewDeleteMfaUnauthorized() *DeleteMfaUnauthorized {
	return &DeleteMfaUnauthorized{}
}

/*
DeleteMfaUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteMfaUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteMfaUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /current-identity/mfa][%d] deleteMfaUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteMfaUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteMfaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMfaNotFound creates a DeleteMfaNotFound with default headers values
func NewDeleteMfaNotFound() *DeleteMfaNotFound {
	return &DeleteMfaNotFound{}
}

/*
DeleteMfaNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DeleteMfaNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteMfaNotFound) Error() string {
	return fmt.Sprintf("[DELETE /current-identity/mfa][%d] deleteMfaNotFound  %+v", 404, o.Payload)
}
func (o *DeleteMfaNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteMfaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

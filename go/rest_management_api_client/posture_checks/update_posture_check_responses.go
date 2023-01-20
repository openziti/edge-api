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

	"github.com/openziti/edge-api/go/rest_model"
)

// UpdatePostureCheckReader is a Reader for the UpdatePostureCheck structure.
type UpdatePostureCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePostureCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePostureCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePostureCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePostureCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePostureCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePostureCheckOK creates a UpdatePostureCheckOK with default headers values
func NewUpdatePostureCheckOK() *UpdatePostureCheckOK {
	return &UpdatePostureCheckOK{}
}

/* UpdatePostureCheckOK describes a response with status code 200, with default header values.

The update request was successful and the resource has been altered
*/
type UpdatePostureCheckOK struct {
	Payload *rest_model.Empty
}

func (o *UpdatePostureCheckOK) Error() string {
	return fmt.Sprintf("[PUT /posture-checks/{id}][%d] updatePostureCheckOK  %+v", 200, o.Payload)
}
func (o *UpdatePostureCheckOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdatePostureCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePostureCheckBadRequest creates a UpdatePostureCheckBadRequest with default headers values
func NewUpdatePostureCheckBadRequest() *UpdatePostureCheckBadRequest {
	return &UpdatePostureCheckBadRequest{}
}

/* UpdatePostureCheckBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdatePostureCheckBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdatePostureCheckBadRequest) Error() string {
	return fmt.Sprintf("[PUT /posture-checks/{id}][%d] updatePostureCheckBadRequest  %+v", 400, o.Payload)
}
func (o *UpdatePostureCheckBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdatePostureCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePostureCheckUnauthorized creates a UpdatePostureCheckUnauthorized with default headers values
func NewUpdatePostureCheckUnauthorized() *UpdatePostureCheckUnauthorized {
	return &UpdatePostureCheckUnauthorized{}
}

/* UpdatePostureCheckUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type UpdatePostureCheckUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdatePostureCheckUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /posture-checks/{id}][%d] updatePostureCheckUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdatePostureCheckUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdatePostureCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePostureCheckNotFound creates a UpdatePostureCheckNotFound with default headers values
func NewUpdatePostureCheckNotFound() *UpdatePostureCheckNotFound {
	return &UpdatePostureCheckNotFound{}
}

/* UpdatePostureCheckNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type UpdatePostureCheckNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdatePostureCheckNotFound) Error() string {
	return fmt.Sprintf("[PUT /posture-checks/{id}][%d] updatePostureCheckNotFound  %+v", 404, o.Payload)
}
func (o *UpdatePostureCheckNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdatePostureCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

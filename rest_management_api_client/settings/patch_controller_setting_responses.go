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

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// PatchControllerSettingReader is a Reader for the PatchControllerSetting structure.
type PatchControllerSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchControllerSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchControllerSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchControllerSettingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchControllerSettingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchControllerSettingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchControllerSettingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchControllerSettingOK creates a PatchControllerSettingOK with default headers values
func NewPatchControllerSettingOK() *PatchControllerSettingOK {
	return &PatchControllerSettingOK{}
}

/* PatchControllerSettingOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchControllerSettingOK struct {
	Payload *rest_model.Empty
}

func (o *PatchControllerSettingOK) Error() string {
	return fmt.Sprintf("[PATCH /controller-settings/{id}/effective][%d] patchControllerSettingOK  %+v", 200, o.Payload)
}
func (o *PatchControllerSettingOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchControllerSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchControllerSettingBadRequest creates a PatchControllerSettingBadRequest with default headers values
func NewPatchControllerSettingBadRequest() *PatchControllerSettingBadRequest {
	return &PatchControllerSettingBadRequest{}
}

/* PatchControllerSettingBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchControllerSettingBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchControllerSettingBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /controller-settings/{id}/effective][%d] patchControllerSettingBadRequest  %+v", 400, o.Payload)
}
func (o *PatchControllerSettingBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchControllerSettingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchControllerSettingUnauthorized creates a PatchControllerSettingUnauthorized with default headers values
func NewPatchControllerSettingUnauthorized() *PatchControllerSettingUnauthorized {
	return &PatchControllerSettingUnauthorized{}
}

/* PatchControllerSettingUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type PatchControllerSettingUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchControllerSettingUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /controller-settings/{id}/effective][%d] patchControllerSettingUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchControllerSettingUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchControllerSettingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchControllerSettingNotFound creates a PatchControllerSettingNotFound with default headers values
func NewPatchControllerSettingNotFound() *PatchControllerSettingNotFound {
	return &PatchControllerSettingNotFound{}
}

/* PatchControllerSettingNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchControllerSettingNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchControllerSettingNotFound) Error() string {
	return fmt.Sprintf("[PATCH /controller-settings/{id}/effective][%d] patchControllerSettingNotFound  %+v", 404, o.Payload)
}
func (o *PatchControllerSettingNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchControllerSettingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchControllerSettingTooManyRequests creates a PatchControllerSettingTooManyRequests with default headers values
func NewPatchControllerSettingTooManyRequests() *PatchControllerSettingTooManyRequests {
	return &PatchControllerSettingTooManyRequests{}
}

/* PatchControllerSettingTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchControllerSettingTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchControllerSettingTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /controller-settings/{id}/effective][%d] patchControllerSettingTooManyRequests  %+v", 429, o.Payload)
}
func (o *PatchControllerSettingTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchControllerSettingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// DeleteCaReader is a Reader for the DeleteCa structure.
type DeleteCaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteCaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCaOK creates a DeleteCaOK with default headers values
func NewDeleteCaOK() *DeleteCaOK {
	return &DeleteCaOK{}
}

/* DeleteCaOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteCaOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteCaOK) Error() string {
	return fmt.Sprintf("[DELETE /cas/{id}][%d] deleteCaOK  %+v", 200, o.Payload)
}
func (o *DeleteCaOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteCaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCaBadRequest creates a DeleteCaBadRequest with default headers values
func NewDeleteCaBadRequest() *DeleteCaBadRequest {
	return &DeleteCaBadRequest{}
}

/* DeleteCaBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteCaBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCaBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cas/{id}][%d] deleteCaBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCaBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCaUnauthorized creates a DeleteCaUnauthorized with default headers values
func NewDeleteCaUnauthorized() *DeleteCaUnauthorized {
	return &DeleteCaUnauthorized{}
}

/* DeleteCaUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteCaUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCaUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cas/{id}][%d] deleteCaUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCaUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCaNotFound creates a DeleteCaNotFound with default headers values
func NewDeleteCaNotFound() *DeleteCaNotFound {
	return &DeleteCaNotFound{}
}

/* DeleteCaNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DeleteCaNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCaNotFound) Error() string {
	return fmt.Sprintf("[DELETE /cas/{id}][%d] deleteCaNotFound  %+v", 404, o.Payload)
}
func (o *DeleteCaNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCaTooManyRequests creates a DeleteCaTooManyRequests with default headers values
func NewDeleteCaTooManyRequests() *DeleteCaTooManyRequests {
	return &DeleteCaTooManyRequests{}
}

/* DeleteCaTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteCaTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCaTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cas/{id}][%d] deleteCaTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteCaTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCaServiceUnavailable creates a DeleteCaServiceUnavailable with default headers values
func NewDeleteCaServiceUnavailable() *DeleteCaServiceUnavailable {
	return &DeleteCaServiceUnavailable{}
}

/* DeleteCaServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type DeleteCaServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCaServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /cas/{id}][%d] deleteCaServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DeleteCaServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

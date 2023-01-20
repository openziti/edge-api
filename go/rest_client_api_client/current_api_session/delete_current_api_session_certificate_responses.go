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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/go/rest_model"
)

// DeleteCurrentAPISessionCertificateReader is a Reader for the DeleteCurrentAPISessionCertificate structure.
type DeleteCurrentAPISessionCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCurrentAPISessionCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCurrentAPISessionCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCurrentAPISessionCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCurrentAPISessionCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCurrentAPISessionCertificateOK creates a DeleteCurrentAPISessionCertificateOK with default headers values
func NewDeleteCurrentAPISessionCertificateOK() *DeleteCurrentAPISessionCertificateOK {
	return &DeleteCurrentAPISessionCertificateOK{}
}

/* DeleteCurrentAPISessionCertificateOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteCurrentAPISessionCertificateOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteCurrentAPISessionCertificateOK) Error() string {
	return fmt.Sprintf("[DELETE /current-api-session/certificates/{id}][%d] deleteCurrentApiSessionCertificateOK  %+v", 200, o.Payload)
}
func (o *DeleteCurrentAPISessionCertificateOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteCurrentAPISessionCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCurrentAPISessionCertificateBadRequest creates a DeleteCurrentAPISessionCertificateBadRequest with default headers values
func NewDeleteCurrentAPISessionCertificateBadRequest() *DeleteCurrentAPISessionCertificateBadRequest {
	return &DeleteCurrentAPISessionCertificateBadRequest{}
}

/* DeleteCurrentAPISessionCertificateBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteCurrentAPISessionCertificateBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCurrentAPISessionCertificateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /current-api-session/certificates/{id}][%d] deleteCurrentApiSessionCertificateBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCurrentAPISessionCertificateBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCurrentAPISessionCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCurrentAPISessionCertificateUnauthorized creates a DeleteCurrentAPISessionCertificateUnauthorized with default headers values
func NewDeleteCurrentAPISessionCertificateUnauthorized() *DeleteCurrentAPISessionCertificateUnauthorized {
	return &DeleteCurrentAPISessionCertificateUnauthorized{}
}

/* DeleteCurrentAPISessionCertificateUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DeleteCurrentAPISessionCertificateUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCurrentAPISessionCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /current-api-session/certificates/{id}][%d] deleteCurrentApiSessionCertificateUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCurrentAPISessionCertificateUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCurrentAPISessionCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

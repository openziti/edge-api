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

package well_known

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListWellKnownCasReader is a Reader for the ListWellKnownCas structure.
type ListWellKnownCasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWellKnownCasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWellKnownCasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListWellKnownCasOK creates a ListWellKnownCasOK with default headers values
func NewListWellKnownCasOK() *ListWellKnownCasOK {
	return &ListWellKnownCasOK{}
}

/* ListWellKnownCasOK describes a response with status code 200, with default header values.

A base64 encoded PKCS7 store
*/
type ListWellKnownCasOK struct {
	Payload string
}

func (o *ListWellKnownCasOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/est/cacerts][%d] listWellKnownCasOK  %+v", 200, o.Payload)
}
func (o *ListWellKnownCasOK) GetPayload() string {
	return o.Payload
}

func (o *ListWellKnownCasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

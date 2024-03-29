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

package informational

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new informational API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for informational API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DetailSpec(params *DetailSpecParams, opts ...ClientOption) (*DetailSpecOK, error)

	DetailSpecBody(params *DetailSpecBodyParams, opts ...ClientOption) (*DetailSpecBodyOK, error)

	ListEnumeratedCapabilities(params *ListEnumeratedCapabilitiesParams, opts ...ClientOption) (*ListEnumeratedCapabilitiesOK, error)

	ListRoot(params *ListRootParams, opts ...ClientOption) (*ListRootOK, error)

	ListSpecs(params *ListSpecsParams, opts ...ClientOption) (*ListSpecsOK, error)

	ListSummary(params *ListSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSummaryOK, error)

	ListVersion(params *ListVersionParams, opts ...ClientOption) (*ListVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DetailSpec returns a single spec resource

  Returns single spec resource embedded within the controller for consumption/documentation/code geneartion
*/
func (a *Client) DetailSpec(params *DetailSpecParams, opts ...ClientOption) (*DetailSpecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailSpec",
		Method:             "GET",
		PathPattern:        "/specs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailSpecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetailSpecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailSpec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DetailSpecBody returns the spec s file

  Return the body of the specification (i.e. Swagger, OpenAPI 2.0, 3.0, etc).
*/
func (a *Client) DetailSpecBody(params *DetailSpecBodyParams, opts ...ClientOption) (*DetailSpecBodyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailSpecBodyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailSpecBody",
		Method:             "GET",
		PathPattern:        "/specs/{id}/spec",
		ProducesMediaTypes: []string{"application/json", "text/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailSpecBodyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetailSpecBodyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailSpecBody: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListEnumeratedCapabilities returns all capabilities this version of the controller is aware of enabled or not
*/
func (a *Client) ListEnumeratedCapabilities(params *ListEnumeratedCapabilitiesParams, opts ...ClientOption) (*ListEnumeratedCapabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEnumeratedCapabilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listEnumeratedCapabilities",
		Method:             "GET",
		PathPattern:        "/enumerated-capabilities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListEnumeratedCapabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListEnumeratedCapabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listEnumeratedCapabilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRoot returns version information
*/
func (a *Client) ListRoot(params *ListRootParams, opts ...ClientOption) (*ListRootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRootParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRoot",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRootReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRootOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRoot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSpecs returns a list of API specs

  Returns a list of spec files embedded within the controller for consumption/documentation/code geneartion
*/
func (a *Client) ListSpecs(params *ListSpecsParams, opts ...ClientOption) (*ListSpecsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSpecsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSpecs",
		Method:             "GET",
		PathPattern:        "/specs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSpecsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSpecsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSpecs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSummary returns a list of accessible resource counts

  This endpoint is usefull for UIs that wish to display UI elements with counts.
*/
func (a *Client) ListSummary(params *ListSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSummary",
		Method:             "GET",
		PathPattern:        "/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListVersion returns version information
*/
func (a *Client) ListVersion(params *ListVersionParams, opts ...ClientOption) (*ListVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

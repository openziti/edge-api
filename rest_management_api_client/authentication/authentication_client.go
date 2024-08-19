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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Authenticate(params *AuthenticateParams, opts ...ClientOption) (*AuthenticateOK, error)

	AuthenticateMfa(params *AuthenticateMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AuthenticateMfaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Authenticate authenticates via a method supplied via a query string parameter

Allowed authentication methods include "password", "cert", and "ext-jwt"
*/
func (a *Client) Authenticate(params *AuthenticateParams, opts ...ClientOption) (*AuthenticateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthenticateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "authenticate",
		Method:             "POST",
		PathPattern:        "/authenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthenticateReader{formats: a.formats},
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
	success, ok := result.(*AuthenticateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authenticate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AuthenticateMfa completes m f a authentication

Completes MFA authentication by submitting a MFA time based one time token or backup code.
*/
func (a *Client) AuthenticateMfa(params *AuthenticateMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AuthenticateMfaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthenticateMfaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "authenticateMfa",
		Method:             "POST",
		PathPattern:        "/authenticate/mfa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthenticateMfaReader{formats: a.formats},
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
	success, ok := result.(*AuthenticateMfaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authenticateMfa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

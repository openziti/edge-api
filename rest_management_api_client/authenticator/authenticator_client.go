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

package authenticator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authenticator API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authenticator API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAuthenticator(params *CreateAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAuthenticatorCreated, error)

	DeleteAuthenticator(params *DeleteAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAuthenticatorOK, error)

	DetailAuthenticator(params *DetailAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailAuthenticatorOK, error)

	ListAuthenticators(params *ListAuthenticatorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAuthenticatorsOK, error)

	PatchAuthenticator(params *PatchAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAuthenticatorOK, error)

	ReEnrollAuthenticator(params *ReEnrollAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReEnrollAuthenticatorCreated, error)

	UpdateAuthenticator(params *UpdateAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAuthenticatorOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAuthenticator creates an authenticator

Creates an authenticator for a specific identity. Requires admin access.
*/
func (a *Client) CreateAuthenticator(params *CreateAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAuthenticatorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAuthenticator",
		Method:             "POST",
		PathPattern:        "/authenticators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*CreateAuthenticatorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DeleteAuthenticator deletes an authenticator

	Delete an authenticator by id. Deleting all authenticators for an identity will make it impossible to log in.

Requires admin access.
*/
func (a *Client) DeleteAuthenticator(params *DeleteAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAuthenticator",
		Method:             "DELETE",
		PathPattern:        "/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailAuthenticator retrieves a single authenticator

Retrieves a single authenticator by id. Requires admin access.
*/
func (a *Client) DetailAuthenticator(params *DetailAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailAuthenticator",
		Method:             "GET",
		PathPattern:        "/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*DetailAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ListAuthenticators lists authenticators

	Returns a list of authenticators associated to identities. The resources can be sorted, filtered, and paginated.

This endpoint requires admin access.
*/
func (a *Client) ListAuthenticators(params *ListAuthenticatorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAuthenticatorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAuthenticatorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAuthenticators",
		Method:             "GET",
		PathPattern:        "/authenticators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAuthenticatorsReader{formats: a.formats},
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
	success, ok := result.(*ListAuthenticatorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAuthenticators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAuthenticator updates the supplied fields on an authenticator

Update the supplied fields on an authenticator by id. Requires admin access.
*/
func (a *Client) PatchAuthenticator(params *PatchAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchAuthenticator",
		Method:             "PATCH",
		PathPattern:        "/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*PatchAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ReEnrollAuthenticator reverts an authenticator to an enrollment

	Allows an authenticator to be reverted to an enrollment and allows re-enrollment to occur. On success the

created enrollment record response is provided and the source authenticator record will be deleted. The
enrollment created depends on the authenticator. UPDB authenticators result in UPDB enrollments, CERT
authenticators result in OTT enrollments, CERT + CA authenticators result in OTTCA enrollments.
*/
func (a *Client) ReEnrollAuthenticator(params *ReEnrollAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReEnrollAuthenticatorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReEnrollAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reEnrollAuthenticator",
		Method:             "POST",
		PathPattern:        "/authenticators/{id}/re-enroll",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReEnrollAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*ReEnrollAuthenticatorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reEnrollAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAuthenticator updates all fields on an authenticator

Update all fields on an authenticator by id. Requires admin access.
*/
func (a *Client) UpdateAuthenticator(params *UpdateAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAuthenticator",
		Method:             "PUT",
		PathPattern:        "/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*UpdateAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

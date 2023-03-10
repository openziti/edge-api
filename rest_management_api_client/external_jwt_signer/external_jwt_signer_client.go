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

package external_jwt_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new external jwt signer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for external jwt signer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateExternalJWTSigner(params *CreateExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExternalJWTSignerCreated, error)

	DeleteExternalJWTSigner(params *DeleteExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExternalJWTSignerOK, error)

	DetailExternalJWTSigner(params *DetailExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailExternalJWTSignerOK, error)

	ListExternalJWTSigners(params *ListExternalJWTSignersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListExternalJWTSignersOK, error)

	PatchExternalJWTSigner(params *PatchExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchExternalJWTSignerOK, error)

	UpdateExternalJWTSigner(params *UpdateExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExternalJWTSignerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateExternalJWTSigner creates an external JWT signer

  Creates an External JWT Signer. Requires admin access.
*/
func (a *Client) CreateExternalJWTSigner(params *CreateExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExternalJWTSignerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExternalJWTSignerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createExternalJwtSigner",
		Method:             "POST",
		PathPattern:        "/external-jwt-signers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateExternalJWTSignerReader{formats: a.formats},
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
	success, ok := result.(*CreateExternalJWTSignerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createExternalJwtSigner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteExternalJWTSigner deletes an external JWT signer

  Delete an External JWT Signer by id. Requires admin access.

*/
func (a *Client) DeleteExternalJWTSigner(params *DeleteExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExternalJWTSignerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExternalJWTSignerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteExternalJwtSigner",
		Method:             "DELETE",
		PathPattern:        "/external-jwt-signers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExternalJWTSignerReader{formats: a.formats},
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
	success, ok := result.(*DeleteExternalJWTSignerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteExternalJwtSigner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DetailExternalJWTSigner retrieves a single external JWT signer

  Retrieves a single External JWT Signer by id. Requires admin access.
*/
func (a *Client) DetailExternalJWTSigner(params *DetailExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailExternalJWTSignerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailExternalJWTSignerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailExternalJwtSigner",
		Method:             "GET",
		PathPattern:        "/external-jwt-signers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailExternalJWTSignerReader{formats: a.formats},
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
	success, ok := result.(*DetailExternalJWTSignerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailExternalJwtSigner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListExternalJWTSigners lists external JWT signers

  Retrieves a list of external JWT signers for authentication
*/
func (a *Client) ListExternalJWTSigners(params *ListExternalJWTSignersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListExternalJWTSignersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListExternalJWTSignersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listExternalJwtSigners",
		Method:             "GET",
		PathPattern:        "/external-jwt-signers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListExternalJWTSignersReader{formats: a.formats},
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
	success, ok := result.(*ListExternalJWTSignersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listExternalJwtSigners: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchExternalJWTSigner updates the supplied fields on an external JWT signer

  Update only the supplied fields on an External JWT Signer by id. Requires admin access.
*/
func (a *Client) PatchExternalJWTSigner(params *PatchExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchExternalJWTSignerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchExternalJWTSignerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchExternalJwtSigner",
		Method:             "PATCH",
		PathPattern:        "/external-jwt-signers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchExternalJWTSignerReader{formats: a.formats},
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
	success, ok := result.(*PatchExternalJWTSignerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchExternalJwtSigner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateExternalJWTSigner updates all fields on an external JWT signer

  Update all fields on an External JWT Signer by id. Requires admin access.
*/
func (a *Client) UpdateExternalJWTSigner(params *UpdateExternalJWTSignerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExternalJWTSignerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExternalJWTSignerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateExternalJwtSigner",
		Method:             "PUT",
		PathPattern:        "/external-jwt-signers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExternalJWTSignerReader{formats: a.formats},
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
	success, ok := result.(*UpdateExternalJWTSignerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateExternalJwtSigner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new current api session API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for current api session API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCurrentAPISession(params *DeleteCurrentAPISessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCurrentAPISessionOK, error)

	DetailCurrentIdentityAuthenticator(params *DetailCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailCurrentIdentityAuthenticatorOK, error)

	ExtendCurrentIdentityAuthenticator(params *ExtendCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtendCurrentIdentityAuthenticatorOK, error)

	ExtendVerifyCurrentIdentityAuthenticator(params *ExtendVerifyCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtendVerifyCurrentIdentityAuthenticatorOK, error)

	GetCurrentAPISession(params *GetCurrentAPISessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentAPISessionOK, error)

	ListCurrentIdentityAuthenticators(params *ListCurrentIdentityAuthenticatorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListCurrentIdentityAuthenticatorsOK, error)

	PatchCurrentIdentityAuthenticator(params *PatchCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchCurrentIdentityAuthenticatorOK, error)

	UpdateCurrentIdentityAuthenticator(params *UpdateCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCurrentIdentityAuthenticatorOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteCurrentAPISession logouts

Terminates the current API session
*/
func (a *Client) DeleteCurrentAPISession(params *DeleteCurrentAPISessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCurrentAPISessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCurrentAPISessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCurrentAPISession",
		Method:             "DELETE",
		PathPattern:        "/current-api-session",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCurrentAPISessionReader{formats: a.formats},
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
	success, ok := result.(*DeleteCurrentAPISessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCurrentAPISession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailCurrentIdentityAuthenticator retrieves an authenticator for the current identity

Retrieves a single authenticator by id. Will only show authenticators assigned to the API session's identity.
*/
func (a *Client) DetailCurrentIdentityAuthenticator(params *DetailCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailCurrentIdentityAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailCurrentIdentityAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailCurrentIdentityAuthenticator",
		Method:             "GET",
		PathPattern:        "/current-identity/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailCurrentIdentityAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*DetailCurrentIdentityAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailCurrentIdentityAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ExtendCurrentIdentityAuthenticator allows the current identity to recieve a new certificate associated with a certificate based authenticator

	This endpoint only functions for certificates issued by the controller. 3rd party certificates are not handled.

Allows an identity to extend its certificate's expiration date by using its current and valid client certificate to submit a CSR. This CSR may be passed in using a new private key, thus allowing private key rotation.
The response from this endpoint is a new client certificate which the client must  be verified via the /authenticators/{id}/extend-verify endpoint.
After verification is completion any new connections must be made with new certificate. Prior to verification the old client certificate remains active.
*/
func (a *Client) ExtendCurrentIdentityAuthenticator(params *ExtendCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtendCurrentIdentityAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtendCurrentIdentityAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extendCurrentIdentityAuthenticator",
		Method:             "POST",
		PathPattern:        "/current-identity/authenticators/{id}/extend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtendCurrentIdentityAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*ExtendCurrentIdentityAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extendCurrentIdentityAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ExtendVerifyCurrentIdentityAuthenticator allows the current identity to validate reciept of a new client certificate

	After submitting a CSR for a new client certificate the resulting public certificate must be re-submitted to this endpoint to verify receipt.

After receipt, the new client certificate must be used for new authentication requests.
*/
func (a *Client) ExtendVerifyCurrentIdentityAuthenticator(params *ExtendVerifyCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtendVerifyCurrentIdentityAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtendVerifyCurrentIdentityAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extendVerifyCurrentIdentityAuthenticator",
		Method:             "POST",
		PathPattern:        "/current-identity/authenticators/{id}/extend-verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtendVerifyCurrentIdentityAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*ExtendVerifyCurrentIdentityAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extendVerifyCurrentIdentityAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCurrentAPISession returns the current API session

Retrieves the API session that was used to issue the current request
*/
func (a *Client) GetCurrentAPISession(params *GetCurrentAPISessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentAPISessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentAPISessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCurrentAPISession",
		Method:             "GET",
		PathPattern:        "/current-api-session",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrentAPISessionReader{formats: a.formats},
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
	success, ok := result.(*GetCurrentAPISessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentAPISession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListCurrentIdentityAuthenticators lists authenticators for the current identity

Retrieves a list of authenticators assigned to the current API session's identity; supports filtering, sorting, and pagination.
*/
func (a *Client) ListCurrentIdentityAuthenticators(params *ListCurrentIdentityAuthenticatorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListCurrentIdentityAuthenticatorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCurrentIdentityAuthenticatorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listCurrentIdentityAuthenticators",
		Method:             "GET",
		PathPattern:        "/current-identity/authenticators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCurrentIdentityAuthenticatorsReader{formats: a.formats},
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
	success, ok := result.(*ListCurrentIdentityAuthenticatorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCurrentIdentityAuthenticators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PatchCurrentIdentityAuthenticator updates the supplied fields on an authenticator of this identity

	Update the supplied fields on an authenticator by id. Will only update authenticators assigned to the API

session's identity.
*/
func (a *Client) PatchCurrentIdentityAuthenticator(params *PatchCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchCurrentIdentityAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCurrentIdentityAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchCurrentIdentityAuthenticator",
		Method:             "PATCH",
		PathPattern:        "/current-identity/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCurrentIdentityAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*PatchCurrentIdentityAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchCurrentIdentityAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UpdateCurrentIdentityAuthenticator updates all fields on an authenticator of this identity

	Update all fields on an authenticator by id.  Will only update authenticators assigned to the API session's

identity.
*/
func (a *Client) UpdateCurrentIdentityAuthenticator(params *UpdateCurrentIdentityAuthenticatorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCurrentIdentityAuthenticatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCurrentIdentityAuthenticatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCurrentIdentityAuthenticator",
		Method:             "PUT",
		PathPattern:        "/current-identity/authenticators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCurrentIdentityAuthenticatorReader{formats: a.formats},
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
	success, ok := result.(*UpdateCurrentIdentityAuthenticatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCurrentIdentityAuthenticator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

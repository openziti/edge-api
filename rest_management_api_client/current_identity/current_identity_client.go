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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new current identity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for current identity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMfaRecoveryCodes(params *CreateMfaRecoveryCodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMfaRecoveryCodesOK, error)

	DeleteMfa(params *DeleteMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMfaOK, error)

	DetailMfa(params *DetailMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailMfaOK, error)

	DetailMfaQrCode(params *DetailMfaQrCodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailMfaQrCodeOK, error)

	DetailMfaRecoveryCodes(params *DetailMfaRecoveryCodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailMfaRecoveryCodesOK, error)

	EnrollMfa(params *EnrollMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnrollMfaCreated, error)

	GetCurrentIdentity(params *GetCurrentIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentIdentityOK, error)

	VerifyMfa(params *VerifyMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VerifyMfaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateMfaRecoveryCodes fors a completed m f a enrollment regenerate the recovery codes

Allows regeneration of recovery codes of an MFA enrollment. Requires a current valid time based one time password to interact with. Available after a completed MFA enrollment. This replaces all existing recovery codes.
*/
func (a *Client) CreateMfaRecoveryCodes(params *CreateMfaRecoveryCodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMfaRecoveryCodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMfaRecoveryCodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMfaRecoveryCodes",
		Method:             "POST",
		PathPattern:        "/current-identity/mfa/recovery-codes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMfaRecoveryCodesReader{formats: a.formats},
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
	success, ok := result.(*CreateMfaRecoveryCodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMfaRecoveryCodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMfa disables m f a for the current identity

Disable MFA for the current identity. Requires a current valid time based one time password if MFA enrollment has been completed. If not, code should be an empty string. If one time passwords are not available and admin account can be used to remove MFA from the identity via `DELETE /identities/<id>/mfa`.
*/
func (a *Client) DeleteMfa(params *DeleteMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMfaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMfaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMfa",
		Method:             "DELETE",
		PathPattern:        "/current-identity/mfa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMfaReader{formats: a.formats},
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
	success, ok := result.(*DeleteMfaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMfa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailMfa returns the current status of m f a enrollment

Returns details about the current MFA enrollment. If enrollment has not been completed it will return the current MFA configuration details necessary to complete a `POST /current-identity/mfa/verify`.
*/
func (a *Client) DetailMfa(params *DetailMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailMfaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailMfaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailMfa",
		Method:             "GET",
		PathPattern:        "/current-identity/mfa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailMfaReader{formats: a.formats},
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
	success, ok := result.(*DetailMfaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailMfa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailMfaQrCode shows a q r code for unverified m f a enrollments

Shows an QR code image for unverified MFA enrollments. 404s if the MFA enrollment has been completed or not started.
*/
func (a *Client) DetailMfaQrCode(params *DetailMfaQrCodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailMfaQrCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailMfaQrCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailMfaQrCode",
		Method:             "GET",
		PathPattern:        "/current-identity/mfa/qr-code",
		ProducesMediaTypes: []string{"application/json", "image/png"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailMfaQrCodeReader{formats: a.formats},
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
	success, ok := result.(*DetailMfaQrCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailMfaQrCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailMfaRecoveryCodes fors a completed m f a enrollment view the current recovery codes

Allows the viewing of recovery codes of an MFA enrollment. Requires a current valid time based one time password to interact with. Available after a completed MFA enrollment.
*/
func (a *Client) DetailMfaRecoveryCodes(params *DetailMfaRecoveryCodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailMfaRecoveryCodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailMfaRecoveryCodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailMfaRecoveryCodes",
		Method:             "GET",
		PathPattern:        "/current-identity/mfa/recovery-codes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailMfaRecoveryCodesReader{formats: a.formats},
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
	success, ok := result.(*DetailMfaRecoveryCodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailMfaRecoveryCodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnrollMfa initiates m f a enrollment

Allows authenticator based MFA enrollment. If enrollment has already been completed, it must be disabled before attempting to re-enroll. Subsequent enrollment request is completed via `POST /current-identity/mfa/verify`
*/
func (a *Client) EnrollMfa(params *EnrollMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnrollMfaCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollMfaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollMfa",
		Method:             "POST",
		PathPattern:        "/current-identity/mfa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollMfaReader{formats: a.formats},
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
	success, ok := result.(*EnrollMfaCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollMfa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCurrentIdentity returns the current identity

Returns the identity associated with the API sessions used to issue the current request
*/
func (a *Client) GetCurrentIdentity(params *GetCurrentIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentIdentityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCurrentIdentity",
		Method:             "GET",
		PathPattern:        "/current-identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrentIdentityReader{formats: a.formats},
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
	success, ok := result.(*GetCurrentIdentityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentIdentity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VerifyMfa completes m f a enrollment by verifying a time based one time token

Completes MFA enrollment by accepting a time based one time password as verification. Called after MFA enrollment has been initiated via `POST /current-identity/mfa`.
*/
func (a *Client) VerifyMfa(params *VerifyMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VerifyMfaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyMfaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "verifyMfa",
		Method:             "POST",
		PathPattern:        "/current-identity/mfa/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VerifyMfaReader{formats: a.formats},
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
	success, ok := result.(*VerifyMfaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for verifyMfa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

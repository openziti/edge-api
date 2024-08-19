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

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new enroll API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for enroll API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Enroll(params *EnrollParams, opts ...ClientOption) (*EnrollOK, error)

	EnrollCa(params *EnrollCaParams, opts ...ClientOption) (*EnrollCaOK, error)

	EnrollErOtt(params *EnrollErOttParams, opts ...ClientOption) (*EnrollErOttOK, error)

	EnrollOtt(params *EnrollOttParams, opts ...ClientOption) (*EnrollOttOK, error)

	EnrollOttCa(params *EnrollOttCaParams, opts ...ClientOption) (*EnrollOttCaOK, error)

	EnrollUpdb(params *EnrollUpdbParams, opts ...ClientOption) (*EnrollUpdbOK, error)

	EnrollmentChallenge(params *EnrollmentChallengeParams, opts ...ClientOption) (*EnrollmentChallengeOK, error)

	ExtendRouterEnrollment(params *ExtendRouterEnrollmentParams, opts ...ClientOption) (*ExtendRouterEnrollmentOK, error)

	GetEnrollmentJwks(params *GetEnrollmentJwksParams, opts ...ClientOption) (*GetEnrollmentJwksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Enroll enrolls an identity via one time token

present a OTT and CSR to receive a long-lived client certificate
*/
func (a *Client) Enroll(params *EnrollParams, opts ...ClientOption) (*EnrollOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enroll",
		Method:             "POST",
		PathPattern:        "/enroll",
		ProducesMediaTypes: []string{"application/json", "application/x-pem-file"},
		ConsumesMediaTypes: []string{"application/json", "application/pkcs10", "application/x-pem-file", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollReader{formats: a.formats},
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
	success, ok := result.(*EnrollOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enroll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EnrollCa enrolls an identity with a pre exchanged certificate

	For CA auto enrollment, an identity is not created beforehand.

Instead one will be created during enrollment. The client will present a client certificate that is signed by a
Certificate Authority that has been added and verified (See POST /cas and POST /cas/{id}/verify).

During this process no CSRs are requires as the client should already be in possession of a valid certificate.
*/
func (a *Client) EnrollCa(params *EnrollCaParams, opts ...ClientOption) (*EnrollCaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollCaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollCa",
		Method:             "POST",
		PathPattern:        "/enroll/ca",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollCaReader{formats: a.formats},
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
	success, ok := result.(*EnrollCaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollCa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnrollErOtt enrolls an edge router

Enrolls an edge-router via a one-time-token to establish a certificate based identity.
*/
func (a *Client) EnrollErOtt(params *EnrollErOttParams, opts ...ClientOption) (*EnrollErOttOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollErOttParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollErOtt",
		Method:             "POST",
		PathPattern:        "/enroll/erott",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollErOttReader{formats: a.formats},
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
	success, ok := result.(*EnrollErOttOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollErOtt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EnrollOtt enrolls an identity via one time token

	Enroll an identity via a one-time-token which is supplied via a query string parameter. This enrollment method

expects a PEM encoded CSRs to be provided for fulfillment. It is up to the enrolling identity to manage the
private key backing the CSR request.
*/
func (a *Client) EnrollOtt(params *EnrollOttParams, opts ...ClientOption) (*EnrollOttOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollOttParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollOtt",
		Method:             "POST",
		PathPattern:        "/enroll/ott",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollOttReader{formats: a.formats},
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
	success, ok := result.(*EnrollOttOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollOtt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EnrollOttCa enrolls an identity via one time token with a pre exchanged client certificate

	Enroll an identity via a one-time-token that also requires a pre-exchanged client certificate to match a

Certificate Authority that has been added and verified (See POST /cas and POST /cas{id}/verify). The client
must present a client certificate signed by CA associated with the enrollment. This enrollment is similar to
CA auto enrollment except that is required the identity to be pre-created.

As the client certificate has been pre-exchanged there is no CSR input to this enrollment method.
*/
func (a *Client) EnrollOttCa(params *EnrollOttCaParams, opts ...ClientOption) (*EnrollOttCaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollOttCaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollOttCa",
		Method:             "POST",
		PathPattern:        "/enroll/ottca",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollOttCaReader{formats: a.formats},
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
	success, ok := result.(*EnrollOttCaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollOttCa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnrollUpdb enrolls an identity via one time token

Enrolls an identity via a one-time-token to establish an initial username and password combination
*/
func (a *Client) EnrollUpdb(params *EnrollUpdbParams, opts ...ClientOption) (*EnrollUpdbOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollUpdbParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollUpdb",
		Method:             "POST",
		PathPattern:        "/enroll/updb",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollUpdbReader{formats: a.formats},
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
	success, ok := result.(*EnrollUpdbOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollUpdb: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EnrollmentChallenge allows verification of a controller or cluster of controllers as being the valid target for enrollment

	A caller may submit a nonce and a key id (kid) from the enrollment JWKS endpoint or enrollment JWT that will

be used to sign the nonce. The resulting signature may be validated with the associated public key in order
to verify a networks identity during enrollment. The nonce must be a valid formatted UUID.
*/
func (a *Client) EnrollmentChallenge(params *EnrollmentChallengeParams, opts ...ClientOption) (*EnrollmentChallengeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnrollmentChallengeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enrollmentChallenge",
		Method:             "POST",
		PathPattern:        "/enroll/challenge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrollmentChallengeReader{formats: a.formats},
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
	success, ok := result.(*EnrollmentChallengeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enrollmentChallenge: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ExtendRouterEnrollment extends the life of a currently enrolled router s certificates

	Allows a router to extend its certificates' expiration date by

using its current and valid client certificate to submit a CSR. This CSR may
be passed in using a new private key, thus allowing private key rotation or swapping.

After completion any new connections must be made with certificates returned from a 200 OK
response. The previous client certificate is rendered invalid for use with the controller even if it
has not expired.

This request must be made using the existing, valid, client certificate.
*/
func (a *Client) ExtendRouterEnrollment(params *ExtendRouterEnrollmentParams, opts ...ClientOption) (*ExtendRouterEnrollmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtendRouterEnrollmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extendRouterEnrollment",
		Method:             "POST",
		PathPattern:        "/enroll/extend/router",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtendRouterEnrollmentReader{formats: a.formats},
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
	success, ok := result.(*ExtendRouterEnrollmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extendRouterEnrollment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetEnrollmentJwks lists JSON web keys associated with enrollment

	Returns a list of JSON Web Keys (JWKS) that are used for enrollment signing. The keys listed here are used

to sign and co-sign enrollment JWTs. They can be verified through a challenge endpoint, using the public keys
from this endpoint to verify the target machine has possession of the related private key.
*/
func (a *Client) GetEnrollmentJwks(params *GetEnrollmentJwksParams, opts ...ClientOption) (*GetEnrollmentJwksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnrollmentJwksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnrollmentJwks",
		Method:             "GET",
		PathPattern:        "/enroll/jwks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEnrollmentJwksReader{formats: a.formats},
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
	success, ok := result.(*GetEnrollmentJwksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnrollmentJwks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

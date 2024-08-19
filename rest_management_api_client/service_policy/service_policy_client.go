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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServicePolicy(params *CreateServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServicePolicyCreated, error)

	DeleteServicePolicy(params *DeleteServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServicePolicyOK, error)

	DetailServicePolicy(params *DetailServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailServicePolicyOK, error)

	ListServicePolicies(params *ListServicePoliciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePoliciesOK, error)

	ListServicePolicyIdentities(params *ListServicePolicyIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePolicyIdentitiesOK, error)

	ListServicePolicyPostureChecks(params *ListServicePolicyPostureChecksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePolicyPostureChecksOK, error)

	ListServicePolicyServices(params *ListServicePolicyServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePolicyServicesOK, error)

	PatchServicePolicy(params *PatchServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchServicePolicyOK, error)

	UpdateServicePolicy(params *UpdateServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateServicePolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateServicePolicy creates a service policy resource

Create a service policy resource. Requires admin access.
*/
func (a *Client) CreateServicePolicy(params *CreateServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServicePolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServicePolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createServicePolicy",
		Method:             "POST",
		PathPattern:        "/service-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServicePolicyReader{formats: a.formats},
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
	success, ok := result.(*CreateServicePolicyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createServicePolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteServicePolicy deletes a service policy

Delete a service policy by id. Requires admin access.
*/
func (a *Client) DeleteServicePolicy(params *DeleteServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServicePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServicePolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteServicePolicy",
		Method:             "DELETE",
		PathPattern:        "/service-policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServicePolicyReader{formats: a.formats},
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
	success, ok := result.(*DeleteServicePolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServicePolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailServicePolicy retrieves a single service policy

Retrieves a single service policy by id. Requires admin access.
*/
func (a *Client) DetailServicePolicy(params *DetailServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailServicePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailServicePolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailServicePolicy",
		Method:             "GET",
		PathPattern:        "/service-policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailServicePolicyReader{formats: a.formats},
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
	success, ok := result.(*DetailServicePolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailServicePolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListServicePolicies lists service policies

Retrieves a list of service policy resources; supports filtering, sorting, and pagination. Requires admin access.
*/
func (a *Client) ListServicePolicies(params *ListServicePoliciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listServicePolicies",
		Method:             "GET",
		PathPattern:        "/service-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServicePoliciesReader{formats: a.formats},
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
	success, ok := result.(*ListServicePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServicePolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListServicePolicyIdentities lists identities a service policy affects

Retrieves a list of identity resources that are affected by a service policy; supports filtering, sorting, and pagination. Requires admin access.
*/
func (a *Client) ListServicePolicyIdentities(params *ListServicePolicyIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePolicyIdentitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicePolicyIdentitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listServicePolicyIdentities",
		Method:             "GET",
		PathPattern:        "/service-policies/{id}/identities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServicePolicyIdentitiesReader{formats: a.formats},
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
	success, ok := result.(*ListServicePolicyIdentitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServicePolicyIdentities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListServicePolicyPostureChecks lists posture check a service policy includes

Retrieves a list of posture check resources that are affected by a service policy; supports filtering, sorting, and pagination. Requires admin access.
*/
func (a *Client) ListServicePolicyPostureChecks(params *ListServicePolicyPostureChecksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePolicyPostureChecksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicePolicyPostureChecksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listServicePolicyPostureChecks",
		Method:             "GET",
		PathPattern:        "/service-policies/{id}/posture-checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServicePolicyPostureChecksReader{formats: a.formats},
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
	success, ok := result.(*ListServicePolicyPostureChecksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServicePolicyPostureChecks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListServicePolicyServices lists services a service policy affects

Retrieves a list of service resources that are affected by a service policy; supports filtering, sorting, and pagination. Requires admin access.
*/
func (a *Client) ListServicePolicyServices(params *ListServicePolicyServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicePolicyServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicePolicyServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listServicePolicyServices",
		Method:             "GET",
		PathPattern:        "/service-policies/{id}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServicePolicyServicesReader{formats: a.formats},
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
	success, ok := result.(*ListServicePolicyServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServicePolicyServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchServicePolicy updates the supplied fields on a service policy

Update the supplied fields on a service policy. Requires admin access.
*/
func (a *Client) PatchServicePolicy(params *PatchServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchServicePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchServicePolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchServicePolicy",
		Method:             "PATCH",
		PathPattern:        "/service-policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchServicePolicyReader{formats: a.formats},
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
	success, ok := result.(*PatchServicePolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchServicePolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateServicePolicy updates all fields on a service policy

Update all fields on a service policy by id. Requires admin access.
*/
func (a *Client) UpdateServicePolicy(params *UpdateServicePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateServicePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServicePolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateServicePolicy",
		Method:             "PUT",
		PathPattern:        "/service-policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServicePolicyReader{formats: a.formats},
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
	success, ok := result.(*UpdateServicePolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateServicePolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RequestExtendAuthenticatorHandlerFunc turns a function with the right signature into a request extend authenticator handler
type RequestExtendAuthenticatorHandlerFunc func(RequestExtendAuthenticatorParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RequestExtendAuthenticatorHandlerFunc) Handle(params RequestExtendAuthenticatorParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RequestExtendAuthenticatorHandler interface for that can handle valid request extend authenticator params
type RequestExtendAuthenticatorHandler interface {
	Handle(RequestExtendAuthenticatorParams, interface{}) middleware.Responder
}

// NewRequestExtendAuthenticator creates a new http.Handler for the request extend authenticator operation
func NewRequestExtendAuthenticator(ctx *middleware.Context, handler RequestExtendAuthenticatorHandler) *RequestExtendAuthenticator {
	return &RequestExtendAuthenticator{Context: ctx, Handler: handler}
}

/* RequestExtendAuthenticator swagger:route POST /authenticators/{id}/request-extend Authenticator requestExtendAuthenticator

Indicate a certificate authenticator should be extended and optionally key rolled on next authentication.

Allows a certificate authenticator to be flagged for early extension and optionally private key rolling.
Connecting clients will receive flags in their API Session indicating that an early extension is request and
a hint on whether private keys should be rolled. Clients that do not support extension or cannot roll keys
may ignore one or both flags.

If this request is made against a non-certificate based authenticator, it will return a 403-forbidden error.


*/
type RequestExtendAuthenticator struct {
	Context *middleware.Context
	Handler RequestExtendAuthenticatorHandler
}

func (o *RequestExtendAuthenticator) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRequestExtendAuthenticatorParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

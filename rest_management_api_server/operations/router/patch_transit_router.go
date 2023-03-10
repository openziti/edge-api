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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchTransitRouterHandlerFunc turns a function with the right signature into a patch transit router handler
type PatchTransitRouterHandlerFunc func(PatchTransitRouterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchTransitRouterHandlerFunc) Handle(params PatchTransitRouterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchTransitRouterHandler interface for that can handle valid patch transit router params
type PatchTransitRouterHandler interface {
	Handle(PatchTransitRouterParams, interface{}) middleware.Responder
}

// NewPatchTransitRouter creates a new http.Handler for the patch transit router operation
func NewPatchTransitRouter(ctx *middleware.Context, handler PatchTransitRouterHandler) *PatchTransitRouter {
	return &PatchTransitRouter{Context: ctx, Handler: handler}
}

/* PatchTransitRouter swagger:route PATCH /transit-routers/{id} Router patchTransitRouter

Update the supplied fields on a router

Update the supplied fields on a router. Requires admin access.

*/
type PatchTransitRouter struct {
	Context *middleware.Context
	Handler PatchTransitRouterHandler
}

func (o *PatchTransitRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchTransitRouterParams()
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

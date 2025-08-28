// This file is safe to edit. Once it exists it will not be overwritten

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

package rest_oidc_auth_api_server

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/openziti/edge-api/rest_oidc_auth_api_server/operations"
	"github.com/openziti/edge-api/rest_oidc_auth_api_server/operations/o_id_c_authentication"
)

//go:generate swagger generate server --target ../../edge-api --name ZitiEdgeOIDCAuth --spec ../oidc_auth.yml --model-package rest_model --server-package rest_oidc_auth_api_server --principal interface{} --exclude-main

func configureFlags(api *operations.ZitiEdgeOIDCAuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ZitiEdgeOIDCAuthAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.UrlformConsumer = runtime.DiscardConsumer

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("html producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()

	if api.OIDcAuthenticationAuthenticateCertHandler == nil {
		api.OIDcAuthenticationAuthenticateCertHandler = o_id_c_authentication.AuthenticateCertHandlerFunc(func(params o_id_c_authentication.AuthenticateCertParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.AuthenticateCert has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationAuthenticateExtJWTHandler == nil {
		api.OIDcAuthenticationAuthenticateExtJWTHandler = o_id_c_authentication.AuthenticateExtJWTHandlerFunc(func(params o_id_c_authentication.AuthenticateExtJWTParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.AuthenticateExtJWT has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationAuthenticatePasswordHandler == nil {
		api.OIDcAuthenticationAuthenticatePasswordHandler = o_id_c_authentication.AuthenticatePasswordHandlerFunc(func(params o_id_c_authentication.AuthenticatePasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.AuthenticatePassword has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationAuthenticateUsernameHandler == nil {
		api.OIDcAuthenticationAuthenticateUsernameHandler = o_id_c_authentication.AuthenticateUsernameHandlerFunc(func(params o_id_c_authentication.AuthenticateUsernameParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.AuthenticateUsername has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationCompleteTotpEnrollmentHandler == nil {
		api.OIDcAuthenticationCompleteTotpEnrollmentHandler = o_id_c_authentication.CompleteTotpEnrollmentHandlerFunc(func(params o_id_c_authentication.CompleteTotpEnrollmentParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.CompleteTotpEnrollment has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationDeleteTotpEnrollmentHandler == nil {
		api.OIDcAuthenticationDeleteTotpEnrollmentHandler = o_id_c_authentication.DeleteTotpEnrollmentHandlerFunc(func(params o_id_c_authentication.DeleteTotpEnrollmentParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.DeleteTotpEnrollment has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationListAuthQueriesHandler == nil {
		api.OIDcAuthenticationListAuthQueriesHandler = o_id_c_authentication.ListAuthQueriesHandlerFunc(func(params o_id_c_authentication.ListAuthQueriesParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.ListAuthQueries has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationShowCertAuthHandler == nil {
		api.OIDcAuthenticationShowCertAuthHandler = o_id_c_authentication.ShowCertAuthHandlerFunc(func(params o_id_c_authentication.ShowCertAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.ShowCertAuth has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationShowExtJWTAuthHandler == nil {
		api.OIDcAuthenticationShowExtJWTAuthHandler = o_id_c_authentication.ShowExtJWTAuthHandlerFunc(func(params o_id_c_authentication.ShowExtJWTAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.ShowExtJWTAuth has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationShowPasswordLoginHandler == nil {
		api.OIDcAuthenticationShowPasswordLoginHandler = o_id_c_authentication.ShowPasswordLoginHandlerFunc(func(params o_id_c_authentication.ShowPasswordLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.ShowPasswordLogin has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationShowUsernameLoginHandler == nil {
		api.OIDcAuthenticationShowUsernameLoginHandler = o_id_c_authentication.ShowUsernameLoginHandlerFunc(func(params o_id_c_authentication.ShowUsernameLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.ShowUsernameLogin has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationStartTotpEnrollmentHandler == nil {
		api.OIDcAuthenticationStartTotpEnrollmentHandler = o_id_c_authentication.StartTotpEnrollmentHandlerFunc(func(params o_id_c_authentication.StartTotpEnrollmentParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.StartTotpEnrollment has not yet been implemented")
		})
	}
	if api.OIDcAuthenticationVerifyTotpHandler == nil {
		api.OIDcAuthenticationVerifyTotpHandler = o_id_c_authentication.VerifyTotpHandlerFunc(func(params o_id_c_authentication.VerifyTotpParams) middleware.Responder {
			return middleware.NotImplemented("operation o_id_c_authentication.VerifyTotp has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

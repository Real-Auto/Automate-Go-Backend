// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Automate-Go-Backend/restapi/operations"
	"Automate-Go-Backend/restapi/operations/user"
)

//go:generate swagger generate server --target ..\..\Automate-Go-Backend --name AutomateAPI --spec ..\swagger.yml --principal interface{}

func configureFlags(api *operations.AutomateAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AutomateAPIAPI) http.Handler {
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

	api.JSONProducer = runtime.JSONProducer()

	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.UserLoginHandler == nil {
		api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		})
	}
	if api.ChangePasswordHandler == nil {
		api.ChangePasswordHandler = operations.ChangePasswordHandlerFunc(func(params operations.ChangePasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ChangePassword has not yet been implemented")
		})
	}
	if api.DeleteUserHandler == nil {
		api.DeleteUserHandler = operations.DeleteUserHandlerFunc(func(params operations.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteUser has not yet been implemented")
		})
	}
	if api.UserSignUpHandler == nil {
		api.UserSignUpHandler = user.SignUpHandlerFunc(func(params user.SignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation user.SignUp has not yet been implemented")
		})
	}
	if api.UpdateUserHandler == nil {
		api.UpdateUserHandler = operations.UpdateUserHandlerFunc(func(params operations.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateUser has not yet been implemented")
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

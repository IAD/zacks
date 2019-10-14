// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github/IAD/zacks/internal/app/server/gen/server/restapi/operations"
)

//go:generate swagger generate server --target ../../server --name Zacks --spec ../../../../../../api/swagger.yml --template-dir ../go-swagger-template/templates --exclude-main

func configureFlags(api *operations.ZacksAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ZacksAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Handler for GET /{ticker}
	api.GetTickerHandler = operations.GetTickerHandlerFunc(func(params *operations.GetTickerParams,
		getTickerOK operations.NewGetTickerOKFunc,
		getTickerNotFound operations.NewGetTickerNotFoundFunc,
		getTickerInternalServerError operations.NewGetTickerInternalServerErrorFunc,
	) middleware.Responder {
		return middleware.NotImplemented("operation .GetTicker has not yet been implemented")
	})
	// Handler for GET /{ticker}/history
	api.GetTickerHistoryHandler = operations.GetTickerHistoryHandlerFunc(func(params *operations.GetTickerHistoryParams,
		getTickerHistoryOK operations.NewGetTickerHistoryOKFunc,
		getTickerHistoryNotFound operations.NewGetTickerHistoryNotFoundFunc,
		getTickerHistoryInternalServerError operations.NewGetTickerHistoryInternalServerErrorFunc,
	) middleware.Responder {
		return middleware.NotImplemented("operation .GetTickerHistory has not yet been implemented")
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/version" {
			serveVersion(w, r)
			return
		}

		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusOK)
			return
		}

		cors.AllowAll().Handler(handler).ServeHTTP(w, r)
	})
}

func serveVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

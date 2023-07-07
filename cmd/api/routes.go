package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// ! general
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	// ! users
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	// ! tokens
	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.createActivationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationHandler)

	// ! movies
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.requirePermission("movies:write", app.requireActivatedUser(app.createMovieHandler)))
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requirePermission("movies:read", app.requireActivatedUser(app.showMovieHandler)))
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.requirePermission("movies:read", app.requireActivatedUser(app.listMoviesHandler)))
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requirePermission("movies:write", app.requireActivatedUser(app.updateMovieHandler)))
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requirePermission("movies:write", app.requireActivatedUser(app.deleteMovieHandler)))

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}

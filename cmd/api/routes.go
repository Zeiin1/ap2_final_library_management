package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/", app.showSnippet)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.ServeFiles("/mailer/static/*filepath", http.Dir("static"))

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

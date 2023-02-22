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
	router.HandlerFunc(http.MethodGet, "/", app.showIndexPage)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.ServeFiles("/static/*filepath", http.Dir("internal/mailer/static"))

	return app.recoverPanic(app.authenticate(router))
}

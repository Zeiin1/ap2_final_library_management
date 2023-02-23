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
	router.HandlerFunc(http.MethodPost, "/registration", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/registrationPage", app.registrationPage)
	router.HandlerFunc(http.MethodGet, "/loginPage", app.loginPage)
	router.HandlerFunc(http.MethodGet, "/profile/:id", app.profilePage)
	router.HandlerFunc(http.MethodPost, "/login", app.loginUser)
	router.ServeFiles("/static/*filepath", http.Dir("internal/web/static"))

	return app.recoverPanic(router)
}

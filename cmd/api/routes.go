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
	router.HandlerFunc(http.MethodGet, "/user_info/:id", app.backProfilePage)
	router.HandlerFunc(http.MethodPost, "/update/:id", app.updateUserInfo) //it is for update, i dont know but with
	//put method its not working
	router.HandlerFunc(http.MethodGet, "/delete/:id", app.deleteUser) //actually it must be delete mapping but
	//i find out that html version supports only post and get methods
	router.ServeFiles("/static/*filepath", http.Dir("internal/web/static"))

	return app.recoverPanic(router)
}

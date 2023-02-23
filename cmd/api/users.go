package main

import (
	"awesomeProject/internal/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.FormValue("name")
	surname := r.FormValue("surname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := data.User{
		Surname: surname,
		Name:    username,
		Email:   email,
	}

	err := user.Password.Set(password)
	if err != nil {
		http.Redirect(w, r, "/registrationPage", http.StatusSeeOther)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {

		http.Redirect(w, r, "/registrationPage", http.StatusSeeOther)
		return
	}

	if err != nil {
		http.Redirect(w, r, "/registrationPage", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
	return
}
func (app *application) registrationPage(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	ts, err := template.ParseFiles("./internal/web/templates/reg.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, user)
	if err != nil {
		log.Println(err.Error())
		return
	}

}

func (app *application) showIndexPage(w http.ResponseWriter, r *http.Request) {

	// Template
	ts, err := template.ParseFiles("./internal/web/templates/index.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
func (app *application) loginPage(w http.ResponseWriter, r *http.Request) {

	// Template
	ts, err := template.ParseFiles("./internal/web/templates/login.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	id, active := app.models.Users.Authenticate(email, password)
	if !active {
		http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/profile/%d", id), http.StatusSeeOther)
}
func (app *application) profilePage(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
	}

	user, err := app.models.Users.GetById(id)

	if err != nil {

		fmt.Println(err)
		http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
	}
	ts, err := template.ParseFiles("./internal/web/templates/profile.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, user)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
func (app *application) backProfilePage(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
	}

	user, err := app.models.Users.GetById(id)

	if err != nil {

		fmt.Println(err)
		http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
	}
	ts, err := template.ParseFiles("./internal/web/templates/backProfile.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, user)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
func (app *application) updateUserInfo(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	r.ParseForm()
	email := r.FormValue("email")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	password := r.FormValue("password")
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, fmt.Sprintf("/user_info/%d", id), http.StatusSeeOther)
	}
	user := data.User{
		ID:      id,
		Surname: surname,
		Email:   email,
		Name:    name,
	}
	err = user.Password.Set(password)
	if err != nil {
		fmt.Println(err)
		return
	}
	ok := app.models.Users.Update(&user)
	if !ok {
		fmt.Println("nothing changed")
		http.Redirect(w, r, fmt.Sprintf("/user_info/%d", id), http.StatusSeeOther)
	}

	if err != nil {

		fmt.Println(err)
		http.Redirect(w, r, fmt.Sprintf("/user_info/%d", id), http.StatusSeeOther)
	}
	http.Redirect(w, r, fmt.Sprintf("/user_info/%d", id), http.StatusSeeOther)

}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, fmt.Sprintf("/user_info/%d", id), http.StatusSeeOther)
	}
	ok := app.models.Users.Delete(id)
	if !ok {
		http.Redirect(w, r, fmt.Sprintf("/user_info/%d", id), http.StatusSeeOther)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

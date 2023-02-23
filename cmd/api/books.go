package main

import (
	"awesomeProject/internal/data"

	"html/template"
	"log"
	"net/http"
)

func (app *application) addBookPage(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./internal/web/templates/addBook.html")

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

func (app *application) addBookPageToDb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	author := r.FormValue("author")
	year := r.FormValue("year")
	var book = data.Book{
		BookName:     name,
		Author:       author,
		ReleasedYear: year,
	}
	ok := app.models.Books.InsertBook(book)
	if !ok {
		http.Redirect(w, r, "/library/show", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/library/show", http.StatusSeeOther)

}
func (app *application) showAllBooksFromDb(w http.ResponseWriter, r *http.Request) {

	books, err := app.models.Books.Latest()
	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, "could not process")
	}

	data := &templateData{Books: books}
	ts, err := template.ParseFiles("./internal/web/templates/booksPage.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

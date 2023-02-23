package main

import "awesomeProject/internal/data"

// like cmd/web/templates.go
type templateData struct {
	User  *data.User
	Users []*data.User
	Books []*data.Book
}

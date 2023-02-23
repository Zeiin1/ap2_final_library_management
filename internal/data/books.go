package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID           int64  `json:"id"`
	BookName     string `json:"book_name"`
	Author       string `json:"author"`
	ReleasedYear string `json:"released_year"`
}
type BookModel struct {
	DB *sql.DB
}

func (m BookModel) InsertBook(book Book) bool {
	query := `
INSERT INTO books (bookname, author, releasedyear)
VALUES ($1, $2, $3)
`
	args := []any{book.BookName, book.Author, book.ReleasedYear}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func (m *BookModel) Latest() ([]*Book, error) {

	stmt := `SELECT id, bookname, author, releasedyear FROM books
`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []*Book{}

	for rows.Next() {

		s := &Book{}

		err = rows.Scan(&s.ID, &s.BookName, &s.Author, &s.ReleasedYear)
		if err != nil {
			return nil, err
		}

		books = append(books, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

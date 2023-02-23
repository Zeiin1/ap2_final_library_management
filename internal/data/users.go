package data

import (
	"awesomeProject/internal/validator"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user User) error {
	query := `
INSERT INTO users (name, email, surname, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at`
	args := []any{user.Name, user.Email, user.Surname, user.Password.hash}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
func (m UserModel) UpdatePassword(user *User) error {
	query := `
UPDATE users
SET  password_hash = $1, activated = $2
WHERE id = $3`
	args := []any{user.Password.hash /*user.Activated,*/, user.ID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		print("password successfully changed ")
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
SELECT id, created_at, name, email, password_hash, activated, version
FROM users
WHERE email = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		/*	&user.Activated,
			&user.Version,*/
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
func (m UserModel) GetById(id int64) (*User, error) {
	query := `
SELECT id,  name, surname, email
FROM users
WHERE id = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,

		/*	&user.Activated,
			&user.Version,*/
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
func (m UserModel) Delete(id int64) bool {
	query := `
DELETE FROM users WHERE id = $1`
	args := []any{id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}
func (m *UserModel) Authenticate(email, password string) (int, bool) {

	var id int
	var hashedPassword []byte
	query := `
SELECT id, password_hash
FROM users
WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&id,
		&hashedPassword,
	)
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, false
		} else {
			fmt.Println(err)
			return 0, false
		}
	}
	return id, true
}

func (m UserModel) Update(user *User) bool {
	query := `
UPDATE users
SET name = $1, email = $2, password_hash = $3, surname = $4
WHERE id = $5 
RETURNING id`
	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Surname,
		user.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be provided")
}
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")
	ValidateEmail(v, user.Email)
	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}
	if user.Password.hash == nil {
		panic("missing password hash for user")
	}

}

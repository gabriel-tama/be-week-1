// models/user.go

package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint
	Username string
	Name     string
	Password string
}

type UserModel struct {
	db   *pgx.Conn
	salt int
}

func NewUserModel(db *pgx.Conn, salt int) *UserModel {
	return &UserModel{db: db, salt: salt}
}

func (um *UserModel) Create(username, name, password string) (*User, error) {

	var user_id int

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), um.salt)
	if err != nil {
		return nil, err
	}

	// Store the user in the database
	err = um.db.QueryRow(context.Background(), "INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING id",
		username, name, string(hashedPassword)).Scan(&user_id)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       uint(user_id),
		Username: username,
		Name:     name,
	}

	return user, nil
}

func (um *UserModel) FindExistByUsername(username string) (bool, error) {
	var exists bool
	err := um.db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false, errors.New("something went wrong")
	}
	return exists, nil
}

func (um *UserModel) FindUserByUsername(username string) (*User, error) {
	var storedPassword, name string
	var user_id int
	err := um.db.QueryRow(context.Background(), "SELECT id,name,password FROM users WHERE username = $1", username).Scan(&user_id, &name, &storedPassword)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:       uint(user_id),
		Username: username,
		Name:     name,
		Password: storedPassword,
	}
	return user, nil
}

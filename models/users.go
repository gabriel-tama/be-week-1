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
    db *pgx.Conn
}

func NewUserModel(db *pgx.Conn) *UserModel {
    return &UserModel{db: db}
}

func (um *UserModel) Create(username, name, password string) (*User,error) {
    // Check if the username already exists
    exists, err := um.FindExistByUsername(username)
    if err != nil {
        return nil,errors.New("something went wrong")
    }

    if exists {
        return nil,errors.New("username already exist")
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil,errors.New("something went wrong")
    }

    // Store the user in the database
    _, err = um.db.Exec(context.Background(), "INSERT INTO users (username, name, password) VALUES ($1, $2, $3)",
        username, name, string(hashedPassword))
    if err != nil {
        return nil,errors.New("something went wrong")
    }
        user := &User{
        Username: username,
        Name:     name,
    }

    return user,nil
}

func (um *UserModel) FindExistByUsername(username string) (bool,error){
    var exists bool
	err := um.db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false,errors.New("something went wrong")
	}
	return exists,nil
}

func (um *UserModel) FindUserByUsername(username string) (*User,error){
    var storedPassword,name string
	err := um.db.QueryRow(context.Background(), "SELECT name,password FROM users WHERE username = $1", username).Scan(&name,&storedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil,err
		}
		return nil,err
	}
    user := &User{
        Username: username,
        Name:     name,
        Password: storedPassword,
    }
    return user, nil
}





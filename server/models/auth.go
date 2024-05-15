package models

import (
	"errors"

	"github.com/falasefemi2/chat-app/db"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(email, password string) (*User, error) {
	var user User

	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Remove the password from the returned User struct for security
	user.Password = ""

	return &user, nil

}

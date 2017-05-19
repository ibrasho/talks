package model

import (
	"fmt"
	"github.com/pkg/errors"
)

type M interface {
	Id() int
	Save() error
}

type User struct {
	id       int
	Username string
	password string
	Email    string
}

func NewUser(id int, username, email, password string) *User {
	return &User{
		id:       id,
		Username: username,
		Email:    email,
		password: password,
	}
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Save() (err error) {
	return
}

func (u *User) SendEmail() error {
	if err := sendEmail(u.Email); err != nil {
		return err
	}

	return nil
}

func (u *User) String() string {
	return fmt.Sprintf("Email: %v", u.Email)
}

func sendEmail(email string) error {
	// do something
	if email == "" {
		errors.New("empty email")
	}

	return nil
}

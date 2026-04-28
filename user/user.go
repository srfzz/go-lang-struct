package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
	updatedAt time.Time
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt, u.updatedAt)
}

func (u *User) ClearuserDetails() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Firstname ,LastName and Birthdate are Required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

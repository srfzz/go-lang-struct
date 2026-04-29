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
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and Password Are required")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Sarfaraj",
			lastName:  "Ansari",
			birthDate: "23/12/1198",
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
	}, nil
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt, u.updatedAt)
}

func (u *User) ClearuserDetails() {
	u.firstName = ""
	u.lastName = ""
}

func (a *Admin) AdminOutputDetails() {
	fmt.Printf(a.email, a.password, a.User)
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

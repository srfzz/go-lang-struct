package main

import (
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

func main() {
	firstName := getUserData("Please Enter Your FirstName")
	lastName := getUserData("Please Enter Your LastName")
	birthDate := getUserData("Please Enter Your BirtDate (DD/MM/YY):")
	fmt.Println(firstName, lastName, birthDate)
	var appUser User
	appUser = User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}
	outputDetails(&appUser)
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt, appUser.updatedAt)
}

func outputDetails(u *User) {
	u.firstName = "Sarfaraj ansari"
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt, u.updatedAt)
}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value
}

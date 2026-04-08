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
	outputDetails(firstName, lastName, birthDate)
	appUser = User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
		updatedAt: time.Now().UTC(),
	}
}

func outputDetails(firstName, lastname, birthDate string) {
	fmt.Println(firstName, lastname, birthDate)
}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value
}

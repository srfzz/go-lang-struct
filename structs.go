package main

import (
	"fmt"

	"structexample.com/user"
)

func main() {
	firstName := getUserData("Please Enter Your FirstName")
	lastName := getUserData("Please Enter Your LastName")
	birthDate := getUserData("Please Enter Your BirtDate (DD/MM/YY):")

	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthDate)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	appUser.OutputDetails()
	appUser.ClearuserDetails()
	appUser.OutputDetails()
}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value
}

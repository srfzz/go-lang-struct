package main

import "fmt"

func main() {
	firstName := getUserData("Please Enter Your FirstName")
	lastName := getUserData("Please Enter Your LastName")
	birthDate := getUserData("Please Enter Your BirtDate (DD/MM/YY):")
	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value
}

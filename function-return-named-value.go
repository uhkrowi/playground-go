package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Moh"
	middleName = "Nurul"
	lastName = "Uhkrowi"

	return // equals to return firstName, middleName, lastName
	// return middleName, lastName, firstName <= use this if you want to switch returned variables
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}
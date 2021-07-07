package main

import "fmt"

func getFullName(firstName string, lastName string, fatherName string) string {
	return firstName + " " + lastName + " bin " + fatherName
}

func main() {
	fmt.Println(getFullName("Nurul", "Uhkrowi", "Ero Suntara"))
}
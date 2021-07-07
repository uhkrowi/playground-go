package main

import "fmt"

func getSeparatedName() (string, string, string, int) {
	return "Moh", "Nurul", "Uhkrowi", 26
}

func main() {
	var _, _, lastName, age = getSeparatedName()

	//fmt.Println(firstName)
	//fmt.Println(middleName)
	fmt.Println(lastName+"'s", "age:", age)
}
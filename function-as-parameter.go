package main

import "fmt"

func greeting(name string, filteredName func(string) string) string {
	return "Halo mang " + filteredName(name)
}

func filterWord(word string) string {
	// if word equals to some texts, change it to triple dots
	if word == "Anjing" {
		return "..."
	}

	return word
}

func main() {
	fmt.Println(greeting("Arul", filterWord))
	fmt.Println(greeting("Anjing", filterWord))
}

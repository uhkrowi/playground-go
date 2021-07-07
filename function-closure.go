package main

import "fmt"

func main() {
	mathGrade := 80
	physicGrade := 75

	// create closure with anonymous function
	average := func() int {
		return (mathGrade + physicGrade) / 2
	}

	fmt.Println(average)
}

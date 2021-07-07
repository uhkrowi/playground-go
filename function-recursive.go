package main

import "fmt"

func getFactorial(number int) int {
	if number > 1 {
		return number * getFactorial(number-1)
	}

	return number
}

func main() {
	fmt.Println(getFactorial(5))
}
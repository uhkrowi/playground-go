package main

import (
	"fmt"
)

/*
	IIFE stands for Immediately Invoked Function Expression
	You can run immediately the function you just created & assign the value to variables
	look at the parentheses after declaring the anonymous function
*/

func main() {
	grades := []int{1, 2, 3, 4, 5}

	min, max := func() (int, int) {
		min, max := 0, 0

		for _, v := range grades {
			switch {
			case v > max:
				max = v
			case v < min:
				min = v
			}
		}

		return min, max
	}()

	fmt.Printf("data: %v \nmin: %v \nmax: %v", grades, min, max)
}

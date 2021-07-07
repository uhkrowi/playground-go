package main

import "fmt"

/*
	Defer function is function that runs after any function run, whether the invoker function is success or error
*/

func logging() {
	fmt.Println("log: function has been running")
}

func runApp() {
	defer logging()

	fmt.Println("application start running")
}

func main() {
	runApp()
}


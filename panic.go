package main

import "fmt"

/*
	Panic is useful when you want to throw a stopper message then exit the application
*/

func tryPanic() {
	panic("panik ga?! panik ga?! ya panik lah masa ngga!")
}

func main() {
	tryPanic()
	fmt.Println("perhatiin, ini keprint gak")
}

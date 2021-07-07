package main

import "fmt"

/*
	Recover is used to avoid stopping & exit program when panic happens
*/

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured: ", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func tryPanicAgain() {
	defer catch()

	panic("ini komentar paniknya")
}

func main() {
	tryPanicAgain()
	fmt.Println("perhatiin, ini keprint gak")
}

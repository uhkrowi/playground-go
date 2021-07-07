package main

import "fmt"

func sayHello(name string) string {
	return "Halo mang " + name
}

func main() {
	greeting := sayHello

	fmt.Println(greeting("Arul"))
}
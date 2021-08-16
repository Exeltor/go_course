package main

import "fmt"

// Use the "defer" keyword to show that a deferred func runs after the func containing it exits

func main() {
	defer printSomething()
	fmt.Println("I was called second")
	fmt.Println("I was called third")
}

func printSomething() {
	fmt.Println("I was called first but I will run when my parent function exits")
}

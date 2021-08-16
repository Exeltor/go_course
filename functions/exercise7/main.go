package main

import "fmt"

// assign a func to a variable and call it

func main() {
	f := func() {
		fmt.Println("I am a function assigned to a variable")
	}

	f()
}

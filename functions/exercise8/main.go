package main

import "fmt"

// create a func which returns a func
// assign the reteured func to a variable
// call the returned func

func main() {
	f := returning()

	f()
}

func returning() func() {
	return func() {
		fmt.Println("I am a returned func")
	}
}

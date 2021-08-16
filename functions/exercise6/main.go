package main

import "fmt"

// build and use an anonymous function

func main() {
	func() {
		fmt.Println("I am an anonymous func")
	}()
}

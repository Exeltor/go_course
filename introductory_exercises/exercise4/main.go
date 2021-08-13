package main

import "fmt"

// Declare custom type, create global variable, print value, print type, assign value and print

type custom int

var x custom

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

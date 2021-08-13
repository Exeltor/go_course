package main

import "fmt"

// Declare custom type, create global variable, print value, print type, assign value and print

type custom int

var x custom
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
}

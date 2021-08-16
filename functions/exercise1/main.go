package main

import "fmt"

// create function with identifier "foo" that returns an int
// create function with identifier "bar" that returns an int and a string
// call both funcs
// print out result

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x, y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "Hello world"
}

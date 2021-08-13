package main

import "fmt"

// Printing declared but unassigned global variables

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
}

package main

import "fmt"

// Declaration and assignment of global variables
// Declare and assign "s" with Sprintf

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

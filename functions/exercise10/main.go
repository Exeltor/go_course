package main

import "fmt"

// create a func which encloses the scope of a variable

func main() {
	x := incrementToTen()

	fmt.Println(x)
}

func incrementToTen() int {
	x := 0
	for x <= 10 {
		x++
	}

	return x
}

package main

import "fmt"

// write expressions with comparators and assign the results to variables

func main() {
	x := 10
	y := 20

	equal := x == y
	less_equal := x <= y
	more_equal := x >= y
	not_equal := x != y
	less := x < y
	more := x > y

	fmt.Println("equal:", equal)
	fmt.Println("less or equal:", less_equal)
	fmt.Println("more or equal:", more_equal)
	fmt.Println("not equal:", not_equal)
	fmt.Println("less than:", less)
	fmt.Println("more than:", more)
}

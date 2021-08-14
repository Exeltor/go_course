package main

import "fmt"

// using a composite literal, create an array which holds 5 ints and assign values to each position
// range over the array and print the values
// using format printing, print out the TYPE of the array

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}

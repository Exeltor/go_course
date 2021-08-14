package main

import "fmt"

// using a composite literal, create an slice and assign 10 values
// range over the array and print the values
// using format printing, print out the TYPE of the array

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}

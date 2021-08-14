package main

import "fmt"

// start with slice {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// append 52 to the slice
// print slice
// in ONE STATEMENT append values 53, 54, 55
// print slice
// append this slice to the slice: {56, 57, 58, 59, 60}
// print out slice

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	x = append(x, []int{56, 57, 58, 59, 60}...)
	fmt.Println(x)
}

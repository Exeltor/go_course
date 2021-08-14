package main

import "fmt"

// slice the slice {42, 43, 44, 45, 46, 47, 48, 49, 50, 51} and obtain the following slices
// [42, 43, 44, 45, 46]
// [47, 48, 49, 50, 51]
// [44, 45, 46, 47, 48]
// [43, 44, 45, 46, 47]

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	first := x[:5]
	second := x[5:]
	third := x[2:7]
	fourth := x[1:6]

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
}

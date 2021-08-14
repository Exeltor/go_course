package main

import "fmt"

// start with slice {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// use append and slicing to get values which should print [42, 43, 44, 48, 49, 40, 51]

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

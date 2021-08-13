package main

import "fmt"

// using 'iota' create 4 constants for the last 4 years
// print the constants

const (
	a = 2021 + iota
	b = 2021 + iota
	c = 2021 + iota
	d = 2021 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}

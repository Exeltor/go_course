package main

import "fmt"

// create a for loop in the format "for {}"

func main() {
	x := 0

	for {
		if x > 22 {
			break
		}
		fmt.Println(x)
		x++
	}
}

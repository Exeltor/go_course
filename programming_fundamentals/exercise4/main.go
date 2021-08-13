package main

import "fmt"

// assign int to variable
// print int in decimal, binary and hex
// shift bits of int 1 position to the left and assign to variable
// print that variable in decimal, binary and hex

func main() {
	x := 10
	fmt.Printf("%d\t%b\t%x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%x\n", y, y, y)
}

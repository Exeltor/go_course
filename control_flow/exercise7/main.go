package main

import "fmt"

// if, else if and else

func main() {
	x := "Alexey Zhelezov"

	if x == "Alexey Zhelezov" {
		fmt.Println(x)
	} else if x == "AAA" {
		fmt.Println("else if", x)
	} else {
		fmt.Println("else", x)
	}
}

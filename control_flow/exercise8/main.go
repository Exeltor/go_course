package main

import "fmt"

// switch statement without expression

func main() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

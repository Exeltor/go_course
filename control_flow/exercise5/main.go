package main

import "fmt"

// print modulus dividing by 4 for each number between 10 and 100

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

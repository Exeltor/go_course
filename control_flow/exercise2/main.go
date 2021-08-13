package main

import "fmt"

// print unicodes of uppercase alphabet 3 times from ascii

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		for i := 65; i <= 90; i++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}

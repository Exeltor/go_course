package main

import "fmt"

// create and use an anonymous struct

func main() {
	s := struct {
		name    string
		surname string
	}{
		name:    "James",
		surname: "Bond",
	}

	fmt.Println(s)
}

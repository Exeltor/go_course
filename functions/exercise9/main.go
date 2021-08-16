package main

import "fmt"

// pass a func into a func as an argument (callback)

func main() {
	someFunc(func() {
		fmt.Println("I am a callback function")
	})
}

func someFunc(function func()) {
	function()
}

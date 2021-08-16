package main

import "fmt"

// create user defined struct with
// identifier person
// fields first, last, age

// attach methos to type person with
// identifier speak
// the method should have the person say their name and age

// create value of type person
// call method from the value of type person

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Alexey",
		last:  "Zhelezov",
		age:   22,
	}

	p.speak()
}

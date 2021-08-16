package main

import "fmt"

// create a person struct
// create func called "changeMe" with *person as parameter
// change the value stored at the *person address

// create value of type person
// print out value
// in func main call changeMe
// in func main print out the value

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "Alexey",
		last:  "Zhelezov",
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}

func changeMe(p *person) {
	p.first = "changed"
}

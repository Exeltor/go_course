package main

import "fmt"

// create own type "person" with underlying type struct
// store following data:
// first name
// last name
// favorite ice cream flavors

type person struct {
	firstName        string
	lastName         string
	favoriteFlavours []string
}

func main() {
	p1 := person{
		firstName:        "Person",
		lastName:         "One",
		favoriteFlavours: []string{"Vanilla, Strawberry"},
	}
	p2 := person{
		firstName:        "Person",
		lastName:         "Two",
		favoriteFlavours: []string{"Chocolate, Cream"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for _, v := range p1.favoriteFlavours {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for _, v := range p2.favoriteFlavours {
		fmt.Println(v)
	}
}

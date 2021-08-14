package main

import "fmt"

// create own type "person" with underlying type struct
// store following data:
// first name
// last name
// favorite ice cream flavors
// store persons in a map
// iterate over map and print values

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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

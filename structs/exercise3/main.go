package main

import "fmt"

// create new type "vehicle" with underlying type stuct and fields doors and color
// create two new types "truck" and "sedan" with undelying type struct
// embed "vehicle" in both "truck" and "sedan"
// give truck field "fourWheel" with type bool
// give sedan field "luxury" with type bool
// using composite literal, create type truck and assign values
// using composite literal, create type sedan and assign values
// print each of these values
// print single field from each of these values

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "Black",
		},
		luxury: true,
	}

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "White",
		},
		fourWheel: true,
	}

	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(s.doors)
	fmt.Println(t.fourWheel)
}

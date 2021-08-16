package main

import "fmt"

// create func with identifier foo that
// takes in variadic parameter of type int
// pass value of type []int into func (unfurling)
// return sum of all values

// create func with identifier bar that
// takes parameter of type []int
// return sum of all values passed

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(nums...))
	fmt.Println(bar(nums))
}

func foo(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func bar(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

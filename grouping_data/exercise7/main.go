package main

import "fmt"

// create a slice of a slice of string
// store the following data in the multidimentional array
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.

func main() {
	array1 := []string{"James", "Bond", "Shaken, not stirred"}
	array2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	jointArray := [][]string{array1, array2}
	fmt.Println(jointArray)

	for _, array := range jointArray {
		for _, value := range array {
			fmt.Println(value)
		}
	}
}

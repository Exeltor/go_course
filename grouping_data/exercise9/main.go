package main

import "fmt"

// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

func main() {
	x := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range x {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}

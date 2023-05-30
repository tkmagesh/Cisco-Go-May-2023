package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 4, "pencil": 2, "marker": 1}
	/*
		var productRanks map[string]int = map[string]int{
			"pen":    4,
			"pencil": 2,
			"marker": 1,
		}
	*/

	//
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 2
	*/

	productRanks := map[string]int{
		"pen":    4,
		"pencil": 2,
		"marker": 1,
	}
	fmt.Println(productRanks)

	fmt.Println("Accessing the values")
	fmt.Println("productRanks[\"pen\"] =", productRanks["pen"])

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q]=%d\n", key, val)
	}

	fmt.Println("checking for the existence of a key")
	nonExistingKey := "notepad"
	if val, exists := productRanks[nonExistingKey]; exists {
		fmt.Printf("productRanks[%q] = %d\n", nonExistingKey, val)
	} else {
		fmt.Printf("Key - %q does not exists\n", nonExistingKey)
	}

	fmt.Println("Removing an item")
	delete(productRanks, "pencil")
	fmt.Println(productRanks)

	fmt.Println("Adding a new pair")
	productRanks["scribble-pad"] = 7
	fmt.Println(productRanks)
}

package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	//type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	//using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// Iterating using indexer
	fmt.Println("Iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	// Iterating using range
	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		x := [3]int{20, 10, 30}
		y := [3]int{20, 10, 30}
		fmt.Println(x == y)
	*/
	x := [3]int{20, 10, 30}
	y := x
	y[0] = 100
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	fmt.Println("before sorting, nos =", nos)
	sort(&nos)
	fmt.Println("after sorting, nos =", nos)
}

func sort(x *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}

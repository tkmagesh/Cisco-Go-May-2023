package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}

	//type inference
	// var nos = []int{3, 1, 4, 2, 5}

	//using :=
	nos := []int{3, 1, 4, 2, 5}
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

	//appending new values
	fmt.Println("appending new values")
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	/*
		x := []int{20, 10, 30}
		y := []int{20, 10, 30}
		fmt.Println(x == y)
	*/

	x := []int{20, 10, 30}
	y := x
	y[0] = 100
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	fmt.Println("before sorting, nos =", nos)
	sort(nos)
	fmt.Println("after sorting, nos =", nos)

	fmt.Println("Slicing")
	fmt.Println("nos[2:5] =", nos[2:5])
	fmt.Println("nos[:5] =", nos[:5])
	fmt.Println("nos[5:] =", nos[5:])
}

func sort(x []int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}

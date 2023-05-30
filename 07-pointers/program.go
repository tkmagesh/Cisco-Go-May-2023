package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no
	fmt.Println(noPtr)

	//dereferencing ( from address (pointer) -> value )
	var x int
	x = *noPtr
	fmt.Println(x)

	fmt.Println("Before incrementing, x =", x)
	increment(&x)
	fmt.Println("After incrementing, x =", x)

	n1, n2 := 100, 200
	swap(&n1, &n2)
	fmt.Println(n1, n2)

}

func increment(val *int) {
	*val++
}

func swap(x, y *int) /* do NOT return anything */ {
	*x, *y = *y, *x
}

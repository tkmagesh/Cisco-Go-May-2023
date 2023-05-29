package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)

	var n1, n2 int
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1, n2)
}

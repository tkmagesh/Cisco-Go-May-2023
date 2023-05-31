package main

import "fmt"

type Employee struct {
	Id int
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Occaecat ipsum quis nostrud sunt irure sint laborum nulla in mollit. Dolor aliquip non consectetur eiusmod aute est cillum mollit aliquip cillum. Irure do ullamco consectetur minim enim ipsum labore sint tempor exercitation mollit eiusmod anim qui. Consequat ea ad commodo voluptate. Minim tempor aliqua elit adipisicing id magna aute laborum incididunt ad exercitation magna."
	x = true
	x = 19.99
	fmt.Println(x)

	/* type assertion - 1 */
	x = 200
	// y := x + 300 (will not compile)
	// y := x.(int) + 300
	if val, ok := x.(int); ok {
		y := val + 300
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	/* type assertion - 2 */
	x = 100
	// x = "Occaecat ipsum quis nostrud sunt irure sint laborum nulla in mollit. Dolor aliquip non consectetur eiusmod aute est cillum mollit aliquip cillum. Irure do ullamco consectetur minim enim ipsum labore sint tempor exercitation mollit eiusmod anim qui. Consequat ea ad commodo voluptate. Minim tempor aliqua elit adipisicing id magna aute laborum incididunt ad exercitation magna."
	// x = true
	// x = 19.99
	x = Employee{100}
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 300 =", val+300)
	case string:
		fmt.Println("x is a string, length of x =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x * 99% =", val*(99.0/100.0))
	case Employee:
		fmt.Println("x is an employee with id :", val.Id)
	case struct{}:
		fmt.Println("x is an empty object")
	default:
		fmt.Println("x is not a known type")
	}

}

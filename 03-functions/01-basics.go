package main

import "fmt"

func main() {
	/* no parameters, no return values */
	sayHi()

	/* 1 parameter, no return values */
	greet("Magesh")

	/* 1 parameter, 1 return value */
	print(getGreetMsg("Suresh"))

	/* 2 parameters, 1 return value */
	fmt.Println("add(100,200) =", add(100, 200))

	/* more than 1 return values */
	// fmt.Println(divide(100, 7))

	/*
		q, r := divide(100, 7)
		fmt.Printf("divide(100,7), quotient = %d and remainder = %d\n", q, r)
	*/

	// q := divide(100, 7) // Wrong. # of variables should match the # of return values
	// q, r := divide(100, 7) // Wrong. r is not used
	q, _ := divide(100, 7)
	fmt.Printf("divide(100,7), quotient = %d\n", q)
}

func sayHi() {
	fmt.Println("Hi")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// using "named return results" - preferred when a function returns more than 1 result
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}

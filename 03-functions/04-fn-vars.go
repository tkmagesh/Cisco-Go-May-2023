package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}
	var greetMsg = getGreetMsg("Suresh")
	print(greetMsg)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	var q, r = divide(100, 7)
	fmt.Printf("divide(100,7), quotient = %d and remainder = %d\n", q, r)

	/*
		//anonymous version of 'greet'


		//anonymous version of "getGreetMsg"


		//anonymous version of 'divide'
		var q, r = func(x, y int) (quotient, remainder int) {
			quotient, remainder = x/y, x%y
			return
		}(100, 7)
		fmt.Printf("divide(100,7), quotient = %d and remainder = %d\n", q, r)
	*/
}

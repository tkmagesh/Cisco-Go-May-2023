package main

import "fmt"

func main() {
	//anonymous version of 'sayHi'
	func() {
		fmt.Println("Hi")
	}()

	//anonymous version of 'greet'
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	//anonymous version of "getGreetMsg"
	var greetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}("Suresh")
	print(greetMsg)

	//anonymous version of 'divide'
	var q, r = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("divide(100,7), quotient = %d and remainder = %d\n", q, r)
}

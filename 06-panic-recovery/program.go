package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		fmt.Println("deffered - main")
		if err := recover(); err != nil {
			fmt.Println("something went wrong: ", err)
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("divide(100,%d), quotient = %d and remainder = %d\n", divisor, q, r)
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient, remainder = x/y, x%y
	return
}

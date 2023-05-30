package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == DivideByZeroError {
		fmt.Println("please do not divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong. err = ", err)
		return
	}
	fmt.Printf("divide(100,%d), quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(x, y int) (int, int, error) {

	// if y == 0 {
	// 	err := errors.New("divisor cannot be zero")
	// 	return 0, 0, err
	// }

	if y == 0 {
		err := DivideByZeroError
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}

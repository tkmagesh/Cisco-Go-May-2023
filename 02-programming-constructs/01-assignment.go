/*
Write  a program that checks if 97 is a prime number and print the result
*/

package main

import "fmt"

func main() {
	num := 97
	for i := 2; i <= (num / 2); i++ {
		if num%i == 0 {
			fmt.Printf("%d is not a prime\n", num)
			return
		}
	}
	fmt.Printf("%d is a prime\n", num)
}

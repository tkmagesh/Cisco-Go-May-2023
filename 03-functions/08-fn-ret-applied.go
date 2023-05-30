package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/
	/*
		logAdd := getLoggedOperation(add)
		logAdd(100, 200)

		logSubtract := getLoggedOperation(subtract)
		logSubtract(100, 200)
	*/

	/*
		profileAdd := getProfiledOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfiledOperation(subtract)
		profileSubtract(100, 200)
	*/

	logAdd := getLoggedOperation(add)
	profileLogAdd := getProfiledOperation(logAdd)
	profileLogAdd(100, 200)
}

func getLoggedOperation(operation func(int, int)) func(int, int) {
	loggedOperation := func(x, y int) {
		log.Println("Operation started")
		operation(x, y)
		log.Println("Operation completed")
	}
	return loggedOperation
}

func getProfiledOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Printf("Elapsed : %v\n", elapsed)
	}
}

/*
func logOperation(x, y int, operation func(int, int)) {
	log.Println("Operation started")
	operation(x, y)
	log.Println("Operation completed")
}
*/

func add(x, y int) {
	fmt.Printf("Add Result = %d\n", x+y)
}

func subtract(x, y int) {
	fmt.Printf("Subtract Result = %d\n", x-y)
}

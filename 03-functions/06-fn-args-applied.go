package main

import (
	"fmt"
	"log"
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
	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func logOperation(x, y int, operation func(int, int)) {
	log.Println("Operation started")
	operation(x, y)
	log.Println("Operation completed")
}

func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}

func add(x, y int) {
	fmt.Printf("Add Result = %d\n", x+y)
}

func subtract(x, y int) {
	fmt.Printf("Subtract Result = %d\n", x-y)
}

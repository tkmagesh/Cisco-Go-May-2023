package main

import "fmt"

func main() {
	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}

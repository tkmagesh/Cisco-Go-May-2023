package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go generateNos(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func generateNos(ch chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i
	}
}

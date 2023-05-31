package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int)
	go generateNos(ch)
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func generateNos(ch chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

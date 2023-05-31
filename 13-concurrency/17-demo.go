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

	ch := generateNos()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func generateNos() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

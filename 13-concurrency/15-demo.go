package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go generateNos(ch)
	for {
		// fmt.Println(<-ch)

		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
		} else {
			break
		}

	}
	fmt.Println("Done")
}

func generateNos(ch chan int) {
	for i := 1; i <= 20; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1) // increment the counter by any arbitrary value
		go fn(i, wg)
	}
	wg.Wait() // block until the wg counter becomes 0
}

func fn(id int, wg *sync.WaitGroup) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(3 * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the counter ALWAYS by 1
}

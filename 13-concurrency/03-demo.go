package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the counter by any arbitrary value
	go f1(wg)
	f2()
	wg.Wait() // block until the wg counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter ALWAYS by 1
}

func f2() {
	fmt.Println("f2 invoked")
}

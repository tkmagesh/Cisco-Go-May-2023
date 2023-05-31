package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the counter by any arbitrary value
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		f1()
	}(wg)
	f2()
	wg.Wait() // block until the wg counter becomes 0
}

//3rd party api
func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}

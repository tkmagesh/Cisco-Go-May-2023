package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the counter by any arbitrary value
	go fn(wg)
	wg.Wait() // block until the wg counter becomes 0
}

func fn(wg *sync.WaitGroup) {
	for {

	}
}

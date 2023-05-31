package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(20)
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines")
	flag.Parse()
	fmt.Printf("Starting %d goroutines. Hit ENTER to start...\n", count)
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	for i := 1; i <= count; i++ {
		wg.Add(1) // increment the counter by any arbitrary value
		go fn(i, wg)
	}
	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("Done")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the counter ALWAYS by 1
}

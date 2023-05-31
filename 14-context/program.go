package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	valCtx := context.WithValue(cancelCtx, "key-1", "val-1")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	ch := genFib(valCtx)
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(ctx context.Context) chan int {
	fmt.Println("context value :", ctx.Value("key-1"))

	ch := make(chan int)
	go func() {
		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		go genNos(timeoutCtx)

		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
		fmt.Println("genFib exiting")
	}()
	return ch
}

func genNos(ctx context.Context) {
	counter := 1
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("counter : ", counter)
			counter++
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("genNos exiting")
}

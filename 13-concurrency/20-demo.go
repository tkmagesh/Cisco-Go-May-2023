/* handling more than one channel */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib() chan int {
	ch := make(chan int)
	/*
		timeout := make(chan time.Time)
		go func() {
			time.Sleep(5 * time.Second)
			timeout <- time.Now()
		}()
	*/
	// timeout := After(5 * time.Second)
	timeout := time.After(5 * time.Second)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-timeout:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

func After(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}

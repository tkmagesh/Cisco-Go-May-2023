/* handling more than one channel */

package main

import (
	"fmt"
	"time"
)

func main() {

	stopCh := make(chan struct{})
	ch := genFib(stopCh)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// close(stopCh)
		stopCh <- struct{}{}
	}()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(stopCh <-chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
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

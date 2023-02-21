package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()

			}

		}()

	}
	time.Sleep(5 * time.Second)
	fmt.Println(x)

}

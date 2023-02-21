package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var x int64
	group := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {

		go func() {
			group.Add(1)
			for i := 0; i < 100; i++ {

				atomic.AddInt64(&x, 1)

			}
			group.Done()

		}()

	}
	group.Wait()
	fmt.Println(x)

}

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()

	}

	totalCPU := runtime.NumCPU()
	fmt.Println("CPU", totalCPU)
	runtime.GOMAXPROCS(20)
	totalThred := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThred)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("goroutine", totalGoroutine)
	group.Wait()

}

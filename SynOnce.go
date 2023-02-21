package main

import (
	"fmt"
	"sync"
)

var counter = 0

// mutex := sync.Mutex{}

func main() {

	group := sync.WaitGroup{}
	once := sync.Once{}
	var mutex = sync.Mutex{}

	for i := 0; i < 100; i++ {

		go func() {

			group.Add(1)
			mutex.Lock()
			once.Do(Once)
			fmt.Println(counter)
			mutex.Unlock()
			group.Done()

		}()

	}

	group.Wait()
	fmt.Println(counter)

}

func Once() {
	// mutex.Lock()
	counter++
	// mutex.Unlock()

}

package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("hello")
	time.Sleep(1 * time.Second)

}

func main() {

	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronus(group)
	}

	group.Wait()
	fmt.Println("selesai")

}

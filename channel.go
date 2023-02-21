package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := sync.WaitGroup{}

	var channel = make(chan string)
	defer close(channel)
	group.Add(1)
	go func() {

		time.Sleep(2 * time.Second)
		channel <- "Teguh"
		fmt.Println("selesai")
		group.Done()

	}()

	// time.Sleep(5 * time.Second)

	data := <-channel
	fmt.Println(data)
	group.Wait()

}

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 3)
	defer close(channel)
	channel <- "Teguh"
	channel <- "Teguh1"
	channel <- "Teguh2"

	go func() {
		channel <- "Teguh"
		channel <- "Teguh1"
		channel <- "Teguh2"

	}()
	// channel <- "Teguh"
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()
	time.Sleep(2 * time.Second)
	fmt.Println("selesai")

}

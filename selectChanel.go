package main

import (
	"fmt"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hallo teguh"

}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	count := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("data 1", data)
			count++
		case data := <-channel2:
			fmt.Println("data 2", data)
			count++

		default:
			fmt.Println("Menunggu data")
		}
		if count == 2 {
			break
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	fmt.Println("aku")
	// go func() {
	// 	channel <- "susu"
	// }()

	// data1 := <-channel
	// fmt.Println(data1)
	// time.Sleep(1 * time.Nanosecond)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hallo teguh"

}
func GiveMeResponse0() {

	fmt.Println("hello world")

}

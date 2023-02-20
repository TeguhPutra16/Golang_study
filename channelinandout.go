package main

import (
	"fmt"
	"time"
)

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Teguh Putra"

}
func OnlyOut(channle <-chan string) {
	data := <-channle
	fmt.Println(data)

}

func main() {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
	close(channel)

}

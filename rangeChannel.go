package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan string)

	go func() {

		for i := 0; i < 10; i++ {

			channel <- strconv.Itoa(i)

		}
		close(channel)

	}()

	for data := range channel {
		fmt.Println("menerima data", data)

	}

	// time.Sleep(3 * time.Second)
	fmt.Println("selesai")

}

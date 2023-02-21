package main

import (
	"fmt"
	"time"
)

func main() {
	channel := time.Tick(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)

	}()

	for data := range channel {
		fmt.Println(data)

	}
}

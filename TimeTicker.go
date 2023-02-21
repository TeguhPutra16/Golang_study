package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for data := range ticker.C {
		fmt.Println(data)

	}
}

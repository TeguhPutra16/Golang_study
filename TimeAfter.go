package main

import (
	"fmt"
	"time"
)

func main() {

	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello world!")

	for counter := 0; counter < 10; counter++ {
		fmt.Println(<-channel)
	}

	fmt.Println("End")
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

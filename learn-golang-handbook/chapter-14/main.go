package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// Start a goroutine to receive from the channel
	go func() {
		v := <-ch
		fmt.Println("Received:", v)
	}()

	// Send to the channel
	ch <- 69

	// Give the goroutine a chance to run.
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done")
}

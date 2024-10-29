package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

}

func exampleTimeout() {
	// Initialize empty context struct.
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(9 * time.Second)
		close(done)
	}()

  select {
  case <- done:
    fmt.Println("Called the API")

  case <- ctxWithTimeout.Done():
    fmt.Println("Timeout expired...", ctxWithTimeout.Err())
    // Do something to handle.
  }
}

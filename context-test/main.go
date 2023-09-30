package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Received cancellation signal. Stopping.")
			return
		default:
			// Simulate some work
			fmt.Println("Worker: Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the context is canceled when we're done

	var wg sync.WaitGroup

	// Start two worker goroutines
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	// Simulate some main work
	time.Sleep(3 * time.Second)

	// Cancel the context to signal workers to stop
	fmt.Println("Main: Cancelling workers...")
	cancel()

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("Main: All workers have finished.")
}

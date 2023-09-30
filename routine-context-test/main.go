package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(150 * time.Millisecond) // Simulate some work
		go printNumbers()
		go printNumbers()
	}
}

func main() {
	go printNumbers() // Start a Goroutine to print numbers concurrently
	go printLetters() // Start a Goroutine to print letters concurrently

	// Sleep for a while to allow Goroutines to execute
	time.Sleep(3 * time.Second)

	fmt.Println("Main function exiting.")
}

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	doSomethingDangerous()
	fmt.Println("Program continues running after potential panic.")
}

func doSomethingDangerous() {
	// Simulate a panic condition.
	panic("A panic occurred!")
}

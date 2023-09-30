package main

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Suffix = "  Folder creating"
	s.Color("red", "bold")      // Set the spinner color to a bold red// Build our new spinner
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
}

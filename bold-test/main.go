// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// ANSI escape code for red text
// 	red := "\033[31m"
// 	// ANSI escape code for bold text
// 	bold := "\033[1m"
// 	// ANSI escape code to reset text style to normal
// 	reset := "\033[0m"

// 	// Text with multiple font styles
// 	text := bold + "This is bold " + red + "and red" + reset

//		fmt.Println(text)
//	}
// package main

// import "os"

// func main() {
// 	// Set text color to red
// 	// redText := "\x1b[31m" + "This is red text"

// 	// // Reset text styling
// 	// resetText := "\x1b[0m" + "This text has default styling"

// 	// // Print the styled and reset text
// 	// fmt.Println(redText)
// 	// fmt.Println(resetText)

// 	err := os.Mkdir("myfolder", 0755)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create a directory and any necessary parent directories
// 	err = os.MkdirAll("myfolder/subfolder", 0755)
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	projectName := "myproject" // Replace with your project name
	projectRoot := filepath.Join(".", projectName)

	dirs := []string{
		"cmd/myapp",
		"internal/app/handler",
		"internal/app/model",
		"internal/app/repository",
		"internal/app/usecase",
		"internal/config",
		"internal/database/migrations",
		"pkg/logger",
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(projectRoot, dir)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dirPath, err)
			return
		}
		fmt.Printf("Created directory: %s\n", dirPath)
	}

	fmt.Println("Directory structure created successfully!")
}

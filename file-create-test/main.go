package main

import (
	"fmt"
	"os"
)

func main() {
	folderPath := "file-create-test"
	fileName := "example.txt"

	// err := os.MkdirAll(folderPath, os.ModePerm)
	// if err != nil {
	// 	fmt.Println("Error creating folder:", err)
	// 	return
	// }

	filePath := folderPath + "/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	data := "This is some example data that will be written to the file."
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	fmt.Println("File created successfully at:", filePath)
}

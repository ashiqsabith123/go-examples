package main

import "fmt"

type Books struct {
	title string
}

func main() {
	var Book1 Books
	var Book2 *Books
	Book1.title = "Go Programming"
	Book2 = &Book1
	printBook(&Book2)
}

func printBook(book **Books) {
	fmt.Printf("Book title : %s\n", (**book.title))
}

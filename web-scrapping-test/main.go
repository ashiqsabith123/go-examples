package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Specify the URL of the website you want to scrape
	url := "https://www.facebook.com/"

	// Make an HTTP GET request to the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Getting the heading of facebook login page
	doc.Find("h2._8eso").Each(func(index int, element *goquery.Selection) {
		// Print the text of the first <h1> element found on the pageem

		fmt.Println(element.Text())

	})

	// Find and extract links from the HTML document
	// doc.Find("a").Each(func(index int, element *goquery.Selection) {
	// 	// Get the href attribute of each <a> element
	// 	link, _ := element.Attr("href")
	// 	fmt.Println(link)
	// })
}

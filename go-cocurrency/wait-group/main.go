package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var Details = map[int][]string{
	1: {"Ashiq", "Kottayam"},
	2: {"Sabith", "Iduki"},
	3: {"Sanika", "Wayanad"},
}

var chanel = make(chan string)
var wg sync.WaitGroup

func main() {
	// now := time.Now()

	// // Initialize the channel

	// wg.Add(1)

	// go getName(2)

	// wg.Wait()

	// close(chanel)

	// // for name := range chanel {
	// // 	fmt.Println(name)
	// // }

	// fmt.Println(<-chanel)

	// fmt.Println(time.Since(now))

	go fmt.Println("hello")

	time.Sleep(2 * time.Second)
	fmt.Println("hhh")

	log.Printf("jjj")

	time.Now().Add()

	time.Now().Sub()

	log.Fatal()
}

func getName(id int) {
	defer wg.Done()
	for key, value := range Details {
		if key == id {

			time.Sleep(2 * time.Millisecond)
			chanel <- value[0]
		}
	}
}

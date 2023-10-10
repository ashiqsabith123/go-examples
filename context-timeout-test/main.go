package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

var details = map[int]string{
	1: "Ashiq",
	2: "sabith",
	3: "sanika",
}

func main() {

	val, err := fetch(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)

}

type Resp struct {
	value string
	err   error
}

func fetch(ct context.Context, userID int) (string, error) {

	ctx, cancel := context.WithTimeout(context.WithValue(ct, "id", userID), 2*time.Second)
	defer cancel()

	respch := make(chan Resp)

	go func() {
		value, err := fetchName(ctx)
		respch <- Resp{
			value: value,
			err:   err,
		}

	}()

	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("context time out")
		case val := <-respch:
			return val.value, val.err

		}
	}

}

func fetchName(ctx context.Context) (string, error) {
	id := ctx.Value("id")

	for i, v := range details {

		if i == id {
			return v, nil
		}
		time.Sleep(1 * time.Second)

	}

	return "", fmt.Errorf("id not found")
}

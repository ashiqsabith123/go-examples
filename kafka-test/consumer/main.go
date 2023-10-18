package main

import (
	"fmt"

	"github.com/IBM/sarama"
	"gocv.io/x/gocv"
)

func main() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("my-topic4", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	fmt.Println(partitionConsumer)
	defer partitionConsumer.Close()

	// Handle Ctrl+C signal to gracefully exit
	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for message := range partitionConsumer.Messages() {
		fmt.Println(message.Value)
		frame, err := gocv.IMDecode(message.Value, gocv.IMReadUnchanged)
		if err != nil {
			fmt.Println("error while decoding", err)
		}
		if frame.Empty() {
			fmt.Println("Received an empty frame")
			continue
		}

		// Display the frame (you can adjust the window title)
		gocv.IMWrite("Received Frame", frame)
		//gocv.WaitKey(1)

		frame.Close() // Re

	}

	// Create a new Mat and load the frame data from the Kafka message

	// Wait for the Ctrl+C signal to exit
	//<-sigchan
}

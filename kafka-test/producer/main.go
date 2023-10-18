package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"gocv.io/x/gocv"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	camera, _ := gocv.OpenVideoCapture(0)
	if camera.IsOpened() {
		defer camera.Close()
	} else {
		fmt.Println("Error: camera not opened")
		return
	}

	frame := gocv.NewMat()
	defer frame.Close()

	// Handle Ctrl+C signal to gracefully exit
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		camera.Read(&frame)
		if frame.Empty() {
			break
		}

		// Convert the frame to a byte slice
		frameBytes := frame.ToBytes()

		// Create a Kafka message with the frame data
		message := &sarama.ProducerMessage{
			Topic:     "my-topic4",
			Partition: 0,
			Value:     sarama.ByteEncoder(frameBytes),
		}

		// Send the message to Kafka
		producer.SendMessage(message)
		fmt.Printf("video send partition: %v ", 0)
	}

	// Wait for the Ctrl+C signal to exit
	<-sigchan
}

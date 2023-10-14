package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Set up the Kafka producer configuration.
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	//config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.MaxMessageBytes = 10 * 1024 * 1024 * 1024
	//config.Producer.Compression = sarama.CompressionGZIP // Use GZIP compression, for example

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	videoData, err := ReadBinaryFile("WhatsApp Video 2023-10-14 at 4.29.38 PM.mp4")
	if err != nil {
		log.Fatalf("Failed to read the video file: %v", err)
	}

	videoBase64 := base64.StdEncoding.EncodeToString(videoData)

	maxChunkSize := 100

	chunks := chunkString(videoBase64, maxChunkSize)

	for i, chunk := range chunks {
		topic := "large-data-topic"
		partition := int32(i) % 3 // Example partition selection logic
		//key := "chunk-" + fmt.Sprint(i)

		message := &sarama.ProducerMessage{
			Topic:     topic,
			Partition: partition,
			//Key:       sarama.StringEncoder(key),
			Value: sarama.StringEncoder(chunk),
		}

		// Send the chunk to Kafka
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send chunk %d: %v", i, err)
		}

		log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		//time.Sleep(3 * time.Second)
	}

	//fmt.Println(videoData)

	// Create a message to send.
	// message := &sarama.ProducerMessage{
	// 	Topic:     "my-topic2",
	// 	Partition: 0,
	// 	Value:     sarama.StringEncoder(videoBase64),
	// }

	// // Send the message.
	// partition, offset, err := producer.SendMessage(message)
	// if err != nil {
	// 	log.Fatalf("Failed to send message: %v", err)
	// }

}

func ReadBinaryFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

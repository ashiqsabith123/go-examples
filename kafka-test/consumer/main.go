package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	videoData []byte
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/ws", sendData)
	router.Run(":8084")

}

func sendData(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	var wg sync.WaitGroup
	//var data []byte
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to the Kafka topic.
	topic := "large-data-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}

	wg.Add(1)

	go func() {
		for message := range partitionConsumer.Messages() {
			log.Printf("Received video message at offset %d\n", message.Offset)

			// Decode base64 to binary data
			videoBase64 := message.Value
			videoData, err := base64.StdEncoding.DecodeString(string(videoBase64))
			if err != nil {
				log.Fatalf("Failed to decode video: %v", err)
			}

			fmt.Println(videoData)
			//data = append(data, videoData...)

			err = conn.WriteMessage(websocket.BinaryMessage, videoData)
			if err != nil {
				log.Println(err)
				return
			}

		}

		wg.Done()
	}()

	wg.Wait()

}

func ReadBinaryFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

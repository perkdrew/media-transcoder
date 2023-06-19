package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Create a new Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	// Subscribe to a Kafka topic
	topics := []string{"test-topic"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}

	// Start consuming messages from the subscribed topics
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigChan:
			// Received termination signal
			fmt.Println("Received termination signal. Shutting down...")
			consumer.Close()
			return
		default:
			// Poll for new messages
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				continue
			}

			fmt.Printf("Received message: %s\n", string(msg.Value))
		}
	}
}

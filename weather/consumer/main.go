package main

import (
	"log"
)

func main() {

	consumer, err := NewConsumer()
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	// Start polling messages and uploading to S3
	log.Println("starting to poll messages from Kafka...")
	consumer.PollMessages()
}

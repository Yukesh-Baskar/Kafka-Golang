package main

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Define Kafka broker address and topic
	brokerUrl := []string{"localhost:9092"}
	topic := "producer-event"

	// Set up Sarama config and consumer group
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // Make sure the version matches your Kafka version

	// Create a new consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokerUrl, "abc", config)
	if err != nil {
		log.Fatal("Error creating consumer group: ", err)
	}
	defer func() {
		if err := consumerGroup.Close(); err != nil {
			log.Fatal("Error closing consumer group: ", err)
		}
	}()

	// Define a consumer handler
	handler := &ConsumerHandler{}

	// Start consuming in an infinite loop
	for {
		// Consume messages from the topic (blocking call)
		if err := consumerGroup.Consume(context.Background(), []string{topic}, handler); err != nil {
			log.Fatal("Error consuming messages: ", err)
		}
	}
}

// ConsumerHandler handles the messages for the consumer group
type ConsumerHandler struct{}

// Setup is called before consumption begins
func (ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	// Initialize your setup here (e.g., database connections)
	return nil
}

// Cleanup is called after consumption ends
func (ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	// Cleanup resources if needed
	return nil
}

// ConsumeClaim processes the messages from a topic partition
func (ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Process messages in this partition
	for message := range claim.Messages() {
		fmt.Printf("Message received: Offset = %d, Key = %s, Value = %s\n", message.Offset, string(message.Key), string(message.Value))

		// Mark the message as processed
		session.MarkMessage(message, "")
	}

	return nil
}

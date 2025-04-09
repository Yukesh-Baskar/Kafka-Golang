package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// just a simple consumer
func main() {
	brokerUrl := []string{"localhost:9092"}
	topic := "producer-event"
	partition := 0

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokerUrl,
		Topic:          topic,
		Partition:      partition,
		MaxBytes:       10e6,
		MaxWait:        time.Duration(time.Second * 5),
		MaxAttempts:    5,
		CommitInterval: time.Second, // Commit offsets every second
	})
	// r.SetOffset(2)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

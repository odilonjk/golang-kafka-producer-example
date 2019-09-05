package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaServer = "localhost:9092"
	topicName   = "transaction"
	tcp         = "tcp"
	partition   = 0
)

type Transaction struct {
	TransactionId string
	Value         float64
	PersonId      int64
	CreationDate  string
}

func main() {

	start := time.Now()

	Configure([]string{kafkaServer}, topicName)

	messages := getTransactionMessages()

	push(context.Background(), messages)

	writer.Close()

	took := time.Since(start)
	fmt.Printf("Elapsed time: %.4fs\n", took.Seconds())

	os.Exit(0)

}

func getTransactionMessages() []kafka.Message {
	messages := []kafka.Message{}
	for i := 1; i < 101; i++ {
		jsonMessage, _ := getTransactionJson()
		message := kafka.Message{
			Value: []byte(jsonMessage),
		}
		messages = append(messages, message)
	}
	return messages
}

func getTransactionJson() ([]byte, error) {
	transaction := getTransaction()
	return json.Marshal(transaction)
}

func push(parent context.Context, messages []kafka.Message) {
	err := writer.WriteMessages(parent, messages...)
	if err != nil {
		fmt.Println("Error pushing into Kafka.")
		os.Exit(1)
	}
}

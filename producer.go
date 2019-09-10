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
	kafkaBroker1 = "localhost:19092"
	kafkaBroker2 = "localhost:29092"
	kafkaBroker3 = "localhost:39092"
	topicName    = "transaction"
	tcp          = "tcp"
	partition    = 0
)

type Transaction struct {
	TransactionId string
	Value         float64
	PersonId      int
	CreationDate  string
}

func main() {

	start := time.Now()

	Configure([]string{kafkaBroker1, kafkaBroker2, kafkaBroker3}, topicName)

	messages := getTransactionMessages()

	push(context.Background(), messages)

	writer.Close()

	took := time.Since(start)
	fmt.Printf("Elapsed time: %.4fs\n", took.Seconds())

	os.Exit(0)

}

func getTransactionMessages() []kafka.Message {
	messages := []kafka.Message{}
	for i := 1; i < 5001; i++ {
		jsonMessage, key := getTransactionJson()
		message := kafka.Message{
			Key:   []byte(string(key)),
			Value: []byte(jsonMessage),
		}
		messages = append(messages, message)
	}
	return messages
}

func getTransactionJson() ([]byte, int) {
	transaction := getTransaction()
	key := transaction.PersonId
	json, _ := json.Marshal(transaction)
	return json, key
}

func push(parent context.Context, messages []kafka.Message) {
	err := writer.WriteMessages(parent, messages...)
	if err != nil {
		fmt.Println("Error pushing into Kafka.")
		os.Exit(1)
	}
}

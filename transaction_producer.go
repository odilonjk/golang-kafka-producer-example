package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	maxId = 4
	minId = 1
)

func getTransaction() Transaction {
	return Transaction{
		TransactionId: getTransactionId(),
		Value:         getValue(),
		PersonId:      getPersonId(),
		CreationDate:  getCreationDate(),
	}
}

func getTransactionId() string {
	transactionId, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return transactionId.String()
}

func getValue() float64 {
	rand.Seed(time.Now().UnixNano())
	randomValue := 1 + rand.Float64()*(999-1)
	return math.Round(randomValue*100) / 100
}

func getPersonId() int {
	return rand.Intn(maxId-minId+minId) + minId
}

func getCreationDate() string {
	return time.Now().Format("2006-01-02T15:04:05")
}

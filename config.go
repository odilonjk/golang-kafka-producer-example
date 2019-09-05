package main

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

const (
	timeoutLimit = 5 * time.Second
	batchSize    = 100
	batchTimeout = 10 * time.Millisecond
)

var writer *kafka.Writer

func Configure(kafkaBrokerUrls []string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout: timeoutLimit,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		BatchSize:        batchSize,
		BatchTimeout:     batchTimeout,
		WriteTimeout:     timeoutLimit,
		ReadTimeout:      timeoutLimit,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}

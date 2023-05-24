package config

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var Writer *kafka.Writer

var (
	kafkaBrokerUrl string
	kafkaVerbose   bool
	kafkaClientId  string
)

func Configure(topic string) (w *kafka.Writer, err error) {

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
	}

	kafkaBrokerUrls := []string{"localhost:9092"}
	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	Writer = w
	return w, nil
}
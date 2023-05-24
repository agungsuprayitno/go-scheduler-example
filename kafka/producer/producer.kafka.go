package producer

import (
	"context"
	"go-rest-postgres/kafka/config"
	"time"

	"github.com/segmentio/kafka-go"
)

func PushMessage(parent context.Context, key, value []byte) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return config.Writer.WriteMessages(parent, message)
}
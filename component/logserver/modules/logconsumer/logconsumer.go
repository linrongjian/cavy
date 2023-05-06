package logconsumer

import (
	"CavyGo/core/network/amqp"
	"fmt"
	"os"
)

type LogConsumer interface {
	Shutdown() error
}

type Option func(*Options)

type Handle func(*Options)

func NewLogConsumer() (LogConsumer, error) {
	var err error
	uri := os.Getenv("LOG_RABBITMQ_URI")
	exchange := os.Getenv("LOG_RABBITMQ_EXCHANGE")
	queue := os.Getenv("LOG_RABBITMQ_QUEUE")
	if uri == "" || exchange == "" || queue == "" {
		return nil, fmt.Errorf("log consumer no env")
	}

	_, err = amqp.NewConsumer(uri, exchange, "topic", queue, "test-key", "log-consumer")
	if err != nil {
		return nil, fmt.Errorf("new log consumer: %s", err)
	}
	return nil, err
}

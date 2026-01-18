package app

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return conn
}

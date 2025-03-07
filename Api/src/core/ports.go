package core

import "github.com/streadway/amqp"

type Producer interface {
	SendMessageToQueue(queueName string, message []byte) error
}

type RabbitMQConnection interface {
	Connect() (*amqp.Connection, error)
}
package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RabbitMQAdapter struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	if user == "" || password == "" || host == "" || port == "" {
		log.Fatal("Error: Algunas variables de entorno no están definidas.")
	}

	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQAdapter{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func (r *RabbitMQAdapter) SendMessageToQueue(queueName string, message []byte) error {

	_, err := r.Channel.QueueDeclare(
		queueName,
		true,
		false, 
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = r.Channel.Publish(
		"", 
		queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         message,
		},
	)

	return err
}
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	

	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	fmt.Println("RABBITMQ_USER:", os.Getenv("RABBITMQ_USER"))
	fmt.Println("RABBITMQ_PASSWORD:", os.Getenv("RABBITMQ_PASSWORD"))
	fmt.Println("RABBITMQ_HOST:", os.Getenv("RABBITMQ_HOST"))
	fmt.Println("RABBITMQ_PORT:", os.Getenv("RABBITMQ_PORT"))


	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"status_machine",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	failOnError(err, "Failed to bind the queue to the exchange")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		log.Printf(" [x] Recibido: %s", d.Body)
	}
}
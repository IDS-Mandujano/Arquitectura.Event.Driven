package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func sendToAPI(message []byte) error {

    messageObj := map[string]string{
        "Content": string(message),
    }

    messageBytes, err := json.Marshal(messageObj)
    if err != nil {
        return fmt.Errorf("error al convertir el mensaje a JSON: %w", err)
    }

    body := bytes.NewBuffer(messageBytes)

    resp, err := http.Post("http://54.208.214.104:3002/receive", "application/json", body)
    if err != nil {
        return fmt.Errorf("error al enviar mensaje a la API: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        responseBody, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("la API respondió con código %d: %s", resp.StatusCode, string(responseBody))
    }

    log.Printf("Mensaje enviado a la API: %s", message)
    return nil
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

		err := sendToAPI(d.Body)
		if err != nil {
			log.Printf("Error al enviar mensaje a la API: %s", err)
		}
	}
}
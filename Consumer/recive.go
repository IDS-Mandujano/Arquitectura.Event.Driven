package main

import (
	"log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://emandujano:Cacatua$99@54.225.236.159:5672/")
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

	for d := range msgs {
		log.Printf(" [x] Recibido: %s", d.Body)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
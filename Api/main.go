package main

import (
	application "api-producer/src/message/Application"
	"api-producer/src/message/Infraestructure"
	"log"
	"net/http"
)

func main() {
	rabbitAdapter, err := infrastructure.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error creating RabbitMQ adapter: %s", err)
	}
	defer rabbitAdapter.Connection.Close()
	defer rabbitAdapter.Channel.Close()

	messageService := application.NewMessageService(rabbitAdapter)

	messageHandler := application.NewMessageHandler(messageService)

	http.HandleFunc("/receive", messageHandler.HandleMessage)

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
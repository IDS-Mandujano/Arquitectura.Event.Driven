package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	upgrader  websocket.Upgrader
}

var wsServer = WebSocketServer{
	clients:   make(map[*websocket.Conn]bool),
	broadcast: make(chan []byte),
	upgrader: websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := wsServer.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar conexión WebSocket:", err)
		return
	}
	defer ws.Close()

	wsServer.clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado:", err)
			delete(wsServer.clients, ws)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-wsServer.broadcast
		for client := range wsServer.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error enviando mensaje:", err)
				client.Close()
				delete(wsServer.clients, client)
			}
		}
	}
}

func listenToRabbitMQ() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal("Error al conectar con RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"machine_status",
		true,  
		false,     
		false,   
		false, 
		nil,     
	)
	if err != nil {
		log.Fatal("Error al declarar la cola:", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error al consumir mensajes:", err)
	}

	log.Println("Esperando mensajes en la cola 'machine_status'...")

	for msg := range msgs {
		log.Printf("Mensaje recibido: %s", msg.Body)
		wsServer.broadcast <- msg.Body
	}
}

func main() {
	go handleMessages()
	go listenToRabbitMQ()

	http.HandleFunc("/ws", handleConnections)

	port := ":8080"
	log.Println("Servidor WebSocket corriendo en http://54.225.236.159" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
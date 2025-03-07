package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQConfig, error) {

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


	dns := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(dns)
	if err != nil {
		log.Printf("Error conectando a RabbitMQ: %s", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error abriendo canal en RabbitMQ: %s", err)
		conn.Close()
		return nil, err
	}

	log.Println("Conexión a RabbitMQ establecida exitosamente.")

	return &RabbitMQConfig{Conn: conn, Ch: ch}, nil
}

func (r *RabbitMQConfig) Close() {
	if r.Ch != nil {
		r.Ch.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
	log.Println("Conexión a RabbitMQ cerrada.")
}
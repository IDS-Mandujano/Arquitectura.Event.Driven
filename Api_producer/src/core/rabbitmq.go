package core

import (
	"fmt"
	"log"

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
		log.Println("Advertencia: No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	dns := fmt.Sprintf("amqp://emandujano:Cacatua$99@54.225.236.159:5672/")

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
package adapters

import (
	"fmt"
	"gym-system/src/core"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQProducer struct {
	RabbitMQ *core.RabbitMQConfig
}

func NewRabbitMQProducer(rabbitMQ *core.RabbitMQConfig) *RabbitMQProducer {
	return &RabbitMQProducer{RabbitMQ: rabbitMQ}
}

func (r *RabbitMQProducer) PublishMachineStatus(machineID int, status string) error {
	message := fmt.Sprintf("El estado de la máquina %d ha sido actualizado a %s.", machineID, status)
	err := r.PublishMessage("machine_status_updates", message)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
	}
	return nil
}

func (r *RabbitMQProducer) PublishMessage(queueName, message string) error {
	
	err := r.RabbitMQ.Ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false, 
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error al declarar el intercambio: %s", err)
		return err
	}

	_, err = r.RabbitMQ.Ch.QueueDeclare(
		"status_machine",
		true,
		false, 
		false, 
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error al declarar la cola: %s", err)
		return err
	}

	err = r.RabbitMQ.Ch.QueueBind(
		"status_machine",
		"",
		"logs", 
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error al enlazar la cola con el intercambio: %s", err)
		return err
	}

	err = r.RabbitMQ.Ch.Publish(
		"logs",
		"",     
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Error al publicar el mensaje: %s", err)
		return err
	}

	log.Printf("Mensaje enviado al intercambio 'logs': %s", message)

	return nil
}
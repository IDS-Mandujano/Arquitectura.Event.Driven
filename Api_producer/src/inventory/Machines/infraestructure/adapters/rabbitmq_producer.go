package adapters

import (
	"fmt"
	"gym-system/src/core"
	"gym-system/src/inventory/Machines/domain/repository"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	RabbitMQ *core.RabbitMQConfig
}

var _ repository.MachineStatusPublisher = (*RabbitMQProducer)(nil)

func NewRabbitMQProducer(rabbitMQ *core.RabbitMQConfig) *RabbitMQProducer {
	return &RabbitMQProducer{RabbitMQ: rabbitMQ}
}

func (r *RabbitMQProducer) PublishMachineStatus(machineID int, status string) error {
	message := fmt.Sprintf("El estado de la máquina %d ha sido actualizado a %s.", machineID, status)
	return r.PublishMessage("machine_status_updates", message)
}

func (r *RabbitMQProducer) PublishMessage(queueName, message string) error {
	err := r.RabbitMQ.Ch.Publish(
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
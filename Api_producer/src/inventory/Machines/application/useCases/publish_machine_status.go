package machineusecases

import (
	"fmt"
	"gym-system/src/inventory/Machines/infraestructure/adapters"
)

type PublishMachineStatusService struct {
	RabbitMQProducer *adapters.RabbitMQProducer
}

func NewPublishMachineStatusService(rabbitMQProducer *adapters.RabbitMQProducer) *PublishMachineStatusService {
	return &PublishMachineStatusService{RabbitMQProducer: rabbitMQProducer}
}

func (service *PublishMachineStatusService) Publish(machineID int, status string) error {
	err := service.RabbitMQProducer.PublishMachineStatus(machineID, status)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
	}
	return nil
}

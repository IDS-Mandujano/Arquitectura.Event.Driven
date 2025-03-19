package machineusecases

import (
	"fmt"
	"gym-system/src/inventory/Machines/domain/repository"
)

type PublishMachineStatusService struct {
	Publisher repository.MachineStatusPublisher
}

func NewPublishMachineStatusService(publisher repository.MachineStatusPublisher) *PublishMachineStatusService {
	return &PublishMachineStatusService{Publisher: publisher}
}

func (service *PublishMachineStatusService) Publish(machineID int, status string) error {
	err := service.Publisher.PublishMachineStatus(machineID, status)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje: %w", err)
	}
	return nil
}
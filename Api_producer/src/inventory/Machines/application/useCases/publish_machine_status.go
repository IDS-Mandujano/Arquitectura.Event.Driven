package machineusecases

import (
	"fmt"
	"gym-system/src/inventory/Machines/application/repository"
)

type PublishMachineStatusUseCase struct {
	MachineStatusRepository repository.MachineStatusRepository
}

func NewPublishMachineStatusUseCase(machineStatusRepository repository.MachineStatusRepository) *PublishMachineStatusUseCase {
	return &PublishMachineStatusUseCase{MachineStatusRepository: machineStatusRepository}
}

func (useCase *PublishMachineStatusUseCase) Execute(machineID int, status string) error {
	err := useCase.MachineStatusRepository.PublishMachineStatus(machineID, status)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje: %w", err)
	}
	return nil
}

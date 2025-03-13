package machineusecases

import (
	"gym-system/src/inventory/Machines/domain/repository"
	"log"
)

type UpdateMachine struct {
	db                    repository.IMachineRepository
	publishStatusService  *PublishMachineStatusService
}

func NewUpdateMachine(db repository.IMachineRepository, publishStatusService *PublishMachineStatusService) *UpdateMachine {
	return &UpdateMachine{
		db:                   db,
		publishStatusService: publishStatusService,
	}
}

func (updateMachine *UpdateMachine) Execute(id int, cname string, ctype string, cstatus string) {
	updateMachine.db.Update(id, cname, ctype, cstatus)
	
	err := updateMachine.publishStatusService.Publish(id, cstatus)
	if err != nil {
		log.Printf("Error al publicar el estado de la máquina: %s", err)
	}
}
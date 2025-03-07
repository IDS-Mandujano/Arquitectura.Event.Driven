package machineusecases

import "gym-system/src/inventory/Machines/domain/repository"

type GetStatusMachine struct {
	db repository.IMachineRepository
}

func NewMachineStatus(db repository.IMachineRepository) *GetStatusMachine {
	return &GetStatusMachine{db: db}
}

func (getMStatus *GetStatusMachine) Execute(id int) (string, error) {
	return getMStatus.db.GetStatus(id)
}

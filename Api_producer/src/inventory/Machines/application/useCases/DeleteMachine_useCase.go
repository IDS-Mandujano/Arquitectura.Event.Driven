package machineusecases

import "gym-system/src/inventory/Machines/domain/repository"

type DeleteMachine struct {
	db repository.IMachineRepository
}

func NewDeleteMachine(db repository.IMachineRepository) *DeleteMachine {
	return &DeleteMachine{db: db}
}

func (deleteMachine *DeleteMachine) Execute(id int) {
	deleteMachine.db.Delete(id)
}

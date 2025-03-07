package machineusecases

import "gym-system/src/inventory/Machines/domain/repository"

type ListMachine struct {
	db repository.IMachineRepository
}

func NewListMachine(db repository.IMachineRepository) *ListMachine {
	return &ListMachine{db: db}
}

func (listMachine *ListMachine) Execute() ([]map[string]interface{}, error) {
	return listMachine.db.GetAll()
}

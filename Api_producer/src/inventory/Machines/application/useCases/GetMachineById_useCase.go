package machineusecases

import "gym-system/src/inventory/Machines/domain/repository"

type GetMachineById struct {
	db repository.IMachineRepository
}

func NewMachineById(db repository.IMachineRepository) *GetMachineById {
	return &GetMachineById{db: db}
}

func (getMachine *GetMachineById) Execute(id int) ([]map[string]interface{}, error) {
	return getMachine.db.GetById(id)
}

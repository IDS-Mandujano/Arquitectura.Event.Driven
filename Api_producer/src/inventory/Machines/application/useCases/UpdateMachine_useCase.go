package machineusecases

import "gym-system/src/inventory/Machines/domain/repository"

type UpdateMachine struct {
	db repository.IMachineRepository
}

func NewUpdateMachine(db repository.IMachineRepository) *UpdateMachine {
	return &UpdateMachine{db: db}
}

func (updateMachine *UpdateMachine) Execute(id int, cname string, ctype string, cstatus string) {
	updateMachine.db.Update(id, cname, ctype, cstatus)
}

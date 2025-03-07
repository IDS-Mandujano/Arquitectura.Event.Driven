package machineusecases

import "gym-system/src/inventory/Machines/domain/repository"

type CreateMachine struct {
	db repository.IMachineRepository
}

func NewCreateMachine(db repository.IMachineRepository) *CreateMachine {
	return &CreateMachine{db: db}
}

func (createMachine *CreateMachine) Execute(cname string, ctype string, cstatus string){
	createMachine.db.Save(cname, ctype, cstatus)
}
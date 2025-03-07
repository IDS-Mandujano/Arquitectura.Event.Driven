package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type CreateEquipament struct {
	db repository.IEquipamentRepository
}

func NewCreateEquipament(db repository.IEquipamentRepository) *CreateEquipament {
	return &CreateEquipament{db: db}
}

func (ce *CreateEquipament) Execute(cname string, category string, ccondition string){
	ce.db.Save(cname, category, ccondition)
}
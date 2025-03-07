package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type UpdateEquipment struct {
	db repository.IEquipamentRepository
}

func NewUpdateEquipment(db repository.IEquipamentRepository) *UpdateEquipment {
	return &UpdateEquipment{db: db}
}

func (updateEquipment *UpdateEquipment) Execute (id int,cname string, category string, ccondition string){
	updateEquipment.db.Update(id,cname,category,ccondition)
}
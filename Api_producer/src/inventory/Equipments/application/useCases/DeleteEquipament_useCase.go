package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type DeleteEquipment struct {
	db repository.IEquipamentRepository
}

func NewDeleteEquipment (db repository.IEquipamentRepository) *DeleteEquipment {
	return &DeleteEquipment{db: db}
}

func (deleteEquipment *DeleteEquipment) Execute(id int){
	deleteEquipment.db.Delete(id)
}
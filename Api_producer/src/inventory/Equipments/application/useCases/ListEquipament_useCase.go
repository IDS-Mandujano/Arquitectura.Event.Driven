package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type ListEquipament struct {
	db repository.IEquipamentRepository
}

func NewListEquipment (db repository.IEquipamentRepository) *ListEquipament {
	return &ListEquipament{db: db}
}

func (listEquipment *ListEquipament) Execute () ([]map[string]interface{}, error){
	return listEquipment.db.GetAll()
}
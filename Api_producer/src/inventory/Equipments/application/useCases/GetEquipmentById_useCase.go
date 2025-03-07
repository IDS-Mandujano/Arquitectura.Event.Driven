package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type GetEquipmentById struct {
	db repository.IEquipamentRepository
}

func NewEquipmentById(db repository.IEquipamentRepository) *GetEquipmentById {
	return &GetEquipmentById{db: db}
}

func (getEquipment *GetEquipmentById) Execute(id int) ([]map[string]interface{},error){
	return getEquipment.db.GetById(id)
}
package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type GetEquipmentByCondition struct {
	db repository.IEquipamentRepository
}

func NewEquipmentByCondition (db repository.IEquipamentRepository) *GetEquipmentByCondition {
	return &GetEquipmentByCondition{db: db}
}

func (getEquipmentByCondition *GetEquipmentByCondition) Execute (condition string) ([]map[string]interface{}, error){
	return getEquipmentByCondition.db.GetCondition(condition)
}
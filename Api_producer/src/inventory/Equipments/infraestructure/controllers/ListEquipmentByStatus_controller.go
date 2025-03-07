package equipamentcontrollers

import (
	equipamentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetEquipmentByCondition struct {
	useCase *equipamentusecases.GetEquipmentByCondition
}

func NewEquipmentCondition(useCase *equipamentusecases.GetEquipmentByCondition) *GetEquipmentByCondition {
	return &GetEquipmentByCondition{useCase: useCase}
}

func (listEquipmentByCondition *GetEquipmentByCondition) Execute(g *gin.Context) {
	condition := g.Param("condition")

	equipment, err := listEquipmentByCondition.useCase.Execute(condition)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el equipo"})
		return
	}

	if len(equipment) == 0 {
		g.JSON(http.StatusNotFound, gin.H{"error": "No hay equipos con la condición especificada"})
		return
	}

	g.JSON(http.StatusOK, equipment)
}

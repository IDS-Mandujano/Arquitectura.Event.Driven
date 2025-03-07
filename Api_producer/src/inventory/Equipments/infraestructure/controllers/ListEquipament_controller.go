package equipamentcontrollers

import (
	equipamentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListEquipamentController struct {
	useCase equipamentusecases.ListEquipament
}

func NewListEquipmentController(useCase equipamentusecases.ListEquipament) *ListEquipamentController {
	return &ListEquipamentController{useCase: useCase}
}

func (listEquipment_controller *ListEquipamentController) Execute(g *gin.Context) {
	equipments, err := listEquipment_controller.useCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": "Error al obtener los equipos"})
		return
	}

	g.JSON(http.StatusOK, equipments)
}

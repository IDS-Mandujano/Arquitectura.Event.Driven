package equipamentcontrollers

import (
	equipamentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEquipmentController struct {
	useCase equipamentusecases.DeleteEquipment
}

func NewDeleteEquipment(useCase equipamentusecases.DeleteEquipment) *DeleteEquipmentController {
	return &DeleteEquipmentController{useCase: useCase}
}

func (deleteEquipment *DeleteEquipmentController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Id del equipo invalido"})
		return
	}

	deleteEquipment.useCase.Execute(id)
	g.JSON(http.StatusOK, gin.H{"message": "Equipo eliminado correctamente"})
}

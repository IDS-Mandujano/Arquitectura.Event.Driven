package equipamentcontrollers

import (
	equipamentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEquipmentController struct {
	useCase equipamentusecases.UpdateEquipment
}

func NewUpdateEquipmentController(useCase equipamentusecases.UpdateEquipment) *UpdateEquipmentController {
	return &UpdateEquipmentController{useCase: useCase}
}

func (updateEquipment *UpdateEquipmentController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Id del equipo invalido"})
		return
	}

	var input struct {
		Name      string `json:"name"`
		Category  string `json:"category"`
		Condition string `json:"condition"`
	}

	if err := g.ShouldBindJSON(&input); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEquipment.useCase.Execute(id, input.Name, input.Category, input.Condition)
	g.JSON(http.StatusOK, gin.H{"message": "Equipo editado con exito"})
}

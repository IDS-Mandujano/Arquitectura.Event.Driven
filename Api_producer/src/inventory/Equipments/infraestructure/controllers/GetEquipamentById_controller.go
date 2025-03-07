package equipamentcontrollers

import (
	equipamentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEquipmentByIdController struct {
	useCase equipamentusecases.GetEquipmentById
}

func NewEquipmentByIdController(useCase equipamentusecases.GetEquipmentById) *GetEquipmentByIdController {
	return &GetEquipmentByIdController{useCase: useCase}
}

func (getEquipment *GetEquipmentByIdController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Id no valido"})
		return
	}

	equipment, err := getEquipment.useCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el equipo"})
		return
	}

	if equipment == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Equipo no encontrado"})
		return
	}

	g.JSON(http.StatusOK, equipment)
}

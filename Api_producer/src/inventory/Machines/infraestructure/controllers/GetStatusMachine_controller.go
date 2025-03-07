package machinecontrollers

import (
	machineusecases "gym-system/src/inventory/Machines/application/useCases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetStatusMachineController struct {
	useCase *machineusecases.GetStatusMachine
}

func NewStatusMachine(useCase *machineusecases.GetStatusMachine) *GetStatusMachineController {
	return &GetStatusMachineController{useCase: useCase}
}

func (getStatusMachine *GetStatusMachineController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": "ID no valido"})
		return
	}

	status, err := getStatusMachine.useCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": "Error al obtener estado de la maquina"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": status})
}

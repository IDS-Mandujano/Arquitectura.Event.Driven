package machinecontrollers

import (
	machineusecases "gym-system/src/inventory/Machines/application/useCases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMachineController struct {
	useCase machineusecases.DeleteMachine
}

func NewDeleteMachine(useCase machineusecases.DeleteMachine) *DeleteMachineController {
	return &DeleteMachineController{useCase: useCase}
}

func (deleteMachine *DeleteMachineController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Id del equipo invalido"})
		return
	}

	deleteMachine.useCase.Execute(id)
	g.JSON(http.StatusOK, gin.H{"message": "Maqui eliminado correctamente"})
}
